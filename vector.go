package main

import (
	"math"
)

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func Add(a, b *Vector) *Vector {
	return NewVector(
		a.x+b.x,
		a.y+b.y,
		a.z+b.z,
	)
}

func Sub(a, b *Vector) *Vector {
	return NewVector(
		a.x-b.x,
		a.y-b.y,
		a.z-b.z,
	)
}

func Scale(alpha float64, a *Vector) *Vector {
	return NewVector(
		alpha*a.x,
		alpha*a.y,
		alpha*a.z,
	)
}

func Dot(a, b *Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Norm(a *Vector) float64 {
	return math.Sqrt(Dot(a, a)) // sqrt(||a||^2)
}

func Normalize(a *Vector) *Vector {
	norm_inverse := 1.0 / Norm(a)
	return Scale(norm_inverse, a)
}
