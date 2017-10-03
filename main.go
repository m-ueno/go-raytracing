package main

import (
	"flag"
	"log"
)

const EPSILON = 1.0 / 128

func main() {
	var size = flag.Int("size", 0, "Image size in pixels")
	var antialiasing = flag.Bool("aa", false, "Enable antialiasing (slow)")

	flag.Parse()
	if *size == 0 {
		log.Fatalln("Please provide `-size <size>`")
		return
	}

	log.Println("start")
	log.Println("antialiasing:", *antialiasing)
	log.Println("size:", *size)

	scene := newScene33(*size)
	scene.render(*antialiasing)

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
	return newVector(
		-1.0+(float64(x)+dx)/float64(imageSize)*2,
		1.0-(float64(y)+dy)/float64(imageSize)*2,
		0.0,
	)
}
