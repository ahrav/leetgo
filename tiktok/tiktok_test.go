package tiktok

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "Finds min subarray length for exact match",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			name:     "Finds min subarray length for no match",
			target:   100,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Finds min subarray length for single element",
			target:   4,
			nums:     []int{4},
			expected: 1,
		},
		{
			name:     "Finds min subarray length for empty array",
			target:   4,
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Finds min subarray length for multiple possible subarrays",
			target:   11,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinSubArrayLen(tt.target, tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	for i := 0; i < b.N; i++ {
		MinSubArrayLen(target, nums)
	}
}

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Finds max product for positive numbers",
			nums:     []int{2, 3, 4, 5},
			expected: 120,
		},
		{
			name:     "Finds max product for negative numbers",
			nums:     []int{-2, -3, -4, -5},
			expected: 120,
		},
		{
			name:     "Finds max product for mixed numbers",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "Finds max product for single element",
			nums:     []int{4},
			expected: 4,
		},
		{
			name:     "Finds max product for array with zero",
			nums:     []int{0, 2, 3, -2, 4},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MaxProduct(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaxProduct(b *testing.B) {
	nums := []int{2, 3, -2, 4}
	for i := 0; i < b.N; i++ {
		MaxProduct(nums)
	}
}
