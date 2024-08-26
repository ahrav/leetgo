package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleElement",
			nums:     []int{10},
			expected: 1,
		},
		{
			name:     "AllIncreasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "AllDecreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "MixedElements",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Duplicates",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := LengthOfLIS(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "EmptyArray",
			nums: []int{},
		},
		{
			name: "SingleElement",
			nums: []int{10},
		},
		{
			name: "AllIncreasing",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "AllDecreasing",
			nums: []int{5, 4, 3, 2, 1},
		},
		{
			name: "MixedElements",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
		},
		{
			name: "Duplicates",
			nums: []int{2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = LengthOfLIS(tt.nums)
			}
		})
	}
}

func TestLengthOfLISBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleElement",
			nums:     []int{10},
			expected: 1,
		},
		{
			name:     "AllIncreasing",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "AllDecreasing",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 1,
		},
		{
			name:     "MixedElements",
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			name:     "Duplicates",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := LengthOfLISBinarySearch(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLengthOfLISBinarySearch(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "EmptyArray",
			nums: []int{},
		},
		{
			name: "SingleElement",
			nums: []int{10},
		},
		{
			name: "AllIncreasing",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "AllDecreasing",
			nums: []int{5, 4, 3, 2, 1},
		},
		{
			name: "MixedElements",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
		},
		{
			name: "Duplicates",
			nums: []int{2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = LengthOfLISBinarySearch(tt.nums)
			}
		})
	}
}
