package recursion

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

func TestKthGrammar(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		expected int
	}{
		{
			name:     "FirstSymbol",
			n:        1,
			k:        1,
			expected: 0,
		},
		{
			name:     "SecondRowFirstSymbol",
			n:        2,
			k:        1,
			expected: 0,
		},
		{
			name:     "SecondRowSecondSymbol",
			n:        2,
			k:        2,
			expected: 1,
		},
		{
			name:     "ThirdRowFirstSymbol",
			n:        3,
			k:        1,
			expected: 0,
		},
		{
			name:     "ThirdRowSecondSymbol",
			n:        3,
			k:        2,
			expected: 1,
		},
		{
			name:     "ThirdRowThirdSymbol",
			n:        3,
			k:        3,
			expected: 1,
		},
		{
			name:     "ThirdRowFourthSymbol",
			n:        3,
			k:        4,
			expected: 0,
		},
		{
			name:     "FourthRowFifthSymbol",
			n:        4,
			k:        5,
			expected: 1,
		},
		{
			name:     "LargeNAndK",
			n:        10,
			k:        512,
			expected: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := KthGrammar(tt.n, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "EmptyArray",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "SingleElement",
			nums:     []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "TwoElements",
			nums:     []int{1, 2},
			expected: [][]int{{1, 2}, {2, 1}},
		},
		{
			name:     "ThreeElements",
			nums:     []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name:     "DuplicateElements",
			nums:     []int{1, 1},
			expected: [][]int{{1, 1}, {1, 1}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := Permute(tt.nums)
			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}

func BenchmarkPermute(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{"SmallSet", []int{1, 2, 3}},
		{"MediumSet", []int{1, 2, 3, 4, 5}},
		{"LargeSet", []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Permute(tt.nums)
			}
		})
	}
}
