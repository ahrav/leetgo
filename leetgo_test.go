package leetgo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := TwoSum(tt.nums, tt.target)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = TwoSum([]int{2, 7, 11, 15}, 9)
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
			name:     "Example 1",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "Example 2",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "Example 3",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
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
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Example 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Example 3",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MinEatingSpeed(tt.piles, tt.h)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinEatingSpeed(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinEatingSpeed([]int{3, 6, 7, 11}, 8)
	}
}

func TestCountGroups(t *testing.T) {
	tests := []struct {
		name     string
		related  []string
		expected int
	}{
		{
			name:     "Example 1",
			related:  []string{"110", "110", "001"},
			expected: 2,
		},
		{
			name:     "Example 2",
			related:  []string{"110", "110", "011"},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := CountGroups(tt.related)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCountGroups(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CountGroups([]string{"110", "110", "001"})
	}
}

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		name     string
		roman    string
		expected int
	}{
		{
			name:     "Example 1",
			roman:    "III",
			expected: 3,
		},
		{
			name:     "Example 2",
			roman:    "IV",
			expected: 4,
		},
		{
			name:     "Example 3",
			roman:    "IX",
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := RomanToInteger(tt.roman)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkRomanToInteger(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = RomanToInteger("III")
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			name:     "Example 2",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			name:     "Example 3",
			nums1:    []int{0, 0},
			nums2:    []int{0, 0},
			expected: 0.0,
		},
		{
			name:     "Example 4",
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			name:     "Example 5",
			nums1:    []int{2},
			nums2:    []int{},
			expected: 2.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := FindMedianSortedArrays(tt.nums1, tt.nums2)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FindMedianSortedArrays([]int{1, 3}, []int{2})
	}
}

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{90},
			expected: 0,
		},
		{
			name:     "Example 2",
			nums:     []int{9, 4, 1, 7},
			expected: 0,
		},
		{
			name:     "Example 3",
			nums:     []int{6, 6, 0, 1, 1, 4, 6},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MinimumDifference(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinimumDifference(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinimumDifference([]int{6, 6, 0, 1, 1, 4, 6})
	}
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1",
			l1:       &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			name:     "Example 2",
			l1:       &ListNode{Val: 0},
			l2:       &ListNode{Val: 0},
			expected: &ListNode{Val: 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := AddTwoNumbers(tt.l1, tt.l2)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = AddTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "Example 1",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "Example 2",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := GroupAnagrams(tt.strs)
			sortGroups := cmpopts.SortSlices(func(a, b []string) bool {
				return len(a) < len(b) || (len(a) == len(b) && a[0] < b[0])
			})
			sortStrings := cmpopts.SortSlices(func(a, b string) bool {
				return a < b
			})

			if diff := cmp.Diff(tt.expected, actual, sortGroups, sortStrings); diff != "" {
				t.Errorf("GroupAnagrams() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Example 2",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Example 3",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Example 4",
			prices:   []int{1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MaxProfit(tt.prices)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MaxProfit([]int{7, 1, 5, 3, 6, 4})
	}
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "Example 3",
			grid: [][]byte{
				{'1', '0', '1', '1', '1'},
				{'1', '0', '1', '0', '1'},
				{'1', '1', '1', '0', '1'},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := NumIslands(tt.grid)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkNumIslands(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NumIslands([][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		})
	}
}
