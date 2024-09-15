package onefifty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"SingleElement", []int{5}, 5},
		{"AllNegative", []int{-2, -3, -1, -4}, -1},
		{"MixedElements", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"AllPositive", []int{1, 2, 3, 4, 5}, 15},
		{"EmptyArray", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MaxSubArray(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for i := 0; i < b.N; i++ {
		MaxSubArray(nums)
	}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{
			name:     "NoSolution",
			numbers:  []int{1, 2, 3, 4, 5},
			target:   10,
			expected: nil,
		},
		{
			name:     "SinglePair",
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "MultiplePairs",
			numbers:  []int{1, 2, 3, 4, 5, 5},
			target:   8,
			expected: []int{3, 6},
		},
		{
			name:     "NegativeNumbers",
			numbers:  []int{-3, -1, 0, 1, 2},
			target:   -1,
			expected: []int{1, 5},
		},
		{
			name:     "EmptyArray",
			numbers:  []int{},
			target:   5,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := TwoSum(tt.numbers, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 4, 5}
	target := 8
	for i := 0; i < b.N; i++ {
		TwoSum(numbers, target)
	}
}
