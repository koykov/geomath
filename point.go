package geomath

import "math"

type Point struct {
	Latitude, Longitude float64
}

func (p Point) NearestOf(points ...Point) Point {
	p1, _, _ := p.NearestOfWithOptions(Haversine, points...)
	return p1
}

func (p Point) NearestOfWithOptions(method DistanceCalculation, points ...Point) (Point, int, Distance) {
	if len(points) == 0 {
		return Point{}, -1, 0
	}
	var (
		md = Distance(math.MaxFloat64)
		mi int
		fn func(Point, Point) Distance
	)
	switch method {
	case Haversine:
		fn = DistanceHaversine
	case SLC:
		fn = DistanceSLC
	default:
		return Point{}, -1, 0
	}
	for i := 0; i < len(points); i++ {
		dist := fn(p, points[i])
		if dist < md {
			md = dist
			mi = i
		}
	}
	return points[mi], mi, md
}
