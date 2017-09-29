package main

import (
	"fmt"
	"math"
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

func (sc *Scene) rayTrace(ray *Ray) *FColor {
	/* 全shapeとの交差判定 */
	// ip, ok := shape.testIntersection(ray)
	testResult, found := sc.testIntersectionWithAll(ray)

	// set default color
	fcolor := NewFColor(100/255.0, 149/255.0, 237/255.0)

	if found {
		shape := testResult.shape
		ip := testResult.intersectionPoint

		fcolor = sc.shadingAmbient(shape)

		for _, ls := range sc.lightSources {
			if sc.testShadow(ls, ip) {
				continue
			}

			fcolor = FCAdd(fcolor, sc.shadingDiffuse(ip, ls, shape))
			fcolor = FCAdd(fcolor, sc.shadingSpecular(ip, ls, ray, shape))
			//			log.Printf("ip:%v, fc:%v\n", ip, fcolor)
		}
	}

	return fcolor
}

func (sc *Scene) render() {
	size := sc.size
	from := NewVector(0, 0, -5)

	// Sphere
	//	center := NewVector(0, 0, 5)
	//	radius := 1.0

	fmt.Printf("P3\n%d %d\n255\n", size, size)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			screenXYZ := makeEye(x, y, size)

			// 視線方向
			to := Sub(screenXYZ, from)
			to = Normalize(to)

			ray := NewRay(from, to)

			fcolor := sc.rayTrace(ray)
			fmt.Print(fcolor)
		}
	}
}
