# Geo math

Geographical math toolset.

Features:
* distance calculation
* find the nearest pointer

### Distance calculation

```go
p0 := Point{Latitude: 48.85002, Longitude: 2.600178} // Paris
p1 := Point{Latitude: 52.3740, Longitude: 4.8897} // Amsterdam
dist := Haversine(p0, p1)
println(dist.Kilometers()) // +4.242636e+002
```

### Find the nearest point

```go
p := Point{Latitude: 10.498426848097456, Longitude: -66.87011352123069} // Caracas
ps := []Point{
    {Latitude: 52.3740, Longitude: 4.8897},    // Europe/Amsterdam
    {Latitude: 22.3027, Longitude: 114.1772},  // Asia/HongKong
    {Latitude: 34.0522, Longitude: -118.2436}, // US/LA
}
n := p.NearestOf(ps...)
println(n.Latitude, n.Longitude) // +3.405220e+001 -1.182436e+002
```
