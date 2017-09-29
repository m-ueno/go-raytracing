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

func (fc *FColor) String() string {
	// r, g, b := fc.x, fc.y, fc.z
	r := int(math.Ceil(fc.x))
	g := int(math.Ceil(fc.y))
	b := int(math.Ceil(fc.z))
	return fmt.Sprintf("%d %d %d ", r, g, b)
}

func NewFColor(r, g, b float64) *FColor {
	v := NewVector(r, g, b)
	return &FColor{v}
}
