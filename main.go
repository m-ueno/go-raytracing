package main

import (
	"flag"
	"log"
	"os"
)

// EPSILON is used to calcurate reflection
const EPSILON = 1.0 / 128

func main() {
	var antialiasing = flag.Bool("aa", false, "Enable antialiasing (slow)")
	var output = flag.String("output", "", "Output filepath")
	var size = flag.Int("size", 0, "Image size in pixels")

	flag.Parse()
	if *size == 0 {
		log.Fatalln("Please provide `-size <size>`")
		os.Exit(1)
	}

	if *output == "" {
		log.Fatalln("Please provide `-output <path>`")
		os.Exit(1)
	}

	log.Println("start")
	log.Println("antialiasing:", *antialiasing)
	log.Println("output:", *output)
	log.Println("size:", *size)

	f, err := os.OpenFile(*output, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer f.Close()

	scene := newScene33(*size)
	scene.render(*antialiasing, f)

	log.Println("end")
}

// IntersectionPoint is a point where sight vector intersect surface
type IntersectionPoint struct {
	distance float64
	position *Vector
	normal   *Vector
}

// IntersectionTestResult is returned by testIntersectionWithAll()
// intersectionPoint is nil when no intersection found
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
