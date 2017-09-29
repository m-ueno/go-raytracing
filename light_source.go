package main

type Lighting struct {
	direction *Vector
	distance  float64
	intensity float64
}

type LightSource interface {
	LightingAt(*Vector) *Lighting
}

type PointLightSource struct {
	intensity float64
	position  *Vector
}

func NewLighting(direction *Vector, distance float64, intensity float64) *Lighting {
	return &Lighting{
		direction: direction,
		distance:  distance,
		intensity: intensity,
	}
}

func NewPointLightSource(intensity float64, position *Vector) *PointLightSource {
	return &PointLightSource{
		intensity: intensity,
		position: position,
	}
}

func (pls *PointLightSource) LightingAt(p *Vector) *Lighting {
	return NewLighting(
		Normalize(Sub(p, pls.position)),
		Norm(Sub(p, pls.position)),
		pls.intensity,
	)
}
