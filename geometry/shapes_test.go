package geometry

import (
	"math"
	"testing"
)

// Table driven testing of interfaces make test suites easier
// to extend and maintain
func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea  float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle",shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assertShapeValue(t, tt.shape, got, tt.hasArea)
		})
	}
}

func assertShapeValue(t testing.TB, shape Shape, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}

func TestRectangle(t *testing.T) {
	r := Rectangle{Height: 10.0, Width: 10.0}
	t.Run("Perimeter", func(t *testing.T) {
		got := r.Perimeter()
		want := 40.0
		assertFloat(t, got, want)
	})
	t.Run("Area", func(t *testing.T) {
		got := r.Area()
		want := 100.0
		assertFloat(t, got, want)
	})
}

func TestCircle(t *testing.T) {
	c := Circle{Radius: 10.0}
	t.Run("Perimeter", func(t *testing.T) {
		got := c.Perimeter()
		want := 2 * 10.0 * math.Pi
		assertFloat(t, got, want)
	})
	t.Run("Area", func(t *testing.T) {
		got := c.Area()
		want := 10.0 * 10.0 * math.Pi
		assertFloat(t, got, want)
	})
}

func assertFloat(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
