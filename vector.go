package raytracing

import (
	"math"
)

// Vector represents 3d vector
type Vector struct {
	x, y, z float64
}

func newVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

// Add returns pointer of v1 + v2
func Add(a, b *Vector) *Vector {
	return newVector(
		a.x+b.x,
		a.y+b.y,
		a.z+b.z,
	)
}

// Sub returns pointer of v1 - v2
func Sub(a, b *Vector) *Vector {
	return newVector(
		a.x-b.x,
		a.y-b.y,
		a.z-b.z,
	)
}

// Scale returns pointer of scaled vector
func Scale(alpha float64, a *Vector) *Vector {
	return newVector(
		alpha*a.x,
		alpha*a.y,
		alpha*a.z,
	)
}

// Dot returns the sum of the element-wise product of a and b
func Dot(a, b *Vector) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// Norm returns square root of the sum of the squares of the elements
func Norm(a *Vector) float64 {
	return math.Sqrt(Dot(a, a)) // sqrt(||a||^2)
}

// Normalize returns pointer of normalized vector
func Normalize(a *Vector) *Vector {
	normInverse := 1.0 / Norm(a)
	return Scale(normInverse, a)
}
