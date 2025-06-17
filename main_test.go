package main

import "testing"

func TestSumDigits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{-3, 3},
		{0, 0},
		{999999, 54},
		{2147483647, 46},
		{-2147483648, 47},
	}

	for _, test := range tests {
		got := SumDigits(test.n)
		expected := test.want
		if got != expected {
			t.Errorf("SumDigits(%d) expected %d got %d", test.n, expected, got)
		}
	}
}

func BenchmarkSumDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumDigits(i)
	}
}
