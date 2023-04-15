# Geo math

Geographical math toolset.

Features:
* distance calculation
* find the nearest pointer

### Distance calculation

```go
p0 := Point{Latitude: 48.85002, Longitude: 2.600178} // Paris
p1 := Point{Latitude: 52.3740, Longitude: 4.8897} // Amsterdam
dist := DistanceHaversine(p0, p1)
println(dist.Kilometers()) // +4.242636e+002
```

Also, you may use distance calculation based on [spherical law of cosines](https://en.wikipedia.org/wiki/Spherical_law_of_cosines)
by calling function `DistanceSLC`. In example above uses function based on [haversine](https://en.wikipedia.org/wiki/Haversine_formula) formula.

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

By default, method `NearestOf` uses haversine method. You may use [SLC](https://en.wikipedia.org/wiki/Spherical_law_of_cosines)
by calling method `NearestOfWithOptions`.
