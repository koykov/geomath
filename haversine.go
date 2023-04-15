package geomath

import "math"

// DistanceHaversine returns distance between a and b calculated by haversine.
// See https://en.wikipedia.org/wiki/Haversine_formula
func DistanceHaversine(a, b Point) Distance {
	dlat := (b.Latitude - a.Latitude) * deg2rad
	dlon := (b.Longitude - a.Longitude) * deg2rad
	x := math.Pow(math.Sin(dlat/2), 2) + math.Cos(a.Latitude*deg2rad)*math.Cos(b.Latitude*deg2rad)*math.Pow(math.Sin(dlon/2), 2)
	y := 2 * math.Atan2(math.Sqrt(x), math.Sqrt(1-x))
	return Distance(earthRadius * y)
}
