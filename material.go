package main

// Material is reflection information
type Material struct {
	shininess             float64
	k_a                   *fColor
	k_d                   *fColor
	k_s                   *fColor
	usePerfectReflectance bool
	catadioptricFactor    float64 //*fColor
	useRefraction         bool    // 屈折を使用するかどうか
	refractionIndex       float64 // 絶対屈折率
}

func newMaterial() *Material {
	return &Material{
		shininess: 8,
		k_a:       newfColor(0.01, 0.01, 0.01),
		k_d:       newfColor(0.7, 0.5, 0.3),
		k_s:       newfColor(0.3, 0.3, 0.3),
		usePerfectReflectance: false,
		catadioptricFactor:    0.0,
	}
}

func newReflectMaterial() *Material {
	return &Material{
		shininess: 8,
		k_a:       newfColor(0, 0, 0),
		k_d:       newfColor(0, 0, 0),
		k_s:       newfColor(0, 0, 0),
		usePerfectReflectance: true,
		catadioptricFactor:    1.0,
	}
}
