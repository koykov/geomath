package geomath

import "testing"

func TestPoint(t *testing.T) {
	type stage struct {
		city  string
		point Point
		idx   int
	}
	var stages = []stage{
		{"Paris", Point{Latitude: 48.85002, Longitude: 2.600178}, 0},
		{"Bogra", Point{Latitude: 23.31667, Longitude: 90.21667}, 1},
		{"Caracas", Point{Latitude: 10.498426848097456, Longitude: -66.87011352123069}, 2},
	}
	for _, st := range stages {
		t.Run(st.city, func(t *testing.T) {
			_, idx, _ := st.point.NearestOfWithDetails(refPoints...)
			if idx != st.idx {
				t.Errorf("nearest point mismatch: %d vs %d", idx, st.idx)
			}
		})
	}
}

func BenchmarkPoint(b *testing.B) {
	p := Point{Latitude: 48.85002, Longitude: 2.600178}
	for i := 0; i < b.N; i++ {
		_ = p.NearestOf(refPoints...)
	}
}
