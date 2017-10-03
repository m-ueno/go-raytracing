package main

// Ray is half line
type Ray struct {
	start     *Vector
	direction *Vector
}

func newRay(start, direction *Vector) *Ray {
	return &Ray{start, direction}
}
