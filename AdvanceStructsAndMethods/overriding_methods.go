package main

import (
	f "fmt"
	m "math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return m.Sqrt(p.x*p.x + p.y*p.y)
}

func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}

type NamedPoint struct {
	Point
	name string
}

func main() {
	n := NamedPoint{Point{3, 4}, "Pythogoras"}
	f.Println("Calling Point Abs: ", n.Point.Abs())
	f.Println(n.Abs())
}
