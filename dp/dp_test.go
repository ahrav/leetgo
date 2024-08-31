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

func TestLargestDivisibleSubset(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "LargestDivisibleSubsetSingleElement",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "LargestDivisibleSubsetAllDivisible",
			nums:     []int{1, 2, 4, 8},
			expected: []int{1, 2, 4, 8},
		},
		{
			name:     "LargestDivisibleSubsetMixedElements",
			nums:     []int{1, 2, 3, 4, 6, 8, 12},
			expected: []int{1, 2, 4, 8},
		},
		{
			name:     "LargestDivisibleSubsetNoDivisiblePairs",
			nums:     []int{2, 3, 5, 7, 11},
			expected: []int{2},
		},
		{
			name:     "LargestDivisibleSubsetEmptyArray",
			nums:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LargestDivisibleSubset(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLargestDivisibleSubset(b *testing.B) {
	nums := []int{1, 2, 3, 4, 6, 8, 12}
	for i := 0; i < b.N; i++ {
		_ = LargestDivisibleSubset(nums)
	}
}

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"ValidDecoding", "12", 2},
		{"SingleDigitDecoding", "2", 1},
		{"LeadingZeroDecoding", "012", 0},
		{"MultipleZerosDecoding", "100", 0},
		{"LargeNumberDecoding", "226", 3},
		{"EmptyStringDecoding", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := NumDecodings(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkNumDecodings(b *testing.B) {
	input := "226"
	for i := 0; i < b.N; i++ {
		NumDecodings(input)
	}
}

func TestJobScheduling(t *testing.T) {
	tests := []struct {
		name      string
		startTime []int
		endTime   []int
		profit    []int
		expected  int
	}{
		{"SingleJob", []int{1}, []int{2}, []int{50}, 50},
		{"NonOverlappingJobs", []int{1, 3, 6}, []int{2, 5, 9}, []int{50, 10, 40}, 100},
		{"OverlappingJobs", []int{1, 2, 3, 4}, []int{3, 5, 10, 6}, []int{20, 20, 100, 70}, 120},
		{"MixedJobs", []int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}, 120},
		{"EmptyJobs", []int{}, []int{}, []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := JobScheduling(tt.startTime, tt.endTime, tt.profit)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkJobScheduling(b *testing.B) {
	startTime := []int{1, 2, 3, 3}
	endTime := []int{3, 4, 5, 6}
	profit := []int{50, 10, 40, 70}
	for i := 0; i < b.N; i++ {
		JobScheduling(startTime, endTime, profit)
	}
}
