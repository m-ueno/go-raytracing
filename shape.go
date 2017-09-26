package main

import "gonum.org/v1/gonum/mat"

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

func (sp *Sphere) testIntersection(r *Ray) (*IntersectionPoint, bool) {
	v := mat.NewVecDense(3, nil)
	v.SubVec(sp.center, r.start) // v = from - center  : カメラから球の中心

	b := mat.Dot(r.direction, v) // b = to.dot(v)
	c := mat.Dot(v, v) - sp.radius*sp.radius
	d := b*b - c

	return &IntersectionPoint{}, d >= 0
}
