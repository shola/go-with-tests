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
// IMPORTANT: the sides should be in listed in ascending order!
type Triangle struct {
	A float64
	B float64
	C float64
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
	return t.A * t.B / 2
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}