package geomath

import "math"

const (
	earthRadius = 6378137000 // unit: mm
	deg2rad     = math.Pi / 180
)

// Haversine returns distance between a and b.
// See https://en.wikipedia.org/wiki/Haversine_formula
func Haversine(a, b Point) Distance {
	dlat := float64(b.Latitude-a.Latitude) * deg2rad
	dlon := float64(b.Longitude-a.Longitude) * deg2rad
	x := math.Pow(math.Sin(dlat/2), 2) + math.Cos(float64(a.Latitude*deg2rad))*math.Cos(float64(b.Latitude*deg2rad))*math.Pow(math.Sin(dlon/2), 2)
	y := 2 * math.Atan2(math.Sqrt(x), math.Sqrt(1-x))
	return Distance(earthRadius * y)
}
