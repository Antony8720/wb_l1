package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point{
	return &Point{x: x, y: y}
}

func solution(p1, p2 *Point) float64{
	var dx float64
	var dy float64
	if p2.x > p1.x {
		dx = math.Pow(p2.x - p1.x, 2)
	} else{
		dx = math.Pow(p1.x - p2.x, 2)
	}
	if p2.y > p1.y {
		dy = math.Pow(p2.y-p1.y, 2)
	} else {
		dy = math.Pow(p1.y-p2.y, 2)
	}
	return math.Sqrt(dx + dy)
}

func main() {
	p1 := NewPoint(2, 2)
	p2 := NewPoint(1, 1)
	fmt.Println(solution(p1, p2))
}