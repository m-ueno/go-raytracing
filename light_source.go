package main

// Lighting produce light ray
type Lighting struct {
	direction *Vector
	distance  float64
	intensity float64
}

// LightSource produce Lighting
type LightSource interface {
	lightingAt(*Vector) *Lighting
}

// PointLightSource is a light source
type PointLightSource struct {
	intensity float64
	position  *Vector
}

func newLighting(direction *Vector, distance float64, intensity float64) *Lighting {
	return &Lighting{
		direction: direction,
		distance:  distance,
		intensity: intensity,
	}
}

func newPointLightSource(intensity float64, position *Vector) *PointLightSource {
	return &PointLightSource{
		intensity: intensity,
		position:  position,
	}
}

func (pls *PointLightSource) lightingAt(p *Vector) *Lighting {
	return newLighting(
		Normalize(Sub(p, pls.position)),
		Norm(Sub(p, pls.position)),
		pls.intensity,
	)
}
