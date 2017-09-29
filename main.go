package main

import (
	"flag"
	"fmt"
	"log"
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

func shading_ambient(shape *Sphere) *FColor {
	return shape.material.k_a
}

func rayTrace(ray *Ray, shape *Sphere) *FColor {
	/* 交差判定 */
	_, ok := shape.testIntersection(ray)

	fcolor := NewFColor(0, 0, 0)

	if ok {
		fcolor = FCAdd(fcolor, shading_ambient(shape))
	} else {
		fcolor = NewFColor(100, 149, 237)
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
