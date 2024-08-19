package recursionI

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		expected []int
	}{
		{
			name:     "ZeroIndex",
			index:    0,
			expected: []int{1},
		},
		{
			name:     "FirstIndex",
			index:    1,
			expected: []int{1, 1},
		},
		{
			name:     "SecondIndex",
			index:    2,
			expected: []int{1, 2, 1},
		},
		{
			name:     "ThirdIndex",
			index:    3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "FourthIndex",
			index:    4,
			expected: []int{1, 4, 6, 4, 1},
		},
		{
			name:     "LargeIndex",
			index:    10,
			expected: []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := GetRow(tt.index)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkGetRow(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		GetRow(10)
	}
}

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{
			name:     "PositiveExponent",
			x:        2.0,
			n:        3,
			expected: 8.0,
		},
		{
			name:     "NegativeExponent",
			x:        2.0,
			n:        -3,
			expected: 0.125,
		},
		{
			name:     "ZeroExponent",
			x:        2.0,
			n:        0,
			expected: 1.0,
		},
		{
			name:     "OneExponent",
			x:        2.0,
			n:        1,
			expected: 2.0,
		},
		{
			name:     "NegativeBasePositiveExponent",
			x:        -2.0,
			n:        3,
			expected: -8.0,
		},
		{
			name:     "NegativeBaseNegativeExponent",
			x:        -2.0,
			n:        -3,
			expected: -0.125,
		},
		{
			name:     "FractionalBasePositiveExponent",
			x:        0.5,
			n:        3,
			expected: 0.125,
		},
		{
			name:     "FractionalBaseNegativeExponent",
			x:        0.5,
			n:        -3,
			expected: 8.0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MyPow(tt.x, tt.n)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMyPow(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		MyPow(2.0, 3)
	}
}
