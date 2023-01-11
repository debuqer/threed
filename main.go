package main

func main() {
	a := Point{0, 0, 0}
	b := Point{1, 1, 1}

	print(a.DistanceToPoint(b))
	print(a.DistanceToCordination(1, 2, 3))
}
