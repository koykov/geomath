package geomath

import "math"

type Point struct {
	Latitude, Longitude float64
}

func (p Point) NearestOf(points ...Point) Point {
	p1, _, _ := p.NearestOfWithDetails(points...)
	return p1
}

func (p Point) NearestOfWithDetails(points ...Point) (Point, int, Distance) {
	if len(points) == 0 {
		return Point{}, -1, 0
	}
	var (
		md = Distance(math.MaxFloat64)
		mi int
	)
	for i := 0; i < len(points); i++ {
		dist := DistanceHaversine(p, points[i])
		if dist < md {
			md = dist
			mi = i
		}
	}
	return points[mi], mi, md
}
