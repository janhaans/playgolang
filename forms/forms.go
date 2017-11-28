package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius float64
}

func (c *Circle) Surface() float64 {
	return math.Pow(c.Radius, 2.0) * math.Pi
}

type Wheel struct {
	Circle
	Spokes int
}

type GetSurface interface {
	Surface() float64
}

func main() {
	aCircle := Circle{Point{1, 2}, 2.0}
	aWheel := Wheel{Circle{Point{1, 2}, 2.0}, 8}
	forms := map[string]GetSurface{"aCircle": &aCircle, "aWheel": &aWheel}
	for k, v := range forms {
		fmt.Printf("Surface of %s = %f\n", k, v.Surface())
	}

}
