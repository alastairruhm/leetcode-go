package reverseInteger

import "testing"

var (
	input    = 1000000003
	expected = 0
)

func benchmarkReverseInteger(reverseInteger func(int) int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseInteger(input)
	}
}

func BenchmarkReverse(b *testing.B)              { benchmarkReverseInteger(reverse, b) }
func BenchmarkReverseWithGoroutine(b *testing.B) { benchmarkReverseInteger(reverseWithGoroutine, b) }
func BenchmarkReverseIn3msSolution(b *testing.B) { benchmarkReverseInteger(reverseIn3msSolution, b) }
