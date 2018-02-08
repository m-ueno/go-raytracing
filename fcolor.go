package raytracing

import (
	"fmt"
	"math"
)

type fColor struct {
	*Vector
}

func fCAdd(a, b *fColor) *fColor {
	return newfColor(
		a.x+b.x,
		a.y+b.y,
		a.z+b.z,
	)
}

func fCScale(alpha float64, a *fColor) *fColor {
	return newfColor(
		a.x*alpha,
		a.y*alpha,
		a.z*alpha,
	)
}

func (fc *fColor) String() string {
	// r, g, b := fc.x, fc.y, fc.z
	r := int(math.Min(math.Ceil(fc.x*255), 255))
	g := int(math.Min(math.Ceil(fc.y*255), 255))
	b := int(math.Min(math.Ceil(fc.z*255), 255))
	return fmt.Sprintf("%d %d %d ", r, g, b)
}

func newfColor(r, g, b float64) *fColor {
	v := newVector(r, g, b)
	return &fColor{v}
}
