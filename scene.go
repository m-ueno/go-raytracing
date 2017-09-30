package main

import (
	"fmt"
	"math"
	"math/rand"
)

// 物体と光源の集合
type Scene struct {
	shapes           []Shape
	lightSources     []LightSource
	ambientIntensity float64
	size             int
}

func (sc *Scene) shadingAmbient(shape Shape) *FColor {
	return FCScale(sc.ambientIntensity, shape.Material().k_a)
}

func (sc *Scene) shadingDiffuse(
	ip *IntersectionPoint,
	lightSource LightSource,
	shape Shape) *FColor {
	lighting := lightSource.LightingAt(ip.position)
	k_d := shape.Material().k_d

	v_n := ip.normal
	v_l := Scale(-1.0, lighting.direction)
	dot := math.Max(0, Dot(v_l, v_n))
	r_d := FCScale(dot*lighting.intensity, k_d)

	return r_d
}

func (sc *Scene) shadingSpecular(
	ip *IntersectionPoint,
	lightSource LightSource,
	ray *Ray,
	shape Shape) *FColor {

	lighting := lightSource.LightingAt(ip.position)
	k_s := shape.Material().k_s
	v_n := ip.normal
	v_l := Scale(-1.0, lighting.direction)
	to := ray.direction

	v_v := Scale(-1.0, to)
	v_r := Sub(Scale(2.0*Dot(v_n, v_l), v_n), v_l) //# 正反射ベクトル. 交点で入射光が反射する
	// vR = 2 * cos * vN - vL

	dot := math.Max(Dot(v_v, v_r), 0.0)
	r_s := FCScale(lighting.intensity*math.Pow(dot, shape.Material().shininess), k_s)

	return r_s
}

func (sc *Scene) testIntersectionWithAll(ray *Ray) (*IntersectionTestResult, bool) {
	return sc.testIntersectionWithAllFullParam(ray, math.MaxFloat64, false)
}

func (sc *Scene) testIntersectionWithAllFullParam(ray *Ray, maxDist float64, exitOnceFound bool) (*IntersectionTestResult, bool) {
	var nearestShape Shape = nil
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
	lighting := lightSource.LightingAt(ip.position)
	v_l := Scale(-1.0, lighting.direction)
	shadowRay := &Ray{
		direction: v_l,
		start:     Add(ip.position, Scale(EPSILON, v_l)),
	}

	_, found := sc.testIntersectionWithAllFullParam(shadowRay, lighting.distance-EPSILON, true)

	return found
}

var backgroundColor *FColor = NewFColor(100/255.0, 149/255.0, 237/255.0)

func (sc *Scene) rayTrace(ray *Ray) *FColor {
	return sc.rayTraceRecursive(ray, 0)
}

func (sc *Scene) rayTraceRecursive(ray *Ray, recLevel int) *FColor {
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
			fcolor = FCAdd(fcolor, sc.shadingDiffuse(ip, ls, shape))
			fcolor = FCAdd(fcolor, sc.shadingSpecular(ip, ls, ray, shape))
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
					fcolor = FCAdd(fcolor, FCScale(mat.catadioptricFactor, r_reOrNil))
				}
			}
		}
	}

	return fcolor
}

func (sc *Scene) renderHead() {
	fmt.Printf("P3\n%d %d\n255\n", sc.size, sc.size)
}

func (sc *Scene) render(antialiasing bool) {
	size := sc.size
	from := NewVector(0, 0, -5)

	sc.renderHead()

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fcolor := NewFColor(0, 0, 0)

			if antialiasing {
				fcolorAcc := NewFColor(0, 0, 0)
				nSample := 10

				for s := 0; s < nSample; s++ {
					screenXYZ := makeEyeWithSampling(x, y, size, rand.Float64(), rand.Float64())
					to := Normalize(Sub(screenXYZ, from))
					ray := NewRay(from, to)

					fcolorAcc = FCAdd(fcolorAcc, sc.rayTrace(ray))
				}
				fcolor = FCScale(1.0/float64(nSample), fcolorAcc)
			} else {
				screenXYZ := makeEye(x, y, size)
				to := Normalize(Sub(screenXYZ, from))
				ray := NewRay(from, to)

				fcolor = sc.rayTrace(ray)
			}

			fmt.Print(fcolor)
		}
	}
}
