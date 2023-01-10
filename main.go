package main

func main() {
	a := Point{1, 2, 3}
	b := Point{1, 2, -4}

	print(a.DistanceToPoint(b))
	print(a.DistanceToCordination(1, 2, 3))
}
