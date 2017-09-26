package main

import "gonum.org/v1/gonum/mat"

type Ray struct {
	start *mat.VecDense
	direction *mat.VecDense
}

func NewRay(start, direction *mat.VecDense) *Ray {
	return &Ray{start, direction}
}
