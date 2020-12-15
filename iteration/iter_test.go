package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat letter 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

	t.Run("repeat letter 'a' 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})

	t.Run("repeat letter 'a' 1 time", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
