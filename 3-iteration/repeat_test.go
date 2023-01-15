package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat characters with default repeat count 5", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat characters with repeat count passed by parameter", func(t *testing.T) {
		repeated := Repeat("o", 10)
		expected := "oooooooooo"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	// b.N
	// when the benchmark code is executed, it runs b.N times and measures how long it takes
	// to run the benchmarks, from cli, run: go test -bench=.
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
