package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat char 0 number of times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat char 'x' number of times", func(t *testing.T) {
		repeated := Repeat("b", 3)
		expected := "bbb"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
