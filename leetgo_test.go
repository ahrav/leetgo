package leetgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1",
			heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			heights:  []int{1, 1},
			expected: 1,
		},
		{
			name:     "Example 3",
			heights:  []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			name:     "Example 4",
			heights:  []int{1, 2, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MaxArea(tt.heights)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaxArea(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	}
}

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name     string
		number   int
		expected string
	}{
		{
			name:     "Example 1",
			number:   3,
			expected: "III",
		},
		{
			name:     "Example 2",
			number:   4,
			expected: "IV",
		},
		{
			name:     "Example 3",
			number:   9,
			expected: "IX",
		},
		{
			name:     "Example 4",
			number:   58,
			expected: "LVIII",
		},
		{
			name:     "Example 5",
			number:   1994,
			expected: "MCMXCIV",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := IntToRoman(tt.number)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = IntToRoman(1994)
	}
}
