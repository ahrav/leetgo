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
