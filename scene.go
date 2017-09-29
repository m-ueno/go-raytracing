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

func (sc *Scene) rayTrace(ray *Ray, shape Shape) *FColor {
	/* 交差判定 */
	ip, ok := shape.testIntersection(ray)

	fcolor := sc.shadingAmbient(shape)

	if ok {
		for _, ls := range sc.lightSources {
			fcolor = FCAdd(fcolor, sc.shadingDiffuse(ip, ls, shape))
			//			log.Printf("ip:%v, fc:%v\n", ip, fcolor)
		}
	} else {
		fcolor = NewFColor(100/255.0, 149/255.0, 237/255.0)
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

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			screenXYZ := makeEye(x, y, size)

			// 視線方向
			to := Sub(screenXYZ, from)
			to = Normalize(to)

			ray := NewRay(from, to)

			shape := sc.shapes[0]

			fcolor := sc.rayTrace(ray, shape)
			fmt.Print(fcolor)
		}
	}
}
