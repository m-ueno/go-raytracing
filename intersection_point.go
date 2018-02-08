package raytracing

// EPSILON is used to calcurate reflection
const EPSILON = 1.0 / 128

// IntersectionPoint is a point where sight vector intersect surface
type IntersectionPoint struct {
	distance float64
	position *Vector
	normal   *Vector
}

// IntersectionTestResult is returned by testIntersectionWithAll()
// intersectionPoint is nil when no intersection found
type IntersectionTestResult struct {
	intersectionPoint *IntersectionPoint
	shape             Shape
}
