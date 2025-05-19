package geometry

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}
type Rectangle struct {
	Height float64
	Width float64
}
type Triangle struct {
	Base float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}

func (t Triangle) Perimeter() float64 {
	return 0.0
}