package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	log.Println("hello")

	flag.Parse()

	args := flag.Args()
	size, _ := strconv.Atoi(args[0])

	render(size)

	log.Println("bye")
}

type IntersectionPoint struct {
	distance float64
	position *Vector
	normal   *Vector
}

func shadingAmbient(shape Shape) *FColor {
	return shape.Material().k_a
}

func shadingDiffuse(ip *IntersectionPoint, lightSource LightSource, shape Shape) *FColor {
	lighting := lightSource.LightingAt(ip.position)
	k_d := shape.Material().k_d

	v_n := ip.normal
	v_l := Scale(-1.0, lighting.direction)
	dot := math.Max(0, Dot(v_l, v_n))
	r_d := FCScale(dot*lighting.intensity, k_d)

	return r_d
}

func rayTrace(ray *Ray, shape *Sphere) *FColor {
	/* 交差判定 */
	ip, ok := shape.testIntersection(ray)

	ls := NewPointLightSource(1, NewVector(-5, 5, -5))

	fcolor := NewFColor(0, 0, 0)

	if ok {
		fcolor = FCAdd(fcolor, shadingAmbient(shape))
		fcolor = FCAdd(fcolor, shadingDiffuse(ip, ls, shape))
		log.Printf("ip:%v, fc:%v\n", ip, fcolor)

	} else {
		fcolor = NewFColor(100/255.0, 149/255.0, 237/255.0)
	}

	return fcolor
}

func render(size int) {
	from := NewVector(0, 0, -5)

	// Sphere
	shape := NewSphere(
		NewVector(0, 0, 5),
		1.0,
		NewMaterial(),
	)
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

			fcolor := rayTrace(ray, shape)
			fmt.Print(fcolor)
		}
	}
}

func makeEye(x int, y int, imageSize int) *Vector {
	return NewVector(
		-1.0+float64(x)/float64(imageSize)*2,
		1.0-float64(y)/float64(imageSize)*2,
		0.0,
	)
}
