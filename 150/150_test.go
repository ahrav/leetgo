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

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "EmptySecondArray",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3, 0, 0, 0},
		},
		{
			name:     "EmptyFirstArray",
			nums1:    []int{0, 0, 0},
			m:        0,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{2, 5, 6},
		},
		{
			name:     "NonOverlappingArrays",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "OverlappingArrays",
			nums1:    []int{1, 3, 5, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 4, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "WithDuplicates",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 2, 2},
			n:        3,
			expected: []int{1, 2, 2, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.expected, tt.nums1)
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	for i := 0; i < b.N; i++ {
		Merge(nums1, 3, nums2, 3)
	}
}

func TestRemoveDuplicatesII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
		result   []int
	}{
		{
			name:     "SingleElementArray",
			nums:     []int{1},
			expected: 1,
			result:   []int{1},
		},
		{
			name:     "TwoElementArrayNoDuplicates",
			nums:     []int{1, 2},
			expected: 2,
			result:   []int{1, 2},
		},
		{
			name:     "TwoElementArrayWithDuplicates",
			nums:     []int{1, 1},
			expected: 2,
			result:   []int{1, 1},
		},
		{
			name:     "MultipleDuplicates",
			nums:     []int{1, 1, 1, 2, 2, 3},
			expected: 5,
			result:   []int{1, 1, 2, 2, 3},
		},
		{
			name:     "NoDuplicates",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
			result:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "AllElementsSame",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 2,
			result:   []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := RemoveDuplicatesII(tt.nums)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.result, tt.nums[:result])
		})
	}
}

func BenchmarkRemoveDuplicatesII(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		_ = RemoveDuplicatesII(nums)
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		expected int
		result   []int
	}{
		{
			name:     "SingleElementArray",
			nums:     []int{1},
			val:      1,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "MultipleElementsNoMatch",
			nums:     []int{1, 2, 3, 4, 5},
			val:      6,
			expected: 5,
			result:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "MultipleElementsWithMatch",
			nums:     []int{1, 2, 3, 4, 5},
			val:      3,
			expected: 4,
			result:   []int{1, 2, 4, 5},
		},
		{
			name:     "AllElementsMatch",
			nums:     []int{2, 2, 2, 2, 2},
			val:      2,
			expected: 0,
			result:   []int{},
		},
		{
			name:     "MixedElements",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			expected: 2,
			result:   []int{2, 2},
		},
		{
			name:     "EmptyArray",
			nums:     []int{},
			val:      1,
			expected: 0,
			result:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := RemoveElement(tt.nums, tt.val)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.result, tt.nums[:result])
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{3, 2, 2, 3}
	val := 3
	for i := 0; i < b.N; i++ {
		_ = RemoveElement(nums, val)
	}
}

func TestRemoveDuplicatesI(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
		result   []int
	}{
		{
			name:     "SingleElementArrayNoDuplicates",
			nums:     []int{1},
			expected: 1,
			result:   []int{1},
		},
		{
			name:     "MultipleElementsNoDuplicates",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
			result:   []int{1, 2, 3, 4, 5},
		},
		{
			name:     "MultipleElementsWithDuplicates",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: 3,
			result:   []int{1, 2, 3},
		},
		{
			name:     "AllElementsSame",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 1,
			result:   []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := RemoveDuplicatesI(tt.nums)
			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.result, tt.nums[:result])
		})
	}
}

func BenchmarkRemoveDuplicatesI(b *testing.B) {
	nums := []int{1, 1, 2, 2, 3, 3}
	for i := 0; i < b.N; i++ {
		_ = RemoveDuplicatesI(nums)
	}
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleElement",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "AllSameElements",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "MajorityElementInMiddle",
			nums:     []int{1, 2, 3, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "MajorityElementAtEnd",
			nums:     []int{1, 1, 1, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "MixedElements",
			nums:     []int{3, 3, 4, 2, 4, 4, 2, 4, 4},
			expected: 4,
		},
		{
			name:     "NoMajorityElement",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MajorityElement(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	for i := 0; i < b.N; i++ {
		_ = MajorityElement(nums)
	}
}
