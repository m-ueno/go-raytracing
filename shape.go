package main

import "math"

type Shape interface {
	testIntersection(*Ray) (*IntersectionPoint, bool)
	Material() *Material
}

type Sphere struct {
	center   *Vector
	radius   float64
	material *Material
}

type Plane struct {
	position *Vector
	normal   *Vector
	material *Material
}

func newSphere(center *Vector, radius float64, material *Material) *Sphere {
	return &Sphere{center, radius, material}
}

func (sp *Sphere) testIntersection(ray *Ray) (*IntersectionPoint, bool) {
	v := Sub(ray.start, sp.center) // v = from - center  : カメラから球の中心

	b := Dot(ray.direction, v) // b = to.dot(v)
	c := Dot(v, v) - sp.radius*sp.radius
	d := b*b - c

	if d < 0 { // 2次方程式が実数解を持たない
		return &IntersectionPoint{}, false
	}

	det := math.Sqrt(d)
	a := Dot(ray.direction, ray.direction)

	t1 := (-b - det) / a
	t2 := (-b + det) / a

	t := 0.0

	if t1 > 0 && t2 > 0 {
		t = math.Min(t1, t2)
	} else {
		t = math.Max(t1, t2)
	}

	if t >= 0 {
		// 視線ベクトルの延長線上に交点がある
		i_position := Add(ray.start, Scale(t, ray.direction)) // pos = A + c*B
		normal := Normalize(Sub(i_position, sp.center))

		return &IntersectionPoint{
			distance: t,
			position: i_position,
			normal:   normal,
		}, true
	} else {
		// 視線ベクトルから逆向き
		return &IntersectionPoint{}, false
	}
}

// Getter
func (sp *Sphere) Material() *Material {
	return sp.material
}

func newPlane(normal, position *Vector, material *Material) *Plane {
	return &Plane{
		normal:   normal,
		position: position,
		material: material,
	}
}

func (pl *Plane) testIntersection(ray *Ray) (*IntersectionPoint, bool) {
	s := ray.start
	d := ray.direction
	n := pl.normal

	cos := Dot(d, n)
	if cos == 0 {
		return nil, false
	} else {
		t := Dot(Add(Scale(-1.0, s), pl.position), n) / cos

		if t > 0 {
			return &IntersectionPoint{
				distance: t,
				normal:   n,
				position: Add(s, Scale(t, d)),
			}, true
		} else {
			return nil, false
		}
	}
}

func (pl *Plane) Material() *Material {
	return pl.material
}
