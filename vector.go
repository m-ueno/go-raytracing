package main

import "gonum.org/v1/gonum/mat"

func normalize(v *mat.VecDense) *mat.VecDense {
	vv := newVector(0, 0, 0)
	norm := mat.Norm(v, 2)
	vv.ScaleVec(1.0/norm, v)
	return vv
}

func add(a, b *mat.VecDense) *mat.VecDense {
	v := newVector(0, 0, 0)
	v.AddVec(a, b)

	return v
}

func sub(a, b *mat.VecDense) *mat.VecDense {
	v := newVector(0, 0, 0)
	v.SubVec(a, b)
	return v
}

func scale(alpha float64, a *mat.VecDense) *mat.VecDense {
	v := newVector(0, 0, 0)
	v.ScaleVec(alpha, a)
	return v
}
