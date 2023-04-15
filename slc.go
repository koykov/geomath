package geomath

import "math"

// DistanceSLC returns distance between a and b calculated by spherical law of cosines.
// See https://en.wikipedia.org/wiki/Spherical_law_of_cosines
func DistanceSLC(a, b Point) Distance {
	theta := a.Longitude - b.Longitude
	dist := (math.Sin(a.Latitude*deg2rad) *
		math.Sin(b.Latitude*deg2rad)) +
		(math.Cos(a.Latitude*deg2rad) *
			math.Cos(b.Latitude*deg2rad) *
			math.Cos(theta*deg2rad))
	dist = math.Acos(dist)
	dist = dist * rad2deg
	dist = dist * 60 * 1.853159616
	return Distance(dist) * 1e6
}
