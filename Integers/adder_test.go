package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	asserdCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("expected '%d' but got '%d'", want, got)
		}
	}

	t.Run("2 add 2 should be equal 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		asserdCorrectMessage(t, sum, expected)
	})
}
