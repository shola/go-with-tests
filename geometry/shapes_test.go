package geometry

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Height: 12, Width: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		assertFloat(t, got, tt.want)
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
