package geomath

import (
	"math"
	"strconv"
	"testing"
)

func TestDistanceSLC(t *testing.T) {
	type stage struct {
		Point
		exp float64
	}
	var stages = []stage{
		{
			Point: Point{Latitude: 47.457131, Longitude: -10.615383},
			exp:   1112.8102127620043,
		},
		{
			Point: Point{Latitude: 61.265367, Longitude: 19.939854},
			exp:   1508.5367561512999,
		},
		{
			Point: Point{Latitude: 56.64524, Longitude: -21.312707},
			exp:   1745.5623350630779,
		},
		{
			Point: Point{Latitude: 37.028064, Longitude: 18.732122},
			exp:   1957.5437654460789,
		},
		{
			Point: Point{Latitude: 38.823754, Longitude: -17.029437},
			exp:   2112.0358748853623,
		},
		{
			Point: Point{Latitude: 70.687405, Longitude: 11.08298},
			exp:   2217.1191451897485,
		},
	}
	ref := Point{Latitude: 51.0363432, Longitude: 3.7351858}
	const prec = 1 / 1e6
	for i, st := range stages {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			dist := DistanceSLC(st.Point, ref)
			if km := dist.Kilometers(); math.Abs(km-st.exp) > prec {
				t.Errorf("distance mismatch: %f vs %f", km, st.exp)
			}
		})
	}
}

func BenchmarkDistanceSLC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DistanceSLC(Point{Latitude: 47.457131, Longitude: -10.615383}, Point{Latitude: 51.0363432, Longitude: 3.7351858})
	}
}
