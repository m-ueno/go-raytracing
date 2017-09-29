package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"gonum.org/v1/gonum/mat"
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
	position *mat.VecDense
	normal   *mat.VecDense
}

// func newVector(data ...float64) *Vector { // error
func newVector(data ...float64) *mat.VecDense {
	return mat.NewVecDense(len(data), data)
}

func shading_ambient(shape *Sphere) *FColor {
	return NewFColor(30, 40, 50)
}

func rayTrace(ray *Ray, shape *Sphere) *FColor {
	/* 交差判定 */
	_, ok := shape.testIntersection(ray)

	fcolor := NewFColor(0, 0, 0)

	if ok {
		fcolor = add(fcolor, shading_ambient(shape)).(*FColor) // ださい
	} else {
		fcolor = NewFColor(100, 149, 237)
	}

	return fcolor
}

func render(size int) {
	from := newVector(0, 0, -5)

	// Sphere
	shape := NewSphere(newVector(0, 0, 5), 1.0)
	//	center := newVector(0, 0, 5)
	//	radius := 1.0

	fmt.Printf("P3\n%d %d\n255\n", size, size)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			screenXYZ := makeEye(x, y, size)

			// 視線方向
			to := mat.NewVecDense(3, nil)
			to.SubVec(screenXYZ, from)
			to.ScaleVec(1.0/mat.Norm(to, 2), to) // L2距離

			ray := NewRay(from, to)

			fcolor := rayTrace(ray, shape)
			fmt.Print(fcolor)
		}
	}
}

func makeEye(x int, y int, imageSize int) *mat.VecDense {
	return newVector(
		-1.0+float64(x)/float64(imageSize)*2,
		1.0-float64(y)/float64(imageSize)*2,
		0.0,
	)
}
