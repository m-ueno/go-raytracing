package main

import (
	"flag"
	"log"
	"strconv"
)

func main() {
	log.Println("hello")

	flag.Parse()

	args := flag.Args()
	size, _ := strconv.Atoi(args[0])

	scene := NewScene24(size)
	scene.render()

	log.Println("bye")
}

type IntersectionPoint struct {
	distance float64
	position *Vector
	normal   *Vector
}

func makeEye(x int, y int, imageSize int) *Vector {
	return NewVector(
		-1.0+float64(x)/float64(imageSize)*2,
		1.0-float64(y)/float64(imageSize)*2,
		0.0,
	)
}
