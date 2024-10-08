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

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "RotateArrayByZero",
			nums:     []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "RotateArrayByOne",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			name:     "RotateArrayByLength",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "RotateArrayByMoreThanLength",
			nums:     []int{1, 2, 3, 4, 5},
			k:        7,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "RotateArrayWithNegativeNumbers",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "RotateArrayWithSingleElement",
			nums:     []int{1},
			k:        3,
			expected: []int{1},
		},
		{
			name:     "RotateEmptyArray",
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			Rotate(tt.nums, tt.k)
			assert.Equal(t, tt.expected, tt.nums)
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	k := 3
	for i := 0; i < b.N; i++ {
		Rotate(nums, k)
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "MaxProfitSingleDay",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "MaxProfitIncreasingPrices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "MaxProfitDecreasingPrices",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "MaxProfitMixedPrices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "MaxProfitNoProfit",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MaxProfit(tt.prices)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		_ = MaxProfit(prices)
	}
}

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "MaxProfitIISingleDay",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "MaxProfitIIIncreasingPrices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "MaxProfitIIDecreasingPrices",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "MaxProfitIIMixedPrices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "MaxProfitIINoProfit",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "MaxProfitIIEmptyArray",
			prices:   []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MaxProfitII(tt.prices)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaxProfitII(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4}
	for i := 0; i < b.N; i++ {
		_ = MaxProfitII(prices)
	}
}

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "CanJumpSingleElement",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "CanJumpAllZeros",
			nums:     []int{0, 0, 0, 0},
			expected: false,
		},
		{
			name:     "CanJumpReachEnd",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "CanJumpCannotReachEnd",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
		{
			name:     "CanJumpEmptyArray",
			nums:     []int{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := CanJump(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCanJump(b *testing.B) {
	nums := []int{2, 3, 1, 1, 4}
	for i := 0; i < b.N; i++ {
		_ = CanJump(nums)
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleJump",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "MultipleJumps",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "3 Jumps",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Jump(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkJump(b *testing.B) {
	nums := []int{2, 3, 1, 1, 4}
	for i := 0; i < b.N; i++ {
		_ = Jump(nums)
	}
}

func TestHIndex(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "SinglePaper",
			citations: []int{1},
			expected:  1,
		},
		{
			name:      "NoCitations",
			citations: []int{0, 0, 0, 0},
			expected:  0,
		},
		{
			name:      "AllCitationsSame",
			citations: []int{3, 3, 3, 3},
			expected:  3,
		},
		{
			name:      "MixedCitations",
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			name:      "EmptyArray",
			citations: []int{},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := HIndex(tt.citations)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkHIndex(b *testing.B) {
	citations := []int{3, 0, 6, 1, 5}
	for i := 0; i < b.N; i++ {
		_ = HIndex(citations)
	}
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		expected int
	}{
		{
			name:     "StrStrNeedleAtStart",
			haystack: "hello",
			needle:   "he",
			expected: 0,
		},
		{
			name:     "StrStrNeedleAtEnd",
			haystack: "hello",
			needle:   "lo",
			expected: 3,
		},
		{
			name:     "StrStrNeedleInMiddle",
			haystack: "hello",
			needle:   "ll",
			expected: 2,
		},
		{
			name:     "StrStrNeedleNotFound",
			haystack: "hello",
			needle:   "world",
			expected: -1,
		},
		{
			name:     "StrStrEmptyNeedle",
			haystack: "hello",
			needle:   "",
			expected: 0,
		},
		{
			name:     "StrStrEmptyHaystack",
			haystack: "",
			needle:   "a",
			expected: -1,
		},
		{
			name:     "StrStrBothEmpty",
			haystack: "",
			needle:   "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := StrStr(tt.haystack, tt.needle)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	haystack := "hello"
	needle := "ll"
	for i := 0; i < b.N; i++ {
		_ = StrStr(haystack, needle)
	}
}
