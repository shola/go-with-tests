package geometry

import (
	"math"
	"testing"
)

func TestRectangle(t *testing.T) {
	r := Rectangle{Height: 10.0, Width: 10.0}
	t.Run("Perimeter", func (t *testing.T) {
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
	t.Run("Perimeter", func (t *testing.T) {
		got := c.Perimeter()
		want := 2 * 10.0 * math.Pi
		assertFloat(t, got, want)
	})
}

func assertFloat(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}