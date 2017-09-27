package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Shape interface {
	testIntersection()
}

type Sphere struct {
	center *mat.VecDense
	radius float64
}

func NewSphere(center *mat.VecDense, radius float64) *Sphere {
	return &Sphere{center, radius}
}

func (sp *Sphere) testIntersection(ray *Ray) (*IntersectionPoint, bool) {
	v := mat.NewVecDense(3, nil)
	v.SubVec(ray.start, sp.center) // v = from - center  : カメラから球の中心

	b := mat.Dot(ray.direction, v) // b = to.dot(v)
	c := mat.Dot(v, v) - sp.radius*sp.radius
	d := b*b - c

	if d < 0 { // 2次方程式が実数解を持たない
		return &IntersectionPoint{}, false
	}

	det := math.Sqrt(d)
	a := mat.Dot(ray.direction, ray.direction) // 大きさ ほぼ1では？

	t1 := (-b - det) / a
	t2 := (-b + det) / a

	t := 0.0

	if t1 > 0 && t2 > 0 {
		t = math.Min(t1, t2)
	} else {
		t = math.Max(t1, t2)
	}

	if t < 0 { // 視線ベクトルから逆向き
		return &IntersectionPoint{}, false
	}

	// t>=0 なら交差ある
	i_position := newVector(0, 0, 0)
	i_position.AddScaledVec(ray.start, t, ray.direction) // pos = A + c*B
	normal := normalize(sub(i_position, sp.center))

	return &IntersectionPoint{
		distance: t,
		position: i_position,
		normal:   normal,
	}, true
}
