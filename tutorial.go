package main

func NewScene24(size int) *Scene {
	shapes := []Shape{
		NewSphere(
			NewVector(0, 0, 5),
			1.0,
			NewMaterial(),
		),
	}
	lightSources := []LightSource{
		NewPointLightSource(1.0, NewVector(-5, 5, -5)),
	}
	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func NewScene25(size int) *Scene {
	shapes := []Shape{
		NewSphere(
			NewVector(3, 0, 25),
			1.0,
			NewMaterial(),
		),
		NewSphere(
			NewVector(2, 0, 20),
			1.0,
			NewMaterial(),
		),
		NewSphere(
			NewVector(1, 0, 15),
			1.0,
			NewMaterial(),
		),
		NewSphere(
			NewVector(0, 0, 10),
			1.0,
			NewMaterial(),
		),
		NewSphere(
			NewVector(-1, 0, 5),
			1.0,
			NewMaterial(),
		),
		&Plane{
			normal:   NewVector(0, 1, 0),
			position: NewVector(0, -1, 0),
			material: NewMaterial(),
		},
	}
	lightSources := []LightSource{
		NewPointLightSource(1.0, NewVector(-5, 5, -5)),
	}
	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func NewScene27(size int) *Scene {
	sc := NewScene25(size)
	sc.lightSources = []LightSource{
		NewPointLightSource(0.5, NewVector(-5, 5, -5)),
		NewPointLightSource(0.5, NewVector(5, 0, -5)),
		NewPointLightSource(0.5, NewVector(5, 20, -5)),
	}
	return sc
}

func NewScene32(size int) *Scene {
	// 完全鏡面反射の部屋
	shapes := []Shape{
		&Sphere{
			center:   NewVector(-0.25, -0.5, 3),
			radius:   0.5,
			material: NewReflectMaterial(),
		},
		&Sphere{
			center:   NewVector(0.80, -0.5, 3),
			radius:   0.5,
			material: NewReflectMaterial(),
		},
		&Plane{ // 床
			material: NewMaterial(),
			normal:   NewVector(0, 1, 0),
			position: NewVector(0, -1, 0),
		},
		&Plane{ // 天井
			material: NewMaterial(),
			normal:   NewVector(0, -1, 0),
			position: NewVector(0, 1, 0),
		},

		&Plane{ // 左
			material: NewMaterial(),
			normal:   NewVector(1, 0, 0),
			position: NewVector(-1, 0, 0),
		},
		&Plane{ // 奥
			material: NewMaterial(),
			normal:   NewVector(0, 0, -1),
			position: NewVector(0, 0, 5),
		},
	}
	lightSources := []LightSource{
		NewPointLightSource(1.0, NewVector(0, 0.9, 2.5)),
	}

	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func NewColoredDiffuseMaterial(k_d *FColor) *Material {
	return &Material{
		shininess: 8,
		k_a:       NewFColor(0.1, 0.1, 0.1),
		k_d:       FCScale(0.7, k_d),
		k_s:       NewFColor(0, 0, 0),
		usePerfectReflectance: false,
		catadioptricFactor:    0.0,
	}
}

func NewScene32_2(size int) *Scene {
	// 完全鏡面反射の部屋
	shapes := []Shape{
		&Sphere{
			center:   NewVector(-0.25, -0.5, 3),
			radius:   0.5,
			material: NewReflectMaterial(),
		},
		&Plane{ // 床
			material: NewColoredDiffuseMaterial(NewFColor(1, 1, 1)),
			normal:   NewVector(0, 1, 0),
			position: NewVector(0, -1, 0),
		},
		&Plane{ // 天井
			material: NewColoredDiffuseMaterial(NewFColor(1, 1, 1)),
			normal:   NewVector(0, -1, 0),
			position: NewVector(0, 1, 0),
		},
		&Plane{ // 右
			material: NewColoredDiffuseMaterial(NewFColor(0, 1, 0)),
			normal:   NewVector(-1, 0, 0),
			position: NewVector(1, 0, 0),
		},
		&Plane{ // 左
			material: NewColoredDiffuseMaterial(NewFColor(1, 0, 0)),
			normal:   NewVector(1, 0, 0),
			position: NewVector(-1, 0, 0),
		},
		&Plane{ // 奥
			material: NewColoredDiffuseMaterial(NewFColor(1, 1, 1)),
			normal:   NewVector(0, 0, -1),
			position: NewVector(0, 0, 5),
		},
	}
	lightSources := []LightSource{
		NewPointLightSource(1.0, NewVector(0, 0.9, 2.5)),
	}

	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}
