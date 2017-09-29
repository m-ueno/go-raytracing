package main

type Ray struct {
	start *Vector
	direction *Vector
}

func NewRay(start, direction *Vector) *Ray {
	return &Ray{start, direction}
}
