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

func TestHaversine(t *testing.T) {
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
			eu:    424.26356829147227,
			as:    9623.886715873092,
			us:    9109.26204607284,
		},
		{
			Point: Point{Latitude: 23.31667, Longitude: 90.21667},
			city:  "Bogra",
			eu:    7675.604417002523,
			as:    2458.4981174464233,
			us:    12976.41077321424,
		},
		{
			Point: Point{Latitude: 10.498426848097456, Longitude: -66.87011352123069},
			city:  "Caracas",
			eu:    7858.884945,
			as:    16384.314714,
			us:    5829.716592,
		},
	}
	const prec = 1 / 1e6
	for _, st := range stages {
		t.Run(st.city, func(t *testing.T) {
			distEU := Haversine(st.Point, refPoints[0])
			if km := distEU.Kilometers(); math.Abs(km-st.eu) > prec {
				t.Errorf("EU distance mismatch: %f vs %f", km, st.eu)
			}

			distHK := Haversine(st.Point, refPoints[1])
			if km := distHK.Kilometers(); math.Abs(km-st.as) > prec {
				t.Errorf("AS distance mismatch: %f vs %f", km, st.as)
			}

			distUS := Haversine(st.Point, refPoints[2])
			if km := distUS.Kilometers(); math.Abs(km-st.us) > prec {
				t.Errorf("US distance mismatch: %f vs %f", km, st.us)
			}
		})
	}
}

func BenchmarkHaversine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Haversine(Point{Latitude: 48.85002, Longitude: 2.600178}, refPoints[0])
	}
}
