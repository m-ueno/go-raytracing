package raytracing

func newScene24(size int) *Scene {
	shapes := []Shape{
		newSphere(
			newVector(0, 0, 5),
			1.0,
			newMaterial(),
		),
	}
	lightSources := []LightSource{
		newPointLightSource(1.0, newVector(-5, 5, -5)),
	}
	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func newScene25(size int) *Scene {
	shapes := []Shape{
		newSphere(
			newVector(3, 0, 25),
			1.0,
			newMaterial(),
		),
		newSphere(
			newVector(2, 0, 20),
			1.0,
			newMaterial(),
		),
		newSphere(
			newVector(1, 0, 15),
			1.0,
			newMaterial(),
		),
		newSphere(
			newVector(0, 0, 10),
			1.0,
			newMaterial(),
		),
		newSphere(
			newVector(-1, 0, 5),
			1.0,
			newMaterial(),
		),
		&Plane{
			normal:   newVector(0, 1, 0),
			position: newVector(0, -1, 0),
			material: newMaterial(),
		},
	}
	lightSources := []LightSource{
		newPointLightSource(1.0, newVector(-5, 5, -5)),
	}
	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func newScene27(size int) *Scene {
	sc := newScene25(size)
	sc.lightSources = []LightSource{
		newPointLightSource(0.5, newVector(-5, 5, -5)),
		newPointLightSource(0.5, newVector(5, 0, -5)),
		newPointLightSource(0.5, newVector(5, 20, -5)),
	}
	return sc
}

func newScene32(size int) *Scene {
	// 完全鏡面反射の部屋
	shapes := []Shape{
		&Sphere{
			center:   newVector(-0.25, -0.5, 3),
			radius:   0.5,
			material: newReflectMaterial(),
		},
		&Sphere{
			center:   newVector(0.80, -0.5, 3),
			radius:   0.5,
			material: newReflectMaterial(),
		},
		&Plane{ // 床
			material: newMaterial(),
			normal:   newVector(0, 1, 0),
			position: newVector(0, -1, 0),
		},
		&Plane{ // 天井
			material: newMaterial(),
			normal:   newVector(0, -1, 0),
			position: newVector(0, 1, 0),
		},

		&Plane{ // 左
			material: newMaterial(),
			normal:   newVector(1, 0, 0),
			position: newVector(-1, 0, 0),
		},
		&Plane{ // 奥
			material: newMaterial(),
			normal:   newVector(0, 0, -1),
			position: newVector(0, 0, 5),
		},
	}
	lightSources := []LightSource{
		newPointLightSource(1.0, newVector(0, 0.9, 2.5)),
	}

	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func newColoredDiffuseMaterial(k_d *fColor) *Material {
	return &Material{
		shininess: 8,
		k_a:       newfColor(0.1, 0.1, 0.1),
		k_d:       fCScale(0.7, k_d),
		k_s:       newfColor(0, 0, 0),
		usePerfectReflectance: false,
		catadioptricFactor:    0.0,
	}
}

func newScene32_2(size int) *Scene {
	// 完全鏡面反射の部屋
	shapes := []Shape{
		&Sphere{
			center:   newVector(-0.25, -0.5, 3),
			radius:   0.5,
			material: newReflectMaterial(),
		},
		&Plane{ // 床
			material: newColoredDiffuseMaterial(newfColor(1, 1, 1)),
			normal:   newVector(0, 1, 0),
			position: newVector(0, -1, 0),
		},
		&Plane{ // 天井
			material: newColoredDiffuseMaterial(newfColor(1, 1, 1)),
			normal:   newVector(0, -1, 0),
			position: newVector(0, 1, 0),
		},
		&Plane{ // 右
			material: newColoredDiffuseMaterial(newfColor(0, 1, 0)),
			normal:   newVector(-1, 0, 0),
			position: newVector(1, 0, 0),
		},
		&Plane{ // 左
			material: newColoredDiffuseMaterial(newfColor(1, 0, 0)),
			normal:   newVector(1, 0, 0),
			position: newVector(-1, 0, 0),
		},
		&Plane{ // 奥
			material: newColoredDiffuseMaterial(newfColor(1, 1, 1)),
			normal:   newVector(0, 0, -1),
			position: newVector(0, 0, 5),
		},
	}
	lightSources := []LightSource{
		newPointLightSource(1.0, newVector(0, 0.9, 2.5)),
	}

	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
	}
}

func NewScene33(size int) *Scene {
	// 完全鏡面反射の部屋
	shapes := []Shape{
		&Sphere{
			center:   newVector(-0.4, -0.65, 3),
			radius:   0.35,
			material: newReflectMaterial(),
		},
		&Sphere{
			center: newVector(0.5, -0.65, 2),
			radius: 0.35,
			material: &Material{
				k_a:                newfColor(0, 0, 0),
				k_d:                newfColor(0, 0, 0),
				k_s:                newfColor(0, 0, 0),
				useRefraction:      true,
				catadioptricFactor: 1.0,
				refractionIndex:    1.51,
			},
		},
		&Plane{ // 床
			material: newColoredDiffuseMaterial(newfColor(1, 1, 1)),
			normal:   newVector(0, 1, 0),
			position: newVector(0, -1, 0),
		},
		&Plane{ // 天井
			material: newColoredDiffuseMaterial(newfColor(1, 1, 1)),
			normal:   newVector(0, -1, 0),
			position: newVector(0, 1, 0),
		},
		&Plane{ // 右
			material: newColoredDiffuseMaterial(newfColor(0, 1, 0)),
			normal:   newVector(-1, 0, 0),
			position: newVector(1, 0, 0),
		},
		&Plane{ // 左
			material: newColoredDiffuseMaterial(newfColor(1, 0, 0)),
			normal:   newVector(1, 0, 0),
			position: newVector(-1, 0, 0),
		},
		&Plane{ // 奥
			material: newColoredDiffuseMaterial(newfColor(1, 1, 1)),
			normal:   newVector(0, 0, -1),
			position: newVector(0, 0, 5),
		},
	}
	lightSources := []LightSource{
		newPointLightSource(1.0, newVector(0, 0.9, 2.5)),
	}

	ambientIntensity := 0.1

	return &Scene{
		shapes:           shapes,
		lightSources:     lightSources,
		ambientIntensity: ambientIntensity,
		size:             size,
		globalRefractionIndex: 1.0,
	}
}
