package main

import (
	"fmt"
	"math"
)

type FColor struct {
	*Vector
}

func FCAdd(a, b *FColor) *FColor {
	return NewFColor(
		a.x+b.x,
		a.y+b.y,
		a.z+b.z,
	)
}

func FCScale(alpha float64, a *FColor) *FColor {
	return NewFColor(
		a.x*alpha,
		a.y*alpha,
		a.z*alpha,
	)
}

func (fc *FColor) String() string {
	// r, g, b := fc.x, fc.y, fc.z
	r := int(math.Min(math.Ceil(fc.x*255), 255))
	g := int(math.Min(math.Ceil(fc.y*255), 255))
	b := int(math.Min(math.Ceil(fc.z*255), 255))
	return fmt.Sprintf("%d %d %d ", r, g, b)
}

func NewFColor(r, g, b float64) *FColor {
	v := NewVector(r, g, b)
	return &FColor{v}
}
