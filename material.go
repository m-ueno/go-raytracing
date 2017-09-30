package main

type Material struct {
	shininess             float64
	k_a                   *FColor
	k_d                   *FColor
	k_s                   *FColor
	usePerfectReflectance bool
	catadioptricFactor    float64 //*FColor
	useRefraction         bool    // 屈折を使用するかどうか
	refractionIndex       float64 // 絶対屈折率
}

func NewMaterial() *Material {
	return &Material{
		shininess: 8,
		k_a:       NewFColor(0.01, 0.01, 0.01),
		k_d:       NewFColor(0.7, 0.5, 0.3),
		k_s:       NewFColor(0.3, 0.3, 0.3),
		usePerfectReflectance: false,
		catadioptricFactor:    0.0,
	}
}

func NewReflectMaterial() *Material {
	return &Material{
		shininess: 8,
		k_a:       NewFColor(0, 0, 0),
		k_d:       NewFColor(0, 0, 0),
		k_s:       NewFColor(0, 0, 0),
		usePerfectReflectance: true,
		catadioptricFactor:    1.0,
	}
}
