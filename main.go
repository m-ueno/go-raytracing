package main

import (
	"flag"
	"log"
	"strconv"
)

const EPSILON = 1.0 / 128

func main() {
	log.Println("start")

	flag.Parse()

	args := flag.Args()
	size, _ := strconv.Atoi(args[0])

	scene := NewScene33(size)
	scene.render(false)

	log.Println("end")
}

type IntersectionPoint struct {
	distance float64
	position *Vector
	normal   *Vector
}

type IntersectionTestResult struct {
	intersectionPoint *IntersectionPoint
	shape             Shape
}

func makeEye(x int, y int, imageSize int) *Vector {
	return makeEyeWithSampling(x, y, imageSize, 0, 0)
}

func makeEyeWithSampling(x, y int, imageSize int, dx, dy float64) *Vector {
	return NewVector(
		-1.0+(float64(x)+dx)/float64(imageSize)*2,
		1.0-(float64(y)+dy)/float64(imageSize)*2,
		0.0,
	)
}
