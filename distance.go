package geomath

type Distance float64

func (d Distance) Millimeters() float64 {
	return float64(d)
}

func (d Distance) Centimeters() float64 {
	return float64(d / 10)
}

func (d Distance) Decimeters() float64 {
	return float64(d / 100)
}

func (d Distance) Meters() float64 {
	return float64(d / 1000)
}

func (d Distance) Kilometers() float64 {
	return float64(d / 1000000)
}
