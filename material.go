package main

type Material struct {
	shininess float64
	k_a       *FColor
	k_d       *FColor
	k_s       *FColor
}

func NewMaterial() *Material {
	return &Material{
		shininess: 8,
		k_a:       NewFColor(0.01, 0.01, 0.01),
		k_d:       NewFColor(0.7, 0.5, 0.3),
		k_s:       NewFColor(0.3, 0.3, 0.3),
	}
}
