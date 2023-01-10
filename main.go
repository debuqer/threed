package main

import "math"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (a *Point) DistanceToCordination(x float64, y float64, z float64) float64 {
	return math.Sqrt(math.Pow(a.X-x, 2) + math.Pow(a.Y-y, 2) + math.Pow(a.Z-z, 2))
}

func (a *Point) DistanceToPoint(b Point) float64 {
	return a.DistanceToCordination(b.X, b.Y, b.Z)
}

func main() {
	a := Point{1, 2, 3}
	b := Point{1, 2, -4}

	print(a.DistanceToPoint(b))
}
