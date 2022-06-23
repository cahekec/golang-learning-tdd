package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectRepeatedString := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Expected '%q', bug got '%q'", want, got)
		}
	}

	t.Run("Test one", func(t *testing.T) {
		got := Repeat("a", 6)
		want := "aaaaaa"

		assertCorrectRepeatedString(t, got, want)
	})

	t.Run("Test two", func(t *testing.T) {
		got := Repeat("b", 5)
		want := "bbbbb"

		assertCorrectRepeatedString(t, got, want)
	})
}

func BenchmarRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}
