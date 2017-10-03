package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

// Scene is set of shapes and lightSources
type Scene struct {
	shapes                []Shape
	lightSources          []LightSource
	ambientIntensity      float64
	size                  int
	globalRefractionIndex float64 // 大気の絶対屈折率
}

func newScene(shapes []Shape, lightSources []LightSource, size int) *Scene {
	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: 0.1,
		size:             size,
		globalRefractionIndex: 1.000293,
	}
}

func (sc *Scene) shadingAmbient(shape Shape) *fColor {
	return fCScale(sc.ambientIntensity, shape.Material().k_a)
}

func (sc *Scene) shadingDiffuse(
	ip *IntersectionPoint,
	lightSource LightSource,
	shape Shape) *fColor {
	lighting := lightSource.lightingAt(ip.position)
	k_d := shape.Material().k_d

	v_n := ip.normal
	v_l := Scale(-1.0, lighting.direction)
	dot := math.Max(0, Dot(v_l, v_n))
	r_d := fCScale(dot*lighting.intensity, k_d)

	return r_d
}

func (sc *Scene) shadingSpecular(
	ip *IntersectionPoint,
	lightSource LightSource,
	ray *Ray,
	shape Shape) *fColor {

	lighting := lightSource.lightingAt(ip.position)
	k_s := shape.Material().k_s
	v_n := ip.normal
	v_l := Scale(-1.0, lighting.direction)
	to := ray.direction

	v_v := Scale(-1.0, to)
	v_r := Sub(Scale(2.0*Dot(v_n, v_l), v_n), v_l) //# 正反射ベクトル. 交点で入射光が反射する
	// vR = 2 * cos * vN - vL

	dot := math.Max(Dot(v_v, v_r), 0.0)
	r_s := fCScale(lighting.intensity*math.Pow(dot, shape.Material().shininess), k_s)

	return r_s
}

func (sc *Scene) testIntersectionWithAll(ray *Ray) (*IntersectionTestResult, bool) {
	return sc.testIntersectionWithAllFullParam(ray, math.MaxFloat64, false)
}

func (sc *Scene) testIntersectionWithAllFullParam(ray *Ray, maxDist float64, exitOnceFound bool) (*IntersectionTestResult, bool) {
	var nearestShape Shape
	nearestIP := &IntersectionPoint{
		distance: 1e10,
		normal:   nil,
		position: nil,
	}

	for _, shape := range sc.shapes {
		ip, found := shape.testIntersection(ray)
		if found && ip.distance < maxDist {
			if ip.distance < nearestIP.distance {
				nearestShape = shape
				nearestIP = ip
			}
			if exitOnceFound {
				break
			}
		}
	}

	testResult := &IntersectionTestResult{
		intersectionPoint: nearestIP,
		shape:             nearestShape,
	}

	return testResult, nearestShape != nil
}

func (sc *Scene) testShadow(lightSource LightSource, ip *IntersectionPoint) bool {
	lighting := lightSource.lightingAt(ip.position)
	v_l := Scale(-1.0, lighting.direction)
	shadowRay := &Ray{
		direction: v_l,
		start:     Add(ip.position, Scale(EPSILON, v_l)),
	}

	_, found := sc.testIntersectionWithAllFullParam(shadowRay, lighting.distance-EPSILON, true)

	return found
}

var backgroundColor = newfColor(100/255.0, 149/255.0, 237/255.0)

func (sc *Scene) rayTrace(ray *Ray) *fColor {
	return sc.rayTraceRecursive(ray, 0)
}

func (sc *Scene) rayTraceRecursive(ray *Ray, recLevel int) *fColor {
	if recLevel > 10 {
		return nil // 交差なし
	}
	testResult, found := sc.testIntersectionWithAll(ray)

	// set default color
	fcolor := backgroundColor

	if found {
		shape := testResult.shape
		mat := shape.Material()
		ip := testResult.intersectionPoint

		fcolor = sc.shadingAmbient(shape)

		for _, ls := range sc.lightSources {
			if sc.testShadow(ls, ip) {
				continue
			}
			fcolor = fCAdd(fcolor, sc.shadingDiffuse(ip, ls, shape))
			fcolor = fCAdd(fcolor, sc.shadingSpecular(ip, ls, ray, shape))
		}

		// 完全鏡面反射
		if mat.usePerfectReflectance {
			v_n := ip.normal                             // 法線ベクトル
			v_v := Scale(-1.0, Normalize(ray.direction)) // 視線ベクトルの逆
			cos := Dot(v_v, v_n)
			if cos > 0 { // 表からの進入
				v_r := Sub(Scale(2*cos, v_n), v_v) // 大元の視線ベクトルの、正反射ベクトル
				reRay := &Ray{
					direction: v_r,
					start:     Add(ip.position, Scale(EPSILON, v_r)),
				}

				r_reOrNil := sc.rayTraceRecursive(reRay, recLevel+1)
				if r_reOrNil != nil {
					fcolor = fCAdd(fcolor, fCScale(mat.catadioptricFactor, r_reOrNil))
				}
			}
		}

		// 屈折
		if mat.useRefraction {
			v := Scale(-1, Normalize(ray.direction)) // 視線の逆ベクトル
			n := Normalize(ip.normal)

			tau1 := sc.globalRefractionIndex
			tau2 := shape.Material().refractionIndex

			// ウラ面判定
			cos1 := Dot(v, n)
			if cos1 < 0 { // 裏から入射 (物体->大気)
				tau1, tau2 = tau2, tau1
				n = Scale(-1, n)
				cos1 = Dot(v, n) // 内積を再計算
			}

			tauR := tau2 / tau1
			cos2 := 1.0 / tauR * math.Sqrt(tauR*tauR-(1-cos1*cos1))
			omega := tauR*cos2 - cos1

			reDir := Sub(Scale(2*cos1, n), v) // 大元の視線ベクトルの、正反射ベクトル (alias: v_r)
			feDir := Scale(1/tauR, Sub(Normalize(ray.direction), Scale(omega, n)))
			feDir = Normalize(feDir)

			reRay := &Ray{ // 正反射方向の半直線
				direction: reDir,
				start:     Add(ip.position, Scale(EPSILON, reDir)),
			}
			feRay := &Ray{ // 屈折方向の半直線
				direction: feDir,
				start:     Add(ip.position, Scale(EPSILON, feDir)),
			}
			rhoP := (tauR*cos1 - cos2) / (tauR*cos1 + cos2) // p偏光反射率
			rhoS := -omega / (tauR*cos2 + cos1)             // s偏光反射率
			cR := (rhoP*rhoP + rhoS*rhoS) / 2               // 完全鏡面反射光の割合
			cT := 1.0 - cR                                  // 屈折光の割合

			reROrNil := sc.rayTraceRecursive(reRay, recLevel+1) // 完全鏡面反射光の放射輝度
			if reROrNil != nil {
				fcolor = fCAdd(fcolor, fCScale(cR*mat.catadioptricFactor, reROrNil))
			}
			feROrNil := sc.rayTraceRecursive(feRay, recLevel+1) // 屈折光の放射輝度
			if feROrNil != nil {
				fcolor = fCAdd(fcolor, fCScale(cT*mat.catadioptricFactor, feROrNil))
			}
		}
	}

	return fcolor
}

func (sc *Scene) renderHead(f *os.File) {
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", sc.size, sc.size))
}

func (sc *Scene) render(antialiasing bool, f *os.File) {
	size := sc.size
	from := newVector(0, 0, -5)

	sc.renderHead(f)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fcolor := newfColor(0, 0, 0)

			if antialiasing {
				fcolorAcc := newfColor(0, 0, 0)
				nSample := 10

				for s := 0; s < nSample; s++ {
					screenXYZ := makeEyeWithSampling(x, y, size, rand.Float64(), rand.Float64())
					to := Normalize(Sub(screenXYZ, from))
					ray := newRay(from, to)

					fcolorAcc = fCAdd(fcolorAcc, sc.rayTrace(ray))
				}
				fcolor = fCScale(1.0/float64(nSample), fcolorAcc)
			} else {
				screenXYZ := makeEye(x, y, size)
				to := Normalize(Sub(screenXYZ, from))
				ray := newRay(from, to)

				fcolor = sc.rayTrace(ray)
			}

			f.WriteString(fmt.Sprint(fcolor))
		}
	}
}
