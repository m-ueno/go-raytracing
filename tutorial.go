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
