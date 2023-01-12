package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := Point{0, 0, 0}
	b := Point{0, 0, 1}

	s := strconv.FormatFloat(a.DistanceToPoint(b), 'f', 2, 64)
	fmt.Println(s)
}
