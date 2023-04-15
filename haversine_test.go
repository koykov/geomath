package geomath

import (
	"math"
	"testing"
)

var refPoints = []Point{
	{Latitude: 52.3740, Longitude: 4.8897},    // Europe/Amsterdam
	{Latitude: 22.3027, Longitude: 114.1772},  // Asia/HongKong
	{Latitude: 34.0522, Longitude: -118.2436}, // US/LA
}

func TestDistanceHaversine(t *testing.T) {
	type stage struct {
		Point
		city string
		eu   float64
		as   float64
		us   float64
	}
	var stages = []stage{
		{
			Point: Point{Latitude: 48.85002, Longitude: 2.600178},
			city:  "Paris",
			eu:    424.263655,
			as:    9623.887061,
			us:    9109.261973,
		},
		{
			Point: Point{Latitude: 23.31667, Longitude: 90.21667},
			city:  "Bogra",
			eu:    7675.604633,
			as:    2458.497796,
			us:    12976.410328,
		},
		{
			Point: Point{Latitude: 10.498426848097456, Longitude: -66.87011352123069},
			city:  "Caracas",
			eu:    7858.884631,
			as:    16384.314701,
			us:    5829.716920,
		},
	}
	const prec = 1 / 1e6
	for _, st := range stages {
		t.Run(st.city, func(t *testing.T) {
			distEU := DistanceHaversine(st.Point, refPoints[0])
			if km := distEU.Kilometers(); math.Abs(km-st.eu) > prec {
				t.Errorf("EU distance mismatch: %f vs %f", km, st.eu)
			}

			distHK := DistanceHaversine(st.Point, refPoints[1])
			if km := distHK.Kilometers(); math.Abs(km-st.as) > prec {
				t.Errorf("AS distance mismatch: %f vs %f", km, st.as)
			}

			distUS := DistanceHaversine(st.Point, refPoints[2])
			if km := distUS.Kilometers(); math.Abs(km-st.us) > prec {
				t.Errorf("US distance mismatch: %f vs %f", km, st.us)
			}
		})
	}
}

func BenchmarkDistanceHaversine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DistanceHaversine(Point{Latitude: 48.85002, Longitude: 2.600178}, refPoints[0])
	}
}
