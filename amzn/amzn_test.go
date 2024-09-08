package amzn

import (
	"fmt"
	"math"
	"reflect"
	"slices"
	"sort"
	"strings"
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

func TestIsPalindromeSList(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected bool
	}{
		{
			name:     "Example 1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			expected: true,
		},
		{
			name:     "Example 2",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: false,
		},
		{
			name:     "Example 3",
			head:     &ListNode{Val: 1},
			expected: true,
		},
		{
			name:     "Example 4",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := IsPalindromeSList(tt.head)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		p, q     *TreeNode
		expected *TreeNode
	}{
		{
			name: "LCA of 5 and 1",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 1},
			expected: &TreeNode{Val: 2},
		},
		{
			name: "LCA of 2 and 4",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p:        &TreeNode{Val: 2},
			q:        &TreeNode{Val: 4},
			expected: &TreeNode{Val: 2},
		},
		{
			name: "LCA of 7 and 9",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{
					Val:   8,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 9},
				},
			},
			p:        &TreeNode{Val: 7},
			q:        &TreeNode{Val: 9},
			expected: &TreeNode{Val: 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LowestCommonAncestor(tt.root, tt.p, tt.q)
			assert.Equal(t, tt.expected.Val, actual.Val)
		})
	}
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{120, 60, 40, 30, 24},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ProductExceptSelf(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ProductExceptSelf([]int{1, 2, 3, 4})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s, t     string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			name:     "Example 3",
			s:        "a",
			t:        "ab",
			expected: false,
		},
		{
			name:     "Example 4",
			s:        "a",
			t:        "a",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsAnagram(tt.s, tt.t)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = IsAnagram("anagram", "nagaram")
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "Example 1",
			n:        15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name:     "Example 2",
			n:        1,
			expected: []string{"1"},
		},
		{
			name:     "Example 3",
			n:        3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "Example 4",
			n:        5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FizzBuzz(tt.n)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FizzBuzz(15)
	}
}

func TestLongestPalindromeSubseq(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "bbbab",
			expected: 4,
		},
		{
			name:     "Example 2",
			s:        "cbbd",
			expected: 2,
		},
		{
			name:     "Example 3",
			s:        "a",
			expected: 1,
		},
		{
			name:     "Example 4",
			s:        "ac",
			expected: 1,
		},
		{
			name:     "Example 5",
			s:        "abcda",
			expected: 3,
		},
		{
			name:     "Example 6",
			s:        "abcdba",
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LongestPalindromeSubseq(tt.s)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLongestPalindromeSubseq(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LongestPalindromeSubseq("bbbab")
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Example 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			name:     "Example 2",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			name:     "Example 4",
			coins:    []int{1},
			amount:   1,
			expected: 1,
		},
		{
			name:     "Example 5",
			coins:    []int{1},
			amount:   2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CoinChange(tt.coins, tt.amount)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCoinChange(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CoinChange([]int{1, 2, 5}, 11)
	}
}

func TestBoundaryOfBinaryTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 9},
					},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 10},
						Right: &TreeNode{Val: 11},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 13},
					},
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 14},
						Right: &TreeNode{Val: 15},
					},
				},
			},
			expected: []int{1, 2, 4, 8, 9, 10, 11, 12, 13, 14, 15, 7, 3},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
				},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
				},
			},
			expected: []int{1, 4, 5, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := BoundaryOfBinaryTree(tt.root)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkBoundaryOfBinaryTree(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = BoundaryOfBinaryTree(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 9},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 10},
					Right: &TreeNode{Val: 11},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 12},
					Right: &TreeNode{Val: 13},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 14},
					Right: &TreeNode{Val: 15},
				},
			},
		})
	}
}

func TestNumberOfDistinctIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{1, 1, 0, 1, 1},
			},
			expected: 3,
		},
		{
			name: "Example 3",
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{1, 1, 0, 1, 0},
			},
			expected: 3,
		},
		{
			name: "Example 4",
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{1, 1, 0, 0, 1},
			},
			expected: 3,
		},
		{
			name: "Example 5",
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1},
				{1, 1, 0, 0, 0},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NumberOfDistinctIslands(tt.grid)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkNumberOfDistinctIslands(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NumberOfDistinctIslands([][]int{
			{1, 1, 0, 1, 1},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1},
			{1, 1, 0, 1, 1},
		})
	}
}

func TestBinaryTreeHeight(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 8},
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: 4,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 8},
					},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: 4,
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := BinaryTreeHeight(tt.root)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkBinaryTreeHeight(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = BinaryTreeHeight(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 8},
				},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 9},
			},
		})
	}
}

func TestBinaryTreeDiameter(t *testing.T) {
	tests := []struct {
		name     string
		tree     *TreeNode
		expected int
	}{
		{
			name:     "Single node tree",
			tree:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name:     "Two node tree",
			tree:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			expected: 2,
		},
		{
			name:     "Three node balanced tree",
			tree:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			expected: 2,
		},
		{
			name:     "Four node unbalanced tree",
			tree:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
			expected: 3,
		},
		{
			name: "Five node tree with longest path not through root",
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: 3,
		},
		{
			name:     "Nil tree",
			tree:     nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinaryTreeDiameter(tt.tree)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkBinaryTreeDiameter(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = BinaryTreeDiameter(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{Val: 3},
		})
	}
}

func TestDistanceK(t *testing.T) {
	tests := []struct {
		name   string
		root   *TreeNode
		target *TreeNode
		k      int
		want   []int
	}{
		{
			name:   "Empty tree",
			root:   nil,
			target: nil,
			k:      2,
			want:   nil,
		},
		{
			name:   "Single node tree, k=0",
			root:   &TreeNode{Val: 1},
			target: &TreeNode{Val: 1},
			k:      0,
			want:   []int{1},
		},
		{
			name: "Target is leaf, k exceeds tree depth",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			target: &TreeNode{Val: 2},
			k:      3,
			want:   []int{},
		},
		{
			name: "Complex tree, multiple nodes at distance k",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 12},
					Right: &TreeNode{Val: 8},
				},
			},
			target: &TreeNode{Val: 5},
			k:      2,
			want:   []int{1, 4, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.root != nil && tt.root.Left != nil {
				tt.target = tt.root.Left
			}

			got := DistanceK(tt.root, tt.target, tt.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDistanceK(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = DistanceK(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
			},
			Right: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 8},
			},
		}, &TreeNode{Val: 5}, 2)
	}
}

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 0, 0, 0, 1, 1, 1},
			k:        3,
			expected: 8,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 0, 0, 0, 0, 1, 1, 1},
			k:        3,
			expected: 6,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 0, 0, 0, 0, 1, 1, 1},
			k:        0,
			expected: 3,
		},
		{
			name:     "Example 4",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			k:        0,
			expected: 8,
		},
		{
			name:     "Example 5",
			nums:     []int{0, 0, 0, 0, 0, 0, 0, 0},
			k:        3,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LongestOnes(tt.nums, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLongestOnes(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LongestOnes([]int{1, 1, 0, 0, 0, 1, 1, 1}, 3)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "abcabcbb",
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "bbbbb",
			expected: 1,
		},
		{
			name:     "Example 3",
			s:        "pwwkew",
			expected: 3,
		},
		{
			name:     "Example 4",
			s:        "",
			expected: 0,
		},
		{
			name:     "Example 5",
			s:        " ",
			expected: 1,
		},
		{
			name:     "Example 6",
			s:        "au",
			expected: 2,
		},
		{
			name:     "Example 7",
			s:        "dvdf",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := LengthOfLongestSubstring(tt.s)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LengthOfLongestSubstring("abcabcbb")
	}
}

func TestMostVisitedPattern(t *testing.T) {
	tests := []struct {
		name      string
		username  []string
		timestamp []int
		website   []string
		expected  []string
	}{
		{
			name:      "Example 1",
			username:  []string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"},
			timestamp: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			website:   []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"},
			expected:  []string{"home", "about", "career"},
		},
		{
			name:      "Example 2",
			username:  []string{"u1", "u1", "u1", "u2", "u2", "u2"},
			timestamp: []int{1, 2, 3, 4, 5, 6},
			website:   []string{"a", "b", "a", "a", "b", "c"},
			expected:  []string{"a", "b", "a"},
		},
		{
			name:      "Example 3",
			username:  []string{"u1", "u2", "u1", "u2", "u1", "u2"},
			timestamp: []int{1, 2, 3, 4, 5, 6},
			website:   []string{"a", "a", "a", "a", "a", "a"},
			expected:  []string{"a", "a", "a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MostVisitedPattern(tt.username, tt.website, tt.timestamp)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMostVisitedPattern(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MostVisitedPattern([]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"}, []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func TestSortJumbled(t *testing.T) {
	tests := []struct {
		name    string
		mapping []int
		nums    []int
		want    []int
	}{
		{
			name:    "Empty input",
			mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			nums:    []int{},
			want:    []int{},
		},
		{
			name:    "Single digit numbers",
			mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			nums:    []int{0, 1, 2, 3, 4, 5},
			want:    []int{5, 4, 3, 2, 1, 0},
		},
		{
			name:    "All same mapped value",
			mapping: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			nums:    []int{10, 20, 30, 40, 50},
			want:    []int{10, 20, 30, 40, 50},
		},
		{
			name:    "Large numbers",
			mapping: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			nums:    []int{1000000, 10000, 100, 1},
			want:    []int{1, 100, 10000, 1000000},
		},
		{
			name:    "Duplicate numbers",
			mapping: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			nums:    []int{555, 901, 555, 123, 901},
			want:    []int{123, 555, 555, 901, 901},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortJumbled(tt.mapping, tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortJumbled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSortJumbled(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = SortJumbled([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{1000000, 10000, 100, 1})
	}
}

func TestPlatesBetweenCandles(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		queries [][]int
		want    []int
	}{
		{
			name:    "Empty string",
			s:       "",
			queries: [][]int{{0, 0}},
			want:    nil,
		},
		{
			name:    "No plates",
			s:       "|||",
			queries: [][]int{{0, 2}},
			want:    []int{0},
		},
		{
			name:    "All plates",
			s:       "****",
			queries: [][]int{{0, 3}},
			want:    []int{0},
		},
		{
			name:    "Single plate between candles",
			s:       "|*|",
			queries: [][]int{{0, 2}},
			want:    []int{1},
		},
		{
			name:    "Multiple queries with varying results",
			s:       "||**||**|*",
			queries: [][]int{{2, 5}, {5, 9}},
			want:    []int{0, 2},
		},
		{
			name:    "Queries with no plates between candles",
			s:       "|**|**|",
			queries: [][]int{{0, 2}, {4, 6}},
			want:    []int{0, 0},
		},
		{
			name:    "Overlapping queries",
			s:       "|**|*|**|",
			queries: [][]int{{0, 5}, {2, 7}},
			want:    []int{3, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlatesBetweenCandles(tt.s, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlatesBetweenCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPlatesBetweenCandles(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = PlatesBetweenCandles("||**||**|*", [][]int{{2, 5}, {5, 9}})
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Two different characters",
			input:    "ab",
			expected: "a",
		},
		{
			name:     "Two same characters",
			input:    "aa",
			expected: "aa",
		},
		{
			name:     "Odd length palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "Even length palindrome",
			input:    "abccba",
			expected: "abccba",
		},
		{
			name:     "Multiple palindromes",
			input:    "abcddcbafg",
			expected: "abcddcba",
		},
		{
			name:     "Palindrome at the beginning",
			input:    "abbacd",
			expected: "abba",
		},
		{
			name:     "Palindrome at the end",
			input:    "cdabba",
			expected: "abba",
		},
		{
			name:     "No palindrome longer than 1",
			input:    "abcdef",
			expected: "a",
		},
		{
			name:     "All same characters",
			input:    "aaaaa",
			expected: "aaaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := LongestPalindromeDP(tt.input)
			result := LongestPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("LongestPalindrome(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LongestPalindrome("abcddcbafg")
	}
}

func TestLongestPalindromeDP(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Two different characters",
			input:    "ab",
			expected: "a",
		},
		{
			name:     "Two same characters",
			input:    "aa",
			expected: "aa",
		},
		{
			name:     "Odd length palindrome",
			input:    "racecar",
			expected: "racecar",
		},
		{
			name:     "Even length palindrome",
			input:    "abccba",
			expected: "abccba",
		},
		{
			name:     "Multiple palindromes",
			input:    "abcddcbafg",
			expected: "abcddcba",
		},
		{
			name:     "Palindrome at the beginning",
			input:    "abbacd",
			expected: "abba",
		},
		{
			name:     "Palindrome at the end",
			input:    "cdabba",
			expected: "abba",
		},
		{
			name:     "No palindrome longer than 1",
			input:    "abcdef",
			expected: "a",
		},
		{
			name:     "All same characters",
			input:    "aaaaa",
			expected: "aaaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestPalindromeDP(tt.input)
			if result != tt.expected {
				t.Errorf("LongestPalindrome(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestPalindromeDP(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LongestPalindromeDP("abcddcbafg")
	}
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Empty input",
			digits:   "",
			expected: nil,
		},
		{
			name:     "Single digit",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Two digits",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Three digits",
			digits:   "456",
			expected: []string{"gjm", "gjn", "gjo", "gkm", "gkn", "gko", "glm", "gln", "glo", "hjm", "hjn", "hjo", "hkm", "hkn", "hko", "hlm", "hln", "hlo", "ijm", "ijn", "ijo", "ikm", "ikn", "iko", "ilm", "iln", "ilo"},
		},
		{
			name:     "Digits with no mapping",
			digits:   "01",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LetterCombinations(tt.digits)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("LetterCombinations(%s) = %v, want %v", tt.digits, result, tt.expected)
			}
		})
	}
}

func BenchmarkLetterCombinations(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LetterCombinations("456")
	}
}

func TestSearchRotatedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Target in first half of rotated array",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: 1,
		},
		{
			name:     "Target in second half of rotated array",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: 5,
		},
		{
			name:     "Target is first element",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   4,
			expected: 0,
		},
		{
			name:     "Target is last element",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   2,
			expected: 6,
		},
		{
			name:     "Target not in array",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "Array with one element, target present",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Array with one element, target not present",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "Array with two elements, target in first position",
			nums:     []int{2, 1},
			target:   2,
			expected: 0,
		},
		{
			name:     "Array with two elements, target in second position",
			nums:     []int{2, 1},
			target:   1,
			expected: 1,
		},
		{
			name:     "Rotated array with duplicate elements, target present",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: 3,
		},
		{
			name:     "Rotated array with duplicate elements, target not present",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := SearchRotatedArray(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSearchRotatedArray(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = SearchRotatedArray([]int{4, 5, 6, 7, 0, 1, 2}, 5)
	}
}

func TestNumTeams(t *testing.T) {
	tests := []struct {
		name     string
		rating   []int
		expected int
	}{
		{
			name:     "Empty array",
			rating:   []int{},
			expected: 0,
		},
		{
			name:     "Array with less than 3 elements",
			rating:   []int{1, 2},
			expected: 0,
		},
		{
			name:     "Array with 3 elements in ascending order",
			rating:   []int{1, 2, 3},
			expected: 1,
		},
		{
			name:     "Array with 3 elements in descending order",
			rating:   []int{3, 2, 1},
			expected: 1,
		},
		{
			name:     "Array with 4 elements with multiple teams",
			rating:   []int{2, 5, 3, 4},
			expected: 1,
		},
		{
			name:     "Array with duplicate ratings",
			rating:   []int{1, 2, 3, 4, 4, 5},
			expected: 16,
		},
		{
			name:     "Array with negative ratings",
			rating:   []int{-3, -2, -1, 0, 1, 2},
			expected: 20,
		},
		{
			name:     "Large array with mixed ratings",
			rating:   []int{1, 10, 5, 8, 3, 7, 2, 9, 4, 6},
			expected: 38,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := NumTeams(tt.rating)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkNumTeams(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NumTeams([]int{1, 10, 5, 8, 3, 7, 2, 9, 4, 6})
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "All positive numbers",
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "Large numbers",
			input:    []int{1000000, -1000000, 0},
			expected: [][]int{{-1000000, 0, 1000000}},
		},
		{
			name:     "Duplicate solutions",
			input:    []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ThreeSum(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkThreeSum(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	}
}

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Empty input",
			intervals: nil,
			expected:  nil,
		},
		{
			name:      "Single interval",
			intervals: [][]int{{1, 3}},
			expected:  [][]int{{1, 3}},
		},
		{
			name:      "Two overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 4}},
			expected:  [][]int{{1, 4}},
		},
		{
			name:      "Two non-overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}},
			expected:  [][]int{{1, 2}, {3, 4}},
		},
		{
			name:      "Multiple overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Multiple non-overlapping intervals",
			intervals: [][]int{{1, 4}, {5, 6}},
			expected:  [][]int{{1, 4}, {5, 6}},
		},
		{
			name:      "Multiple overlapping and non-overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {1, 2}, {3, 4}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MergeIntervals(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeIntervals(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}

func BenchmarkMergeIntervals(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}

func TestMergeIntervals2(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Empty input",
			intervals: nil,
			expected:  nil,
		},
		{
			name:      "Single interval",
			intervals: [][]int{{1, 3}},
			expected:  [][]int{{1, 3}},
		},
		{
			name:      "Two overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 4}},
			expected:  [][]int{{1, 4}},
		},
		{
			name:      "Two non-overlapping intervals",
			intervals: [][]int{{1, 2}, {3, 4}},
			expected:  [][]int{{1, 2}, {3, 4}},
		},
		{
			name:      "Multiple overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Multiple non-overlapping intervals",
			intervals: [][]int{{1, 4}, {5, 6}},
			expected:  [][]int{{1, 4}, {5, 6}},
		},
		{
			name:      "Multiple overlapping and non-overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {1, 2}, {3, 4}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MergeIntervals2(tt.intervals)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeIntervals(%v) = %v, want %v", tt.intervals, result, tt.expected)
			}
		})
	}
}

func BenchmarkMergeIntervals2(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MergeIntervals2([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	}
}

func TestLowestCommonAncestorDFS(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		pVal     int
		qVal     int
		expected int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{Val: 2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			pVal:     5,
			qVal:     1,
			expected: 3,
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 6},
					Right: &TreeNode{Val: 2,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 8},
				},
			},
			pVal:     5,
			qVal:     4,
			expected: 5,
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			pVal:     1,
			qVal:     2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := findNode(tt.root, tt.pVal)
			q := findNode(tt.root, tt.qVal)
			result := LowestCommonAncestorDFS(tt.root, p, q)
			if result == nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tt.expected, result.Val)
			}
		})
	}
}

// Helper function to find a node with a specific value in the tree.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := findNode(root.Left, val)
	if left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func BenchmarkLowestCommonAncestorDFS(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p := &TreeNode{Val: 5}
		q := &TreeNode{Val: 1}
		_ = LowestCommonAncestorDFS(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 6},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 4},
				},
			},
			Right: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 8},
			},
		}, p, q)
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Multiple elements",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Large input",
			nums:     []int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4},
			k:        3,
			expected: []int{3, 4, 1},
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			k:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TopKFrequent(tt.nums, tt.k)
			slices.Sort(result)
			slices.Sort(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = TopKFrequent([]int{1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}, 3)
	}
}

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			expected: []int{-1},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: []int{2, -1},
		},
		{
			name:     "Three elements",
			nums:     []int{1, 2, 1},
			expected: []int{2, -1, 2},
		},
		{
			name:     "Multiple elements",
			nums:     []int{1, 2, 1, 3, 4, 2},
			expected: []int{2, 3, 3, 4, -1, 3},
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			expected: []int{-1, -1, -1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := NextGreaterElements(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("NextGreaterElements(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkNextGreaterElements(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NextGreaterElements([]int{1, 2, 1, 3, 4, 2})
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "Two elements",
			nums:     []int{2, 3},
			expected: 1,
		},
		{
			name:     "Three elements",
			nums:     []int{2, 3, 1},
			expected: 1,
		},
		{
			name:     "Multiple elements",
			nums:     []int{2, 3, 1, 1, 4},
			expected: 2,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4,
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
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Jump([]int{2, 3, 1, 1, 4})
	}
}

func TestMinimumAddedCoins(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Single coin",
			coins:    []int{1},
			amount:   2,
			expected: 1,
		},
		{
			name:     "Two coins",
			coins:    []int{1, 2},
			amount:   3,
			expected: 0,
		},
		{
			name:     "Three coins",
			coins:    []int{1, 4, 10},
			amount:   19,
			expected: 2,
		},
		{
			name:     "Multiple coins",
			coins:    []int{1, 2, 5, 10},
			amount:   13,
			expected: 1,
		},
		{
			name:     "All same coins",
			coins:    []int{1, 1, 1, 1, 1},
			amount:   10,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MinimumAddedCoins(tt.coins, tt.amount)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinimumAddedCoins(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinimumAddedCoins([]int{1, 2, 5, 10}, 13)
	}
}

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Single word",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "Multiple words",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "No words",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "Empty string",
			s:        "",
			wordDict: []string{"leet", "code"},
			expected: false,
		},
		{
			name:     "Empty dictionary",
			s:        "leetcode",
			wordDict: nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := WordBreak(tt.s, tt.wordDict)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkWordBreak(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = WordBreak("applepenapple", []string{"apple", "pen"})
	}
}

func TestIsValidParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "Single parenthesis",
			s:        "(",
			expected: false,
		},
		{
			name:     "Two parentheses",
			s:        "()",
			expected: true,
		},
		{
			name:     "Multiple parentheses",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "Multiple parentheses with invalid order",
			s:        "(]",
			expected: false,
		},
		{
			name:     "Multiple parentheses with invalid nesting",
			s:        "([)]",
			expected: false,
		},
		{
			name:     "Multiple parentheses with valid nesting",
			s:        "{[]}",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := IsValidParenthesis(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkIsValidParenthesis(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = IsValidParenthesis("{[]}")
	}
}

func TestCanJumpBackwards(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Two elements",
			nums:     []int{1, 0},
			expected: true,
		},
		{
			name:     "Three elements",
			nums:     []int{2, 3, 1},
			expected: true,
		},
		{
			name:     "Multiple elements",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Cannot jump backwards",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CanJumpBackwards(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCanJumpBackwards(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CanJumpBackwards([]int{2, 3, 1, 1, 4})
	}
}

func TestCanJumpForwards(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Single element",
			nums:     []int{0},
			expected: true,
		},
		{
			name:     "Two elements",
			nums:     []int{1, 0},
			expected: true,
		},
		{
			name:     "Three elements",
			nums:     []int{2, 3, 1},
			expected: true,
		},
		{
			name:     "Multiple elements",
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Cannot jump backwards",
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CanJumpForwards(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCanJumpForwards(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CanJumpForwards([]int{2, 3, 1, 1, 4})
	}
}

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{
			name:     "Single character",
			word1:    "a",
			word2:    "b",
			expected: "ab",
		},
		{
			name:     "Two characters",
			word1:    "ab",
			word2:    "cd",
			expected: "acbd",
		},
		{
			name:     "Three characters",
			word1:    "abc",
			word2:    "def",
			expected: "adbecf",
		},
		{
			name:     "Different lengths",
			word1:    "abc",
			word2:    "defg",
			expected: "adbecfg",
		},
		{
			name:     "Empty strings",
			word1:    "",
			word2:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MergeAlternately(tt.word1, tt.word2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMergeAlternately(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MergeAlternately("abcdefgddhdkslfjsldfh", "defgkllskdieldjflsdj")
	}
}

func TestSudoku(t *testing.T) {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	expected := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}

	result := Sudoku(board)

	assert.NotNil(t, result)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Sudoku solver returned incorrect solution.\nExpected:\n%v\nGot:\n%v", formatBoard(expected), formatBoard(result))
	}

}

// Helper function to format the board for better readability in error messages.
func formatBoard(board [][]int) string {
	var sb strings.Builder
	for i, row := range board {
		if i%3 == 0 && i != 0 {
			sb.WriteString("------+-------+------\n")
		}
		for j, num := range row {
			if j%3 == 0 && j != 0 {
				sb.WriteString("| ")
			}
			sb.WriteString(fmt.Sprintf("%d ", num))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			k:        1,
			expected: 2,
		},
		{
			name:     "Three elements",
			nums:     []int{1, 2, 3},
			k:        2,
			expected: 2,
		},
		{
			name:     "Multiple elements",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			k:        1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FindKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			name:     "Single string",
			strs:     []string{"flower"},
			expected: "flower",
		},
		{
			name:     "Two strings with common prefix",
			strs:     []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Two strings with no common prefix",
			strs:     []string{"dog", "racecar", "car"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := LongestCommonPrefix(tt.strs)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LongestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Single element",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "Two elements",
			nums:     []int{0, 1},
			expected: []int{1, 0},
		},
		{
			name:     "Three elements",
			nums:     []int{0, 1, 0},
			expected: []int{1, 0, 0},
		},
		{
			name:     "Multiple elements",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "All zeroes",
			nums:     []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			MoveZeroes(tt.nums)
			assert.Equal(t, tt.expected, tt.nums)
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		nums := []int{0, 1, 0, 3, 12}
		MoveZeroes(nums)
	}
}

func TestMaximumUnits(t *testing.T) {
	tests := []struct {
		name      string
		boxTypes  [][]int
		truckSize int
		expected  int
	}{
		{
			name:      "Single box",
			boxTypes:  [][]int{{1, 3}},
			truckSize: 3,
			expected:  3,
		},
		{
			name:      "Two boxes",
			boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
			truckSize: 10,
			expected:  91,
		},
		{
			name:      "Multiple boxes",
			boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
			truckSize: 4,
			expected:  8,
		},
		{
			name:      "All same boxes",
			boxTypes:  [][]int{{1, 3}, {1, 3}, {1, 3}},
			truckSize: 4,
			expected:  9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MaximumUnits(tt.boxTypes, tt.truckSize)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaximumUnits(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MaximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10)
	}
}

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			k:        1,
			expected: []int{2, 1},
		},
		{
			name:     "Three elements",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: []int{3, 1, 2},
		},
		{
			name:     "Multiple elements",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1, 1},
			k:        2,
			expected: []int{1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RotateArray(tt.nums, tt.k)
			assert.Equal(t, tt.expected, tt.nums)
		})
	}
}

func BenchmarkRotateArray(b *testing.B) {
	b.ReportAllocs()

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < b.N; i++ {
		RotateArray(nums, 3)
	}
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Single number",
			n:        10,
			expected: 4,
		},
		{
			name:     "Two numbers",
			n:        2,
			expected: 0,
		},
		{
			name:     "Multiple numbers",
			n:        20,
			expected: 8,
		},
		{
			name:     "All prime numbers",
			n:        100,
			expected: 25,
		},
		{
			name:     "No prime numbers",
			n:        1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CountPrimes(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCountPrimes(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CountPrimes(100)
	}
}

func TestMinSwaps(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Single element",
			data:     []int{1},
			expected: 0,
		},
		{
			name:     "Two elements",
			data:     []int{1, 0},
			expected: 0,
		},
		{
			name:     "Three elements",
			data:     []int{1, 0, 0},
			expected: 0,
		},
		{
			name:     "Multiple elements",
			data:     []int{1, 0, 0, 0, 1, 0, 1},
			expected: 1,
		},
		{
			name:     "All same elements",
			data:     []int{1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "wrap",
			data:     []int{1, 1, 0, 0, 1},
			expected: 0,
		},
		{
			name:     "Multi Swap",
			data:     []int{0, 1, 1, 1, 0, 0, 1, 1, 0},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MinSwaps(tt.data)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinSwaps(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinSwaps([]int{1, 0, 0, 0, 1, 0, 1})
	}
}

func TestTopKFrequentWords(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		k        int
		expected []string
	}{
		{
			name:     "Single word",
			words:    []string{"i"},
			k:        1,
			expected: []string{"i"},
		},
		{
			name:     "Two words",
			words:    []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:        2,
			expected: []string{"i", "love"},
		},
		{
			name:     "Multiple words",
			words:    []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:        4,
			expected: []string{"the", "is", "sunny", "day"},
		},
		{
			name:     "All same words",
			words:    []string{"i", "i", "i", "i", "i"},
			k:        1,
			expected: []string{"i"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := TopKFrequentWords(tt.words, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTopKFrequentWords(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = TopKFrequentWords([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4)
	}
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "Single digit",
			str:      "1",
			expected: 1,
		},
		{
			name:     "Two digits",
			str:      "42",
			expected: 42,
		},
		{
			name:     "Multiple digits",
			str:      "   -42",
			expected: -42,
		},
		{
			name:     "Leading characters",
			str:      "4193 with words",
			expected: 4193,
		},
		{
			name:     "Trailing characters",
			str:      "words and 987",
			expected: 0,
		},
		{
			name:     "Overflow",
			str:      "-91283472332",
			expected: math.MinInt32,
		},
		{
			name:     "Underflow",
			str:      "91283472332",
			expected: math.MaxInt32,
		},
		{
			name:     "Empty string",
			str:      "",
			expected: 0,
		},
		{
			name:     "Only whitespace",
			str:      "   ",
			expected: 0,
		},
		{
			name:     "Only sign",
			str:      "+",
			expected: 0,
		},
		{
			name:     "Sign and whitespace",
			str:      "  +  ",
			expected: 0,
		},
		{
			name:     "Sign and letters",
			str:      "  +a  ",
			expected: 0,
		},
		{
			name:     "Sign and digits",
			str:      "  +123  ",
			expected: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MyAtoi(tt.str)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MyAtoi("   -42")
	}
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			name:          "Single course",
			numCourses:    1,
			prerequisites: nil,
			expected:      []int{0},
		},
		{
			name:          "Two courses",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			name:          "Multiple courses",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      []int{0, 2, 1, 3},
		},
		{
			name:          "Cycle",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			expected:      []int{},
		},
		{
			name:          "Multiple cycles",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 1}, {3, 2}},
			expected:      []int{},
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: nil,
			expected:      []int{2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindOrder(tt.numCourses, tt.prerequisites)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindOrder(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FindOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	}
}

func TestFindOrderBFS(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			name:          "Single course",
			numCourses:    1,
			prerequisites: nil,
			expected:      []int{0},
		},
		{
			name:          "Two courses",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			name:          "Multiple courses",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			name:          "Cycle",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			expected:      []int{},
		},
		{
			name:          "Multiple cycles",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 1}, {3, 2}},
			expected:      []int{},
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: nil,
			expected:      []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindOrderBFS(tt.numCourses, tt.prerequisites)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindOrderBFS(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FindOrderBFS(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	}
}

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Single character",
			s:        "a",
			expected: 0,
		},
		{
			name:     "Two characters",
			s:        "aa",
			expected: -1,
		},
		{
			name:     "Three characters",
			s:        "aba",
			expected: 1,
		},
		{
			name:     "Multiple characters",
			s:        "leetcode",
			expected: 0,
		},
		{
			name:     "All same characters",
			s:        "cc",
			expected: -1,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FirstUniqChar(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFirstUniqChar(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FirstUniqChar("leetcode")
	}
}

func TestRotateImage(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Single element",
			matrix: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Two elements",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			name: "Three elements",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "Four elements",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RotateImageClockW(tt.matrix)
			assert.Equal(t, tt.expected, tt.matrix)
		})
	}
}

func BenchmarkRotateImage(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		matrix := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		RotateImageClockW(matrix)
	}
}

func TestRotateImageCounterClockW(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Single element",
			matrix: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Two elements",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{2, 4},
				{1, 3},
			},
		},
		{
			name: "Three elements",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			},
		},
		{
			name: "Four elements",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{11, 10, 7, 16},
				{9, 8, 6, 12},
				{1, 4, 3, 14},
				{5, 2, 13, 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RotateImageCounterClockW(tt.matrix)
			assert.Equal(t, tt.expected, tt.matrix)
		})
	}
}

func BenchmarkRotateImageCounterClockW(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		matrix := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		RotateImageCounterClockW(matrix)
	}
}

func TestRotateImageOneEighty(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "Single element",
			matrix: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "Two elements",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{4, 3},
				{2, 1},
			},
		},
		{
			name: "Three elements",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{9, 8, 7},
				{6, 5, 4},
				{3, 2, 1},
			},
		},
		{
			name: "Four elements",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{16, 12, 14, 15},
				{7, 6, 3, 13},
				{10, 8, 4, 2},
				{11, 9, 1, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RotateImageOneEighty(tt.matrix)
			assert.Equal(t, tt.expected, tt.matrix)
		})
	}
}

func BenchmarkRotateImageOneEighty(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		matrix := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		RotateImageOneEighty(matrix)
	}
}

func TestFindRotation(t *testing.T) {
	tests := []struct {
		name     string
		mat      [][]int
		target   [][]int
		expected bool
	}{
		{
			name: "Single element",
			mat: [][]int{
				{1},
			},
			target: [][]int{
				{1},
			},
			expected: true,
		},
		{
			name: "Two elements",
			mat: [][]int{
				{1, 2},
				{3, 4},
			},
			target: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: true,
		},
		{
			name: "Three elements",
			mat: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			target: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: true,
		},
		{
			name: "Four elements",
			mat: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			target: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindRotation(tt.mat, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}

}

func BenchmarkFindRotation(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		matrix := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		target := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		_ = FindRotation(matrix, target)
	}
}

func TestReorganizeString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Single character",
			s:        "a",
			expected: "a",
		},
		{
			name:     "Two characters",
			s:        "aa",
			expected: "",
		},
		{
			name:     "Three characters",
			s:        "aab",
			expected: "aba",
		},
		{
			name:     "Five characters",
			s:        "aaabb",
			expected: "ababa",
		},
		{
			name:     "Six characters",
			s:        "aaabbb",
			expected: "ababab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ReorganizeString(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkReorganizeString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ReorganizeString("aaabbb")
	}
}

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "All positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			k:        9,
			expected: 2,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        -3,
			expected: 2,
		},
		{
			name:     "Mixed positive and negative numbers",
			nums:     []int{-1, 2, -3, 4, -5},
			k:        -3,
			expected: 2,
		},
		{
			name:     "Single element equal to k",
			nums:     []int{5},
			k:        5,
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        5,
			expected: 0,
		},
		{
			name:     "Large array with multiple subarrays",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			k:        3,
			expected: 8,
		},
		{
			name:     "Array with zeros",
			nums:     []int{0, 0, 0, 0, 0},
			k:        0,
			expected: 15,
		},
		{
			name:     "No subarrays sum to k",
			nums:     []int{1, 2, 3, 4, 5},
			k:        20,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := SubarraySum(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSubarraySum(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = SubarraySum([]int{1, 2, 3, 4, 5}, 9)
	}
}

func TestNumberToWords(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected string
	}{
		{
			name:     "Single digit",
			num:      1,
			expected: "One",
		},
		{
			name:     "Two digits",
			num:      12,
			expected: "Twelve",
		},
		{
			name:     "Three digits",
			num:      123,
			expected: "One Hundred Twenty Three",
		},
		{
			name:     "Four digits",
			num:      1234,
			expected: "One Thousand Two Hundred Thirty Four",
		},
		{
			name:     "Five digits",
			num:      12345,
			expected: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name:     "Six digits",
			num:      123456,
			expected: "One Hundred Twenty Three Thousand Four Hundred Fifty Six",
		},
		{
			name:     "Seven digits",
			num:      1234567,
			expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name:     "Eight digits",
			num:      12345678,
			expected: "Twelve Million Three Hundred Forty Five Thousand Six Hundred Seventy Eight",
		},
		{
			name:     "Nine digits",
			num:      123456789,
			expected: "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine",
		},
		{
			name:     "Ten digits",
			num:      1234567890,
			expected: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety",
		},
		{
			name:     "Zero",
			num:      0,
			expected: "Zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := NumberToWords(tt.num)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkNumberToWords(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = NumberToWords(1234567890)
	}
}

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name: "Single interval",
			intervals: [][]int{
				{0, 30},
			},
			expected: 1,
		},
		{
			name: "Two intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
			},
			expected: 2,
		},
		{
			name: "Three intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
			},
			expected: 2,
		},
		{
			name: "Four intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
				{25, 30},
			},
			expected: 2,
		},
		{
			name: "Five intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
				{25, 30},
				{30, 35},
			},
			expected: 2,
		},
		{
			name: "Six intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
				{25, 30},
				{30, 35},
				{35, 40},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MinMeetingRoomsHeap(tt.intervals)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinMeetingRooms(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinMeetingRoomsHeap([][]int{{0, 30}, {5, 10}, {15, 20}, {25, 30}, {30, 35}})
	}
}

func TestMinMeetingRoomsSweep(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{
			name: "Single interval",
			intervals: [][]int{
				{0, 30},
			},
			expected: 1,
		},
		{
			name: "Two intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
			},
			expected: 2,
		},
		{
			name: "Three intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
			},
			expected: 2,
		},
		{
			name: "Four intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
				{25, 30},
			},
			expected: 2,
		},
		{
			name: "Five intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
				{25, 30},
				{30, 35},
			},
			expected: 2,
		},
		{
			name: "Six intervals",
			intervals: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
				{25, 30},
				{30, 35},
				{35, 40},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MinMeetingRoomSweep(tt.intervals)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinMeetingRoomsSweep(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinMeetingRoomSweep([][]int{{0, 30}, {5, 10}, {15, 20}, {25, 30}, {30, 35}})
	}
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "Single course",
			numCourses:    1,
			prerequisites: nil,
			expected:      true,
		},
		{
			name:          "Two courses",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			name:          "Multiple courses",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:          "Cycle",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			expected:      false,
		},
		{
			name:          "Multiple cycles",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 1}, {3, 2}},
			expected:      false,
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: nil,
			expected:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CanFinish(tt.numCourses, tt.prerequisites)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCanFinish(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CanFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	}
}

func TestMaxProfit2(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Single price",
			prices:   []int{7},
			expected: 0,
		},
		{
			name:     "Two prices",
			prices:   []int{7, 1},
			expected: 0,
		},
		{
			name:     "Three prices",
			prices:   []int{7, 1, 5},
			expected: 4,
		},
		{
			name:     "Four prices",
			prices:   []int{7, 1, 5, 3},
			expected: 4,
		},
		{
			name:     "Five prices",
			prices:   []int{7, 1, 5, 3, 6},
			expected: 7,
		},
		{
			name:     "Six prices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "Seven prices",
			prices:   []int{7, 1, 5, 3, 6, 4, 8},
			expected: 11,
		},
		{
			name:     "Eight prices",
			prices:   []int{7, 1, 5, 3, 6, 4, 8, 2},
			expected: 11,
		},
		{
			name:     "Nine prices",
			prices:   []int{7, 1, 5, 3, 6, 4, 8, 2, 10},
			expected: 19,
		},
		{
			name:     "Ten prices",
			prices:   []int{7, 1, 5, 3, 6, 4, 8, 2, 10, 9},
			expected: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := MaxProfit2(tt.prices)
			assert.Equal(t, tt.expected, result)
		})
	}

}

func BenchmarkMaxProfit2(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MaxProfit2([]int{7, 1, 5, 3, 6, 4, 8, 2, 10, 9})
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Empty string",
			s:        "",
			k:        2,
			expected: 0,
		},
		{
			name:     "Single character string",
			s:        "A",
			k:        1,
			expected: 1,
		},
		{
			name:     "All same characters",
			s:        "AAAA",
			k:        2,
			expected: 4,
		},
		{
			name:     "No replacements needed",
			s:        "ABCDE",
			k:        0,
			expected: 1,
		},
		{
			name:     "k greater than string length",
			s:        "ABC",
			k:        5,
			expected: 3,
		},
		{
			name:     "Complex case with multiple character types",
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			name:     "Case with non-consecutive repeating characters",
			s:        "ABCBABBA",
			k:        2,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CharacterReplacement(tt.s, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCharacterReplacement(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CharacterReplacement("ABCBABBA", 2)
	}
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "All oranges already rotten",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "No fresh oranges",
			grid: [][]int{
				{0, 2, 2},
				{0, 0, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "Rot all oranges",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Single row grid",
			grid: [][]int{
				{2, 1, 1, 1, 2},
			},
			expected: 2,
		},
		{
			name: "Single column grid",
			grid: [][]int{
				{2},
				{1},
				{1},
				{1},
				{2},
			},
			expected: 2,
		},
		{
			name: "Large grid",
			grid: [][]int{
				{2, 1, 1, 1, 1, 1},
				{1, 1, 0, 0, 1, 1},
				{1, 0, 1, 1, 0, 1},
				{1, 1, 0, 0, 1, 1},
				{1, 1, 1, 1, 1, 2},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := OrangesRotting(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkOrangesRotting(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = OrangesRotting([][]int{
			{2, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 1, 1},
			{1, 0, 1, 1, 0, 1},
			{1, 1, 0, 0, 1, 1},
			{1, 1, 1, 1, 1, 2},
		})
	}
}

func TestSpiralMatrix2(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			name: "1x1 matrix",
			n:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "3x3 matrix",
			n:    3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "4x4 matrix",
			n:    4,
			expected: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			name:     "0x0 matrix",
			n:        0,
			expected: [][]int{},
		},
		{
			name: "5x5 matrix",
			n:    5,
			expected: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SpiralMatrix2(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SpiralMatrix2(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

func BenchmarkSpiralMatrix2(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = SpiralMatrix2(5)
	}
}

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name:     "4x4 matrix",
			matrix:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			expected: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
		},
		{
			name:     "1x1 matrix",
			matrix:   [][]int{{1}},
			expected: []int{1},
		},
		{
			name:     "2x3 matrix",
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: []int{1, 2, 3, 6, 5, 4},
		},
		{
			name:     "3x2 matrix",
			matrix:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 4, 6, 5, 3},
		},
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := SpiralOrder(tt.matrix)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SpiralOrder() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkSpiralOrder(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = SpiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	}
}

func TestKthFactor(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		expected int
	}{
		{
			name:     "First factor of 12",
			n:        12,
			k:        1,
			expected: 1,
		},
		{
			name:     "Second factor of 12",
			n:        12,
			k:        2,
			expected: 2,
		},
		{
			name:     "Third factor of 12",
			n:        12,
			k:        3,
			expected: 3,
		},
		{
			name:     "Fourth factor of 12",
			n:        12,
			k:        4,
			expected: 4,
		},
		{
			name:     "Fifth factor of 12",
			n:        12,
			k:        5,
			expected: 6,
		},
		{
			name:     "Sixth factor of 12",
			n:        12,
			k:        6,
			expected: 12,
		},
		{
			name:     "K greater than number of factors",
			n:        12,
			k:        7,
			expected: -1,
		},
		{
			name:     "First factor of prime number",
			n:        13,
			k:        1,
			expected: 1,
		},
		{
			name:     "Second factor of prime number",
			n:        13,
			k:        2,
			expected: 13,
		},
		{
			name:     "K greater than number of factors for prime number",
			n:        13,
			k:        3,
			expected: -1,
		},
		{
			name:     "First factor of 1",
			n:        1,
			k:        1,
			expected: 1,
		},
		{
			name:     "K greater than number of factors for 1",
			n:        1,
			k:        2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := KthFactor(tt.n, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkKthFactor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = KthFactor(1000, 3)
	}
}

func TestPartitionStringSingleCharacter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"SingleCharacter", "a", 1},
		{"AllUniqueCharacters", "abcdef", 1},
		{"AllSameCharacters", "aaaaaa", 6},
		{"MixedCharacters", "abac", 2},
		{"LongStringWithRepeats", "abacabadabacaba", 8},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := PartitionString(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkPartitionString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = PartitionString("abacabadabacaba")
	}
}

func TestMinSwapsNoWrap(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{"SingleOne", []int{0, 1, 0, 0, 0}, 0},
		{"AllOnes", []int{1, 1, 1, 1, 1}, 0},
		{"NoOnes", []int{0, 0, 0, 0, 0}, 0},
		{"Mixed", []int{1, 0, 1, 0, 1, 0, 1}, 2},
		{"EdgeCase", []int{1, 0, 0, 1, 0, 1, 0, 1, 0, 1}, 2},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := MinSwapsNoWrap(tt.data)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinSwapsNoWrap(b *testing.B) {
	data := []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	for i := 0; i < b.N; i++ {
		_ = MinSwapsNoWrap(data)
	}
}

func TestAppendCharacters(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{"BothEmpty", "", "", 0},
		{"EmptyS", "", "abc", 3},
		{"EmptyT", "abc", "", 0},
		{"NoAppendNeeded", "abc", "abc", 0},
		{"AppendNeeded", "abc", "abcd", 1},
		{"PartialMatch", "abcde", "ace", 0},
		{"NoMatch", "abc", "def", 3},
		{"LongerS", "abcdefgh", "aceg", 0},
		{"LongerT", "ace", "abcdefgh", 7},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := AppendCharacters(tt.s, tt.t)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkAppendCharacters(b *testing.B) {
	tests := []struct {
		name string
		s    string
		t    string
	}{
		{"BothEmpty", "", ""},
		{"EmptyS", "", "abc"},
		{"EmptyT", "abc", ""},
		{"LongerT", "ace", "abcdefgh"},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = AppendCharacters(tt.s, tt.t)
			}
		})
	}
}

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{"AllUnique", []int{1, 2, 3, 4, 5}, 3, 12},
		{"SomeDuplicates", []int{1, 2, 2, 3, 4}, 3, 9},
		{"AllDuplicates", []int{1, 1, 1, 1, 1}, 2, 0},
		{"EmptyArray", []int{}, 3, 0},
		{"SingleElement", []int{1}, 1, 1},
		{"KGreaterThanArrayLength", []int{1, 2}, 3, 0},
		{"KEqualToArrayLength", []int{1, 2, 3}, 3, 6},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := MaximumSubarraySum(tt.nums, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaximumSubarraySum(b *testing.B) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{"SmallArray", []int{1, 2, 3, 4, 5}, 3},
		{"MediumArray", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{"LargeArray", make([]int, 1000), 500},
		{"AllUnique", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"AllSame", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 5},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MaximumSubarraySum(tt.nums, tt.k)
			}
		})
	}
}

func TestFirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"NoMissingPositive", []int{1, 2, 3}, 4},
		{"SingleElement", []int{1}, 2},
		{"UnorderedElements", []int{3, 4, -1, 1}, 2},
		{"AllNegatives", []int{-1, -2, -3}, 1},
		{"MixedElements", []int{7, 8, 9, 11, 12}, 1},
		{"Duplicates", []int{1, 1, 2, 2}, 3},
		{"EmptyArray", []int{}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := FirstMissingPositive(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFirstMissingPositive(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{"SmallArray", []int{3, 4, -1, 1}},
		{"MediumArray", make([]int, 1000)},
		{"LargeArray", make([]int, 100000)},
		{"AllPositive", []int{1, 2, 3, 4, 5}},
		{"AllNegative", []int{-1, -2, -3, -4, -5}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = FirstMissingPositive(tt.nums)
			}
		})
	}
}

func TestNumberOfWaysSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{"AllZeros", "0000", 0},
		{"AllOnes", "1111", 0},
		{"Alternating01", "010101", 8},
		{"Random", "001101", 6},
		{"Alternating10", "101010", 8},
		{"MixedPattern", "1001001", 12},
		{"SingleZero", "0", 0},
		{"SingleOne", "1", 0},
		{"EmptyString", "", 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NumberOfWaysSlice(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkNumberOfWaysSlice(b *testing.B) {
	tests := []struct {
		name string
		s    string
	}{
		{"SmallString", "010101"},
		{"MediumString", "010101010101010101010101010101"},
		{"LargeString", strings.Repeat("01", 1000)},
		{"AllZeros", strings.Repeat("0", 1000)},
		{"AllOnes", strings.Repeat("1", 1000)},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NumberOfWaysSlice(tt.s)
			}
		})
	}
}

func TestNumberOfWaysCounters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{"AllZeros", "0000", 0},
		{"AllOnes", "1111", 0},
		{"Alternating01", "010101", 8},
		{"Random", "001101", 6},
		{"Alternating10", "101010", 8},
		{"MixedPattern", "1001001", 12},
		{"SingleZero", "0", 0},
		{"SingleOne", "1", 0},
		{"EmptyString", "", 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NumberOfWaysCounters(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkNumberOfWaysCounters(b *testing.B) {
	tests := []struct {
		name string
		s    string
	}{
		{"SmallString", "010101"},
		{"MediumString", "010101010101010101010101010101"},
		{"LargeString", strings.Repeat("01", 1000)},
		{"AllZeros", strings.Repeat("0", 1000)},
		{"AllOnes", strings.Repeat("1", 1000)},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NumberOfWaysCounters(tt.s)
			}
		})
	}
}

func TestReorderLogFiles(t *testing.T) {
	tests := []struct {
		name     string
		logs     []string
		expected []string
	}{
		{
			name: "mixed logs",
			logs: []string{
				"dig1 8 1 5 1",
				"let1 art can",
				"dig2 3 6",
				"let2 own kit dig",
				"let3 art zero",
			},
			expected: []string{
				"let1 art can",
				"let3 art zero",
				"let2 own kit dig",
				"dig1 8 1 5 1",
				"dig2 3 6",
			},
		},
		{
			name: "all digit logs",
			logs: []string{
				"dig1 8 1 5 1",
				"dig2 3 6",
			},
			expected: []string{
				"dig1 8 1 5 1",
				"dig2 3 6",
			},
		},
		{
			name: "all letter logs",
			logs: []string{
				"let1 art can",
				"let2 own kit dig",
				"let3 art zero",
			},
			expected: []string{
				"let1 art can",
				"let3 art zero",
				"let2 own kit dig",
			},
		},
		{
			name: "logs with same content",
			logs: []string{
				"let1 art can",
				"let2 art can",
			},
			expected: []string{
				"let1 art can",
				"let2 art can",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := ReorderLogFiles(tt.logs)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkReorderLogFiles(b *testing.B) {
	tests := []struct {
		name string
		logs []string
	}{
		{"SmallSet", []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}},
		{"MediumSet", []string{
			"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero",
			"let4 art zero", "let5 art zero", "let6 art zero", "let7 art zero", "let8 art zero",
		}},
		{"LargeSet", []string{
			"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero",
			"let4 art zero", "let5 art zero", "let6 art zero", "let7 art zero", "let8 art zero",
			"let9 art zero", "let10 art zero", "let11 art zero", "let12 art zero", "let13 art zero",
			"let14 art zero", "let15 art zero", "let16 art zero", "let17 art zero", "let18 art zero",
		}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ReorderLogFiles(tt.logs)
			}
		})
	}
}

func TestCountTheNumOfKFreeSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int64
	}{
		{
			name:     "example case",
			nums:     []int{1, 2, 3, 4},
			k:        1,
			expected: 8,
		},
		{
			name:     "no k-free subsets",
			nums:     []int{1, 2, 3},
			k:        10,
			expected: 1,
		},
		{
			name:     "single element",
			nums:     []int{5},
			k:        1,
			expected: 1,
		},
		{
			name:     "large k value",
			nums:     []int{1, 2, 3, 4, 5},
			k:        100,
			expected: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := CountTheNumOfKFreeSubsets(tt.nums, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCountTheNumOfKFreeSubsets(b *testing.B) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{"SmallInput", []int{1, 2, 3, 4}, 1},
		{"MediumInput", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2},
		{"LargeInput", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, 3},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CountTheNumOfKFreeSubsets(tt.nums, tt.k)
			}
		})
	}
}

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"AllNumbersPresent", []int{0, 1, 2, 3, 4, 5}, 6},
		{"SingleElementMissing", []int{3, 0, 1}, 2},
		{"FirstElementMissing", []int{1, 2, 3, 4, 5}, 0},
		{"LastElementMissing", []int{0, 1, 2, 3, 4}, 5},
		{"EmptyArray", []int{}, 0},
		{"LargeArray", func() []int {
			nums := make([]int, 999999)
			for i := 0; i < 999999; i++ {
				nums[i] = i
			}
			return nums
		}(), 999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MissingNumber(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMissingNumber(b *testing.B) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"AllNumbersPresent", []int{0, 1, 2, 3, 4, 5}, 6},
		{"SingleElementMissing", []int{3, 0, 1}, 2},
		{"FirstElementMissing", []int{1, 2, 3, 4, 5}, 0},
		{"LastElementMissing", []int{0, 1, 2, 3, 4}, 5},
		{"EmptyArray", []int{}, 0},
		{"LargeArray", func() []int {
			nums := make([]int, 999999)
			for i := 0; i < 999999; i++ {
				nums[i] = i
			}
			return nums
		}(), 999999},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MissingNumber(tt.nums)
			}
		})
	}
}

func TestMissingNumberMath(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"AllNumbersPresent", []int{0, 1, 2, 3, 4, 5}, 6},
		{"SingleElementMissing", []int{3, 0, 1}, 2},
		{"FirstElementMissing", []int{1, 2, 3, 4, 5}, 0},
		{"LastElementMissing", []int{0, 1, 2, 3, 4}, 5},
		{"EmptyArray", []int{}, 0},
		{"LargeArray", func() []int {
			nums := make([]int, 999999)
			for i := 0; i < 999999; i++ {
				nums[i] = i
			}
			return nums
		}(), 999999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MissingNumberMath(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMissingNumberMath(b *testing.B) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"AllNumbersPresent", []int{0, 1, 2, 3, 4, 5}, 6},
		{"SingleElementMissing", []int{3, 0, 1}, 2},
		{"FirstElementMissing", []int{1, 2, 3, 4, 5}, 0},
		{"LastElementMissing", []int{0, 1, 2, 3, 4}, 5},
		{"EmptyArray", []int{}, 0},
		{"LargeArray", func() []int {
			nums := make([]int, 999999)
			for i := 0; i < 999999; i++ {
				nums[i] = i
			}
			return nums
		}(), 999999},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MissingNumberMath(tt.nums)
			}
		})
	}
}

func TestFindDuplicateFloyd(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleDuplicate",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "MultipleDuplicates",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name: "LargeArray",
			nums: func() []int {
				nums := make([]int, 100000)
				for i := 0; i < 99999; i++ {
					nums[i] = i + 1
				}
				nums[99999] = 99999
				return nums
			}(),
			expected: 99999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FindDuplicateFloyd(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFindDuplicateFloyd(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{"SingleDuplicate", []int{1, 3, 4, 2, 2}},
		{"MultipleDuplicates", []int{3, 1, 3, 4, 2}},
		{"LargeArray", func() []int {
			nums := make([]int, 100000)
			for i := 0; i < 99999; i++ {
				nums[i] = i + 1
			}
			nums[99999] = 99999
			return nums
		}()},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = FindDuplicateFloyd(tt.nums)
			}
		})
	}
}

func TestFindDuplicateBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleDuplicate",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "MultipleDuplicates",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name: "LargeArray",
			nums: func() []int {
				nums := make([]int, 100000)
				for i := 0; i < 99999; i++ {
					nums[i] = i + 1
				}
				nums[99999] = 99999
				return nums
			}(),
			expected: 99999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FindDuplicateBinarySearch(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFindDuplicateBinarySearch(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{"SingleDuplicate", []int{1, 3, 4, 2, 2}},
		{"MultipleDuplicates", []int{3, 1, 3, 4, 2}},
		{"LargeArray", func() []int {
			nums := make([]int, 100000)
			for i := 0; i < 99999; i++ {
				nums[i] = i + 1
			}
			nums[99999] = 99999
			return nums
		}()},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = FindDuplicateBinarySearch(tt.nums)
			}
		})
	}
}

func TestMinimumKeypresses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"EmptyString", "", 0},
		{"SingleCharacter", "a", 1},
		{"AllUniqueCharacters", "abcdefghijklmnopqrstuvwxyz", 51},
		{"AllSameCharacter", "aaaaa", 5},
		{"MixedCharacters", "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz", 102},
		{"HighFrequencyCharacters", "aaaabbbbccccddddeeeeffffgggghhhhiiii", 36},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := MinimumKeypresses(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinimumKeypresses(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		{"EmptyString", ""},
		{"SingleCharacter", "a"},
		{"AllUniqueCharacters", "abcdefghijklmnopqrstuvwxyz"},
		{"RepeatingCharacters", "aaabbbccc"},
		{"LongString", "thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsoverthelazydog"},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MinimumKeypresses(tt.input)
			}
		})
	}
}

func TestMinCost(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		costs    []int
		expected int64
	}{
		{
			name:     "SingleElementArray",
			nums:     []int{5},
			costs:    []int{10},
			expected: 0,
		},
		{
			name:     "IncreasingSequence",
			nums:     []int{1, 2, 3, 4, 5},
			costs:    []int{1, 2, 3, 4, 5},
			expected: 14,
		},
		{
			name:     "DecreasingSequence",
			nums:     []int{5, 4, 3, 2, 1},
			costs:    []int{5, 4, 3, 2, 1},
			expected: 10,
		},
		{
			name:     "MixedSequence",
			nums:     []int{3, 1, 4, 1, 5},
			costs:    []int{2, 7, 1, 8, 2},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := MinCost(tt.nums, tt.costs)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinCost(b *testing.B) {
	tests := []struct {
		name  string
		nums  []int
		costs []int
	}{
		{
			name:  "SingleElementArray",
			nums:  []int{5},
			costs: []int{10},
		},
		{
			name:  "IncreasingSequence",
			nums:  []int{1, 2, 3, 4, 5},
			costs: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "DecreasingSequence",
			nums:  []int{5, 4, 3, 2, 1},
			costs: []int{5, 4, 3, 2, 1},
		},
		{
			name:  "MixedSequence",
			nums:  []int{3, 1, 4, 1, 5},
			costs: []int{2, 7, 1, 8, 2},
		},
		{
			name:  "EmptyArray",
			nums:  []int{},
			costs: []int{},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MinCost(tt.nums, tt.costs)
			}
		})
	}
}

func TestMakePalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"PalindromeWithNoChangesNeeded", "racecar", true},
		{"PalindromeWithOneChangeNeeded", "raceecar", true},
		{"PalindromeWithTwoChangesNeeded", "raccecar", true},
		{"NotPalindromeWithMoreThanTwoChanges", "hiello", false},
		{"EmptyStringIsPalindrome", "", true},
		{"SingleCharacterStringIsPalindrome", "a", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MakePalindrome(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMakePalindrome(b *testing.B) {
	tests := []struct {
		name  string
		input string
	}{
		{"PalindromeWithNoChangesNeeded", "racecar"},
		{"PalindromeWithOneChangeNeeded", "raceecar"},
		{"PalindromeWithTwoChangesNeeded", "raccecar"},
		{"NotPalindromeWithMoreThanTwoChanges", "hello"},
		{"EmptyStringIsPalindrome", ""},
		{"SingleCharacterStringIsPalindrome", "a"},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MakePalindrome(tt.input)
			}
		})
	}
}

func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"MinimumOperationsSingleElement", []int{1}, 0},
		{"MinimumOperationsAlreadyPalindrome", []int{1, 2, 3, 2, 1}, 0},
		{"MinimumOperationsTwoElements", []int{1, 2}, 1},
		{"MinimumOperationsGeneralCase", []int{1, 3, 2, 1, 2}, 4},
		{"MinimumOperationsAllSameElements", []int{2, 2, 2, 2}, 0},
		{"MinimumOperationsEmptyArray", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MinimumOperations(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinimumOperations(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{"MinimumOperationsSingleElement", []int{1}},
		{"MinimumOperationsAlreadyPalindrome", []int{1, 2, 3, 2, 1}},
		{"MinimumOperationsTwoElements", []int{1, 2}},
		{"MinimumOperationsGeneralCase", []int{1, 3, 2, 1, 2}},
		{"MinimumOperationsAllSameElements", []int{2, 2, 2, 2}},
		{"MinimumOperationsEmptyArray", []int{}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MinimumOperations(tt.nums)
			}
		})
	}
}

func TestMinimumSwaps(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"SingleElementArray", []int{1}, 0},
		{"AlreadySortedArray", []int{1, 2, 3, 4, 5}, 0},
		{"ReverseSortedArray", []int{5, 4, 3, 2, 1}, 7},
		{"RandomArray", []int{3, 1, 5, 2, 4}, 3},
		{"ArrayWithDuplicates", []int{2, 3, 2, 1, 4}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := MinimumSwaps(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinimumSwaps(b *testing.B) {
	tests := []struct {
		name string
		nums []int
	}{
		{"SingleElementArray", []int{1}},
		{"AlreadySortedArray", []int{1, 2, 3, 4, 5}},
		{"ReverseSortedArray", []int{5, 4, 3, 2, 1}},
		{"RandomArray", []int{3, 1, 5, 2, 4}},
		{"ArrayWithDuplicates", []int{2, 3, 2, 1, 4}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MinimumSwaps(tt.nums)
			}
		})
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "AllStationsReachable",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "NotEnoughGas",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "SingleStation",
			gas:      []int{5},
			cost:     []int{4},
			expected: 0,
		},
		{
			name:     "MultipleStations",
			gas:      []int{2, 3, 4, 5, 1},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "ExactGas",
			gas:      []int{1, 2, 3, 4},
			cost:     []int{1, 2, 3, 4},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := CanCompleteCircuit(tt.gas, tt.cost)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCanCompleteCircuit(b *testing.B) {
	tests := []struct {
		name string
		gas  []int
		cost []int
	}{
		{
			name: "AllStationsReachable",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
		},
		{
			name: "NotEnoughGas",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
		},
		{
			name: "SingleStation",
			gas:  []int{5},
			cost: []int{4},
		},
		{
			name: "MultipleStations",
			gas:  []int{2, 3, 4, 5, 1},
			cost: []int{3, 4, 5, 1, 2},
		},
		{
			name: "ExactGas",
			gas:  []int{1, 2, 3, 4},
			cost: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CanCompleteCircuit(tt.gas, tt.cost)
			}
		})
	}
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		words    []string
		expected []string
	}{
		{
			name: "FindWordsWithNoMatchingWords",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:    []string{"word"},
			expected: []string{},
		},
		{
			name: "FindWordsWithSingleWordMatch",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:    []string{"ab"},
			expected: []string{"ab"},
		},
		{
			name: "FindWordsWithMultipleWordsMatch",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			words:    []string{"ab", "cd"},
			expected: []string{"ab", "cd"},
		},
		{
			name: "FindWordsWithOverlappingWords",
			board: [][]byte{
				{'a', 'b', 'c'},
				{'d', 'e', 'f'},
				{'g', 'h', 'i'},
			},
			words:    []string{"abc", "cfi", "beh", "defi"},
			expected: []string{"abc", "cfi", "beh", "defi"},
		},
		{
			name: "FindWordsWithRepeatedCharacters",
			board: [][]byte{
				{'a', 'a'},
				{'a', 'a'},
			},
			words:    []string{"aaa", "aaaa"},
			expected: []string{"aaa", "aaaa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := FindWords(tt.board, tt.words)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkFindWords(b *testing.B) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
	}{
		{
			name:  "BenchmarkSmallBoard",
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			words: []string{"ab", "cd"},
		},
		{
			name: "BenchmarkLargeBoard",
			board: [][]byte{
				{'a', 'b', 'c', 'd', 'e'},
				{'f', 'g', 'h', 'i', 'j'},
				{'k', 'l', 'm', 'n', 'o'},
				{'p', 'q', 'r', 's', 't'},
				{'u', 'v', 'w', 'x', 'y'},
			},
			words: []string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz"},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = FindWords(tt.board, tt.words)
			}
		})
	}
}

func TestSnakesAndLadders(t *testing.T) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}
	expected := 4
	result := SnakesAndLadders(board)
	assert.Equal(t, expected, result)
}

func BenchmarkSnakesAndLadders(b *testing.B) {
	board := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	for i := 0; i < b.N; i++ {
		_ = SnakesAndLadders(board)
	}
}

func TestExist(t *testing.T) {
	tests := []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			board: [][]byte{
				{'A'},
			},
			word:     "B",
			expected: false,
		},
		{
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABCD",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("board: %v, word: %s", tt.board, tt.word), func(t *testing.T) {
			result := Exist(tt.board, tt.word)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkExist(b *testing.B) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"

	for i := 0; i < b.N; i++ {
		_ = Exist(board, word)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{n: 1, expected: []string{"()"}},
		{n: 2, expected: []string{"(())", "()()"}},
		{n: 3, expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{n: 0, expected: []string{""}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			t.Parallel()
			result := GenerateParenthesis(tt.n)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	n := 3
	for i := 0; i < b.N; i++ {
		_ = GenerateParenthesis(n)
	}
}

func TestUpdateBoard(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		click    []int
		expected [][]byte
	}{
		{
			name: "MineClicked",
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{1, 2},
			expected: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'X', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
		},
		{
			name: "EmptyCellWithNoAdjacentMines",
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{0, 0},
			expected: [][]byte{
				{'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			name: "EmptyCellWithAdjacentMines",
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{1, 1},
			expected: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', '1', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
		},
		{
			name: "ClickOnBoundary",
			board: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{0, 4},
			expected: [][]byte{
				{'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
				{'B', 'B', 'B', 'B', 'B'},
			},
		},
		{
			name: "ClickOnAlreadyRevealedCell",
			board: [][]byte{
				{'B', '1', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			click: []int{0, 0},
			expected: [][]byte{
				{'B', '1', 'E', 'E', 'E'},
				{'E', 'E', 'M', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'E', 'E', 'E', 'E'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := UpdateBoard(tt.board, tt.click)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkUpdateBoard(b *testing.B) {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	click := []int{1, 2}

	for i := 0; i < b.N; i++ {
		_ = UpdateBoard(board, click)
	}
}

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"LongestValidParenthesesEmptyString", "", 0},
		{"LongestValidParenthesesNoValidPairs", "(((((", 0},
		{"LongestValidParenthesesSinglePair", "()", 2},
		{"LongestValidParenthesesNestedPairs", "((()))", 6},
		{"LongestValidParenthesesMixedPairs", "(()())", 6},
		{"LongestValidParenthesesUnbalancedLeft", "(()", 2},
		{"LongestValidParenthesesUnbalancedRight", "())", 2},
		{"LongestValidParenthesesComplex", "(()))())(", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LongestValidParentheses(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestValidParentheses(b *testing.B) {
	input := "(()))())("
	for i := 0; i < b.N; i++ {
		_ = LongestValidParentheses(input)
	}
}

func TestLongestValidParenthesesTwoPass(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"LongestValidParenthesesEmptyString", "", 0},
		{"LongestValidParenthesesNoValidPairs", "(((((", 0},
		{"LongestValidParenthesesSinglePair", "()", 2},
		{"LongestValidParenthesesNestedPairs", "((()))", 6},
		{"LongestValidParenthesesMixedPairs", "(()())", 6},
		{"LongestValidParenthesesUnbalancedLeft", "(()", 2},
		{"LongestValidParenthesesUnbalancedRight", "())", 2},
		{"LongestValidParenthesesComplex", "(()))())(", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LongestValidParenthesesTwoPass(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestValidParenthesesTwoPass(b *testing.B) {
	input := "(()))())("
	for i := 0; i < b.N; i++ {
		_ = LongestValidParenthesesTwoPass(input)
	}
}

func TestGameOfLife(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]int
		expected [][]int
	}{
		{
			name: "GameOfLifeAllDead",
			board: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "GameOfLifeAllAlive",
			board: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "GameOfLifeMixed",
			board: [][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			expected: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
		},
		{
			name: "GameOfLifeEdgeCase",
			board: [][]int{
				{1, 1},
				{1, 0},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			GameOfLife(tt.board)
			assert.Equal(t, tt.expected, tt.board)
		})
	}
}

func BenchmarkGameOfLife(b *testing.B) {
	board := [][]int{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}

	for i := 0; i < b.N; i++ {
		GameOfLife(board)
	}
}

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "SingleElement",
			grid:     [][]int{{5}},
			expected: 5,
		},
		{
			name:     "SingleRow",
			grid:     [][]int{{1, 2, 3}},
			expected: 6,
		},
		{
			name:     "SingleColumn",
			grid:     [][]int{{1}, {2}, {3}},
			expected: 6,
		},
		{
			name: "MultipleRowsAndColumns",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinPathSum(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinPathSum(b *testing.B) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	for i := 0; i < b.N; i++ {
		_ = MinPathSum(grid)
	}
}

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{"SingleChild", []int{5}, 1},
		{"TwoChildrenIncreasing", []int{1, 2}, 3},
		{"TwoChildrenDecreasing", []int{2, 1}, 3},
		{"MultipleChildren", []int{1, 0, 2}, 5},
		{"AllSameRatings", []int{1, 1, 1}, 3},
		{"ComplexRatings", []int{1, 2, 2}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Candy(tt.ratings)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCandy(b *testing.B) {
	ratings := []int{1, 0, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		_ = Candy(ratings)
	}
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{"NoElevation", []int{}, 0},
		{"SingleElevation", []int{1}, 0},
		{"TwoElevations", []int{1, 2}, 0},
		{"SimpleCase", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"FlatSurface", []int{3, 3, 3, 3}, 0},
		{"DescendingSurface", []int{4, 3, 2, 1}, 0},
		{"AscendingSurface", []int{1, 2, 3, 4}, 0},
		{"ComplexCase", []int{4, 2, 0, 3, 2, 5}, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Trap(tt.height)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTrap(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	for i := 0; i < b.N; i++ {
		_ = Trap(height)
	}
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name     string
		heights  [][]int
		expected [][]int
	}{
		{
			name:     "SingleCell",
			heights:  [][]int{{1}},
			expected: [][]int{{0, 0}},
		},
		{
			name: "FlatGrid",
			heights: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 1}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},
		{
			name: "ComplexGrid",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			expected: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
		{
			name: "NoFlow",
			heights: [][]int{
				{10, 10, 10},
				{10, 1, 10},
				{10, 10, 10},
			},
			expected: [][]int{
				{0, 0}, {0, 1}, {0, 2},
				{1, 0}, {1, 2},
				{2, 0}, {2, 1}, {2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := PacificAtlantic(tt.heights)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkPacificAtlantic(b *testing.B) {
	heights := [][]int{
		{1, 2, 2, 3, 5},
		{3, 2, 3, 4, 4},
		{2, 4, 5, 3, 1},
		{6, 7, 1, 4, 5},
		{5, 1, 1, 2, 4},
	}
	for i := 0; i < b.N; i++ {
		_ = PacificAtlantic(heights)
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{"AllZeros", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"AllOnes", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
		{"AllTwos", []int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
		{"Mixed", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"Empty", []int{}, []int{}},
		{"SingleElement", []int{1}, []int{1}},
		{"AlreadySorted", []int{0, 0, 1, 1, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			SortColors(tt.nums)
			assert.Equal(t, tt.expected, tt.nums)
		})
	}
}

func BenchmarkSortColors(b *testing.B) {
	nums := []int{2, 0, 2, 1, 1, 0}
	for i := 0; i < b.N; i++ {
		SortColors(nums)
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"ValidInput", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"SingleSegment", "0000", []string{"0.0.0.0"}},
		{"InvalidLength", "123", []string{}},
		{"LeadingZero", "010010", []string{"0.10.0.10", "0.100.1.0"}},
		{"MaxLength", "111111111111", []string{"111.111.111.111"}},
		{"EmptyInput", "", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := RestoreIpAddresses(tt.input)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkRestoreIpAddresses(b *testing.B) {
	input := "25525511135"
	for i := 0; i < b.N; i++ {
		RestoreIpAddresses(input)
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{"ValidInput", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"SingleElement", []int{1}, 1, []int{1}},
		{"EmptyInput", []int{}, 3, []int{}},
		{"KGreaterThanLength", []int{1, 2, 3}, 4, []int{}},
		{"KIsZero", []int{1, 2, 3}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MaxSlidingWindow(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	for i := 0; i < b.N; i++ {
		MaxSlidingWindow(nums, k)
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"PositiveNumber", 123, 321},
		{"NegativeNumber", -123, -321},
		{"NumberWithTrailingZeros", 120, 21},
		{"SingleDigitNumber", 5, 5},
		{"Zero", 0, 0},
		{"OverflowPositive", 1534236469, 0},
		{"OverflowNegative", -1534236469, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Reverse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	input := 123
	for i := 0; i < b.N; i++ {
		Reverse(input)
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"LongestConsecutiveSequence", []int{100, 4, 200, 1, 3, 2}, 4},
		{"SingleElementSequence", []int{1}, 1},
		{"EmptyArraySequence", []int{}, 0},
		{"NoConsecutiveSequence", []int{10, 5, 100}, 1},
		{"NegativeNumbersSequence", []int{-1, -2, -3, -4, 0, 1}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LongestConsecutive(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	nums := []int{100, 4, 200, 1, 3, 2}
	for i := 0; i < b.N; i++ {
		LongestConsecutive(nums)
	}
}

func TestNumSquares(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"SinglePerfectSquare", 16, 1},
		{"SumOfTwoPerfectSquares", 13, 2},
		{"SumOfThreePerfectSquares", 12, 3},
		{"Zero", 0, 0},
		{"NonPerfectSquare", 15, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := NumSquares(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkNumSquares(b *testing.B) {
	n := 16
	for i := 0; i < b.N; i++ {
		NumSquares(n)
	}
}

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		{"SingleTask", []byte{'A'}, 2, 1},
		{"MultipleTasksNoCooldown", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{"MultipleTasksWithCooldown", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{"AllUniqueTasks", []byte{'A', 'B', 'C', 'D', 'E', 'F'}, 2, 6},
		{"EmptyTasks", []byte{}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LeastInterval(tt.tasks, tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLeastInterval(b *testing.B) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	for i := 0; i < b.N; i++ {
		LeastInterval(tasks, n)
	}
}

func TestKClosest(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		k        int
		expected [][]int
	}{
		{"SinglePoint", [][]int{{1, 2}}, 1, [][]int{{1, 2}}},
		{"MultiplePoints", [][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{"KGreaterThanPoints", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 5, [][]int{{3, 3}, {5, -1}, {-2, 4}}},
		{"KEqualToPoints", [][]int{{1, 3}, {-2, 2}, {2, -2}}, 3, [][]int{{1, 3}, {-2, 2}, {2, -2}}},
		{"EmptyPoints", [][]int{}, 1, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := KClosest(tt.points, tt.k)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkKClosest(b *testing.B) {
	points := [][]int{{1, 3}, {-2, 2}, {2, -2}}
	k := 2
	for i := 0; i < b.N; i++ {
		KClosest(points, k)
	}
}

func TestPeakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected int
	}{
		{"PeakIndexInMiddleOfMountainArray", []int{0, 2, 3, 4, 5, 3, 1}, 4},
		{"PeakIndexAtStartOfMountainArray", []int{5, 4, 3, 2, 1}, 0},
		{"PeakIndexAtEndOfMountainArray", []int{1, 2, 3, 4, 5}, 4},
		{"PeakIndexInSingleElementArray", []int{1}, 0},
		{"PeakIndexInTwoElementArray", []int{1, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := PeakIndexInMountainArray(tt.arr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkPeakIndexInMountainArray(b *testing.B) {
	arr := []int{0, 2, 3, 4, 5, 3, 1}
	for i := 0; i < b.N; i++ {
		PeakIndexInMountainArray(arr)
	}
}

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{"LongestSubstringWithSingleCharacter", "aaaaa", 1, 5},
		{"LongestSubstringWithNoRepeatingCharacters", "abcdef", 2, 0},
		{"LongestSubstringWithMixedCharacters", "aaabb", 3, 3},
		{"LongestSubstringWithMultipleValidSubstrings", "ababbc", 2, 5},
		{"LongestSubstringWithEmptyString", "", 3, 0},
		{"LongestSubstringWithSingleCharacterString", "a", 1, 1},
		{"LongestSubstringWithAllCharactersBelowK", "aabbcc", 3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LongestSubstring(tt.s, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestSubstring(b *testing.B) {
	s := "ababbc"
	k := 2
	for i := 0; i < b.N; i++ {
		LongestSubstring(s, k)
	}
}

func TestWordBreakII(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected []string
	}{
		{"WordBreakIIWithSingleWord", "leetcode", []string{"leet", "code"}, []string{"leet code"}},
		{"WordBreakIIWithMultipleWords", "catsanddog", []string{"cat", "cats", "and", "sand", "dog"}, []string{"cats and dog", "cat sand dog"}},
		{"WordBreakIIWithNoSolution", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := WordBreakII(tt.s, tt.wordDict)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkWordBreakII(b *testing.B) {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	for i := 0; i < b.N; i++ {
		WordBreakII(s, wordDict)
	}
}

func TestSmallestDistancePair(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"SmallestDistancePairWithMultiplePairs", []int{1, 3, 1}, 1, 0},
		{"SmallestDistancePairWithSingleElement", []int{1}, 1, 0},
		{"SmallestDistancePairWithLargeK", []int{1, 6, 1}, 3, 5},
		{"SmallestDistancePairWithNegativeNumbers", []int{-1, -3, -1}, 1, 0},
		{"SmallestDistancePairWithDuplicates", []int{1, 1, 1}, 2, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SmallestDistancePair(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSmallestDistancePair(b *testing.B) {
	nums := []int{1, 3, 1}
	k := 1
	for i := 0; i < b.N; i++ {
		SmallestDistancePair(nums, k)
	}
}

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		target   string
		original []byte
		changed  []byte
		cost     []int
		expected int64
	}{
		{
			name:     "ValidTransformation",
			source:   "abc",
			target:   "def",
			original: []byte{'a', 'b', 'c'},
			changed:  []byte{'d', 'e', 'f'},
			cost:     []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "InvalidTransformation",
			source:   "abc",
			target:   "xyz",
			original: []byte{'a', 'b', 'c'},
			changed:  []byte{'d', 'e', 'f'},
			cost:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "SameSourceAndTarget",
			source:   "abc",
			target:   "abc",
			original: []byte{'a', 'b', 'c'},
			changed:  []byte{'d', 'e', 'f'},
			cost:     []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "EmptySourceAndTarget",
			source:   "",
			target:   "",
			original: []byte{},
			changed:  []byte{},
			cost:     []int{},
			expected: 0,
		},
		{
			name:     "PartialTransformation",
			source:   "abc",
			target:   "abf",
			original: []byte{'c'},
			changed:  []byte{'f'},
			cost:     []int{3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinimumCost(tt.source, tt.target, tt.original, tt.changed, tt.cost)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinimumCost(b *testing.B) {
	source := "abc"
	target := "def"
	original := []byte{'a', 'b', 'c'}
	changed := []byte{'d', 'e', 'f'}
	cost := []int{1, 2, 3}

	for i := 0; i < b.N; i++ {
		MinimumCost(source, target, original, changed, cost)
	}
}

func TestLongestValidSubstring(t *testing.T) {
	tests := []struct {
		name      string
		word      string
		forbidden []string
		expected  int
	}{
		{
			name:      "NoForbiddenWords",
			word:      "abcde",
			forbidden: []string{},
			expected:  5,
		},
		{
			name:      "WithForbiddenWords",
			word:      "abcde",
			forbidden: []string{"bc"},
			expected:  3,
		},
		{
			name:      "AllForbidden",
			word:      "abcde",
			forbidden: []string{"a", "b", "c", "d", "e"},
			expected:  0,
		},
		{
			name:      "PartialForbidden",
			word:      "abcde",
			forbidden: []string{"ab", "de"},
			expected:  3,
		},
		{
			name:      "EmptyWord",
			word:      "",
			forbidden: []string{"a", "b"},
			expected:  0,
		},
		{
			name:      "LongForbidden",
			word:      "abcdefghij",
			forbidden: []string{"abcdefghij"},
			expected:  9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LongestValidSubstring(tt.word, tt.forbidden)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestValidSubstring(b *testing.B) {
	word := "abcdefghij"
	forbidden := []string{"abcdefghij"}

	for i := 0; i < b.N; i++ {
		LongestValidSubstring(word, forbidden)
	}
}

func TestFindMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		machines []int
		expected int
	}{
		{
			name:     "EqualDistribution",
			machines: []int{1, 0, 5},
			expected: 3,
		},
		{
			name:     "Distribution",
			machines: []int{0, 3, 0},
			expected: 2,
		},
		{
			name:     "AlreadyBalanced",
			machines: []int{1, 1, 1},
			expected: 0,
		},
		{
			name:     "SingleMachine",
			machines: []int{4},
			expected: 0,
		},
		{
			name:     "LargeDifference",
			machines: []int{0, 0, 11, 5},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := FindMinMoves(tt.machines)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindMinMoves(b *testing.B) {
	machines := []int{0, 0, 11, 5}

	for i := 0; i < b.N; i++ {
		FindMinMoves(machines)
	}
}

func TestUniqueLetterString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "SingleCharacter",
			s:        "A",
			expected: 1,
		},
		{
			name:     "AllUnique",
			s:        "ABC",
			expected: 10,
		},
		{
			name:     "RepeatingCharacters",
			s:        "AAB",
			expected: 6,
		},
		{
			name:     "EmptyString",
			s:        "",
			expected: 0,
		},
		{
			name:     "LongString",
			s:        "ABABABAB",
			expected: 28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := UniqueLetterString(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkUniqueLetterString(b *testing.B) {
	s := "ABABABAB"

	for i := 0; i < b.N; i++ {
		UniqueLetterString(s)
	}
}

func TestFindAllConcatenatedWordsInADict(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected []string
	}{
		{
			name:     "EmptyInput",
			words:    []string{},
			expected: []string{},
		},
		{
			name:     "NoConcatenatedWords",
			words:    []string{"cat", "dog", "mouse"},
			expected: []string{},
		},
		{
			name:     "SingleConcatenatedWord",
			words:    []string{"cat", "dog", "catdog"},
			expected: []string{"catdog"},
		},
		{
			name:     "MultipleConcatenatedWords",
			words:    []string{"cat", "dog", "catdog", "dogcat", "mouse", "catdogmouse"},
			expected: []string{"catdog", "dogcat", "catdogmouse"},
		},
		{
			name:     "NestedConcatenatedWords",
			words:    []string{"cat", "dog", "catdog", "catdogcat", "catdogcatdog"},
			expected: []string{"catdog", "catdogcat", "catdogcatdog"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := FindAllConcatenatedWordsInADict(tt.words)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkFindAllConcatenatedWordsInADict(b *testing.B) {
	words := []string{"cat", "dog", "catdog", "dogcat", "mouse", "catdogmouse"}

	for i := 0; i < b.N; i++ {
		FindAllConcatenatedWordsInADict(words)
	}
}

func TestFindAllConcatenatedWordsInADictTrie(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected []string
	}{
		{
			name:     "EmptyInput",
			words:    []string{},
			expected: []string{},
		},
		{
			name:     "NoConcatenatedWords",
			words:    []string{"cat", "dog", "mouse"},
			expected: []string{},
		},
		{
			name:     "SingleConcatenatedWord",
			words:    []string{"cat", "dog", "catdog"},
			expected: []string{"catdog"},
		},
		{
			name:     "MultipleConcatenatedWords",
			words:    []string{"cat", "dog", "catdog", "dogcat", "mouse", "catdogmouse"},
			expected: []string{"catdog", "dogcat", "catdogmouse"},
		},
		{
			name:     "NestedConcatenatedWords",
			words:    []string{"cat", "dog", "catdog", "catdogcat", "catdogcatdog"},
			expected: []string{"catdog", "catdogcat", "catdogcatdog"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := FindAllConcatenatedWordsInADictTrie(tt.words)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func BenchmarkFindAllConcatenatedWordsInADicTriet(b *testing.B) {
	words := []string{"cat", "dog", "catdog", "dogcat", "mouse", "catdogmouse"}

	for i := 0; i < b.N; i++ {
		FindAllConcatenatedWordsInADict(words)
	}
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "SingleElement",
			heights:  []int{5},
			expected: 5,
		},
		{
			name:     "MultipleElements",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "AllSameHeight",
			heights:  []int{3, 3, 3, 3},
			expected: 12,
		},
		{
			name:     "DecreasingHeights",
			heights:  []int{6, 5, 4, 3, 2, 1},
			expected: 12,
		},
		{
			name:     "SinglePeak",
			heights:  []int{2, 1, 4, 5, 1, 3, 3},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LargestRectangleArea(tt.heights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLargestRectangleArea(b *testing.B) {
	heights := []int{2, 1, 5, 6, 2, 3}

	for i := 0; i < b.N; i++ {
		LargestRectangleArea(heights)
	}
}

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		name     string
		fruits   []int
		expected int
	}{
		{
			name:     "SingleType",
			fruits:   []int{1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "TwoTypes",
			fruits:   []int{1, 2, 1, 2, 1, 2},
			expected: 6,
		},
		{
			name:     "ThreeTypes",
			fruits:   []int{1, 2, 3, 2, 2},
			expected: 4,
		},
		{
			name:     "Empty",
			fruits:   []int{},
			expected: 0,
		},
		{
			name:     "SingleElement",
			fruits:   []int{1},
			expected: 1,
		},
		{
			name:     "AlternatingTypes",
			fruits:   []int{1, 2, 1, 2, 1, 2, 1, 2},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := TotalFruit(tt.fruits)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTotalFruit(b *testing.B) {
	fruits := []int{1, 2, 1, 2, 1, 2, 1, 2}

	for i := 0; i < b.N; i++ {
		TotalFruit(fruits)
	}
}

func TestReversePairs(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleElement",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "NoPairs",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "WithPairs",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 4,
		},
		{
			name:     "Mixed",
			nums:     []int{1, 3, 2, 3, 1},
			expected: 2,
		},
		{
			name:     "Duplicates",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ReversePairs(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkReversePairs(b *testing.B) {
	nums := []int{1, 3, 2, 3, 1}

	for i := 0; i < b.N; i++ {
		ReversePairs(nums)
	}
}

func TestGetMaxLen(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "PositiveNumbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "NegativeNumbers",
			nums:     []int{-1, -2, -3, -4, -5},
			expected: 4,
		},
		{
			name:     "MixedNumbers",
			nums:     []int{1, -2, -3, 4, -5, 6},
			expected: 4,
		},
		{
			name:     "SingleZero",
			nums:     []int{0},
			expected: 0,
		},
		{
			name:     "MultipleZeros",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "AlternatingSigns",
			nums:     []int{1, -1, 1, -1, 1},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := GetMaxLen(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkGetMaxLen(b *testing.B) {
	nums := []int{1, -2, -3, 4, -5, 6}

	for i := 0; i < b.N; i++ {
		GetMaxLen(nums)
	}
}

func TestMaximumBooks(t *testing.T) {
	tests := []struct {
		name     string
		books    []int
		expected int64
	}{
		{
			name:     "SingleBook",
			books:    []int{5},
			expected: 5,
		},
		{
			name:     "IncreasingOrder",
			books:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "DecreasingOrder",
			books:    []int{5, 4, 3, 2, 1},
			expected: 25,
		},
		{
			name:     "MixedOrder",
			books:    []int{3, 1, 4, 1, 5},
			expected: 26,
		},
		{
			name:     "AllSame",
			books:    []int{2, 2, 2, 2, 2},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MaximumBooks(tt.books)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaximumBooks(b *testing.B) {
	books := []int{3, 1, 4, 1, 5}

	for i := 0; i < b.N; i++ {
		MaximumBooks(books)
	}
}

func TestReachNumber(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		expected int
	}{
		{
			name:     "PositiveTarget",
			target:   5,
			expected: 5,
		},
		{
			name:     "NegativeTarget",
			target:   -5,
			expected: 5,
		},
		{
			name:     "ZeroTarget",
			target:   0,
			expected: 0,
		},
		{
			name:     "LargeTarget",
			target:   1000,
			expected: 47,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ReachNumber(tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkReachNumber(b *testing.B) {
	target := 1000

	for i := 0; i < b.N; i++ {
		ReachNumber(target)
	}
}

func TestClosedIsland(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "SingleIsland",
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1},
				{1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1},
				{1, 1, 1, 1, 1},
			},
			expected: 1,
		},
		{
			name: "MultipleIslands",
			grid: [][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			expected: 2,
		},
		{
			name: "NoIslands",
			grid: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: 0,
		},
		{
			name: "EdgeCaseSmallGrid",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			expected: 0,
		},
		{
			name: "EdgeCaseSingleCell",
			grid: [][]int{
				{0},
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ClosedIsland(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkClosedIsland(b *testing.B) {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	}

	for i := 0; i < b.N; i++ {
		ClosedIsland(grid)
	}
}

func TestShortestBridge(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "SingleBridge",
			grid: [][]int{
				{0, 1, 0, 0, 0},
				{0, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expected: 1,
		},
		{
			name: "MultipleBridges",
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "NoBridgeNeeded",
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expected: -1,
		},
		{
			name: "EdgeCaseSingleCell",
			grid: [][]int{
				{0},
			},
			expected: -1,
		},
		{
			name: "EdgeCaseSmallGrid",
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ShortestBridge(tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkShortestBridge(b *testing.B) {
	grid := [][]int{
		{0, 1, 0, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	for i := 0; i < b.N; i++ {
		ShortestBridge(grid)
	}
}

func TestMostExpensiveItem(t *testing.T) {
	tests := []struct {
		name     string
		primeOne int
		primeTwo int
		expected int
	}{
		{
			name:     "Basic",
			primeOne: 3,
			primeTwo: 5,
			expected: 7,
		},
		{
			name:     "SamePrimes",
			primeOne: 7,
			primeTwo: 7,
			expected: 35,
		},
		{
			name:     "LargePrimes",
			primeOne: 101,
			primeTwo: 103,
			expected: 10199,
		},
		{
			name:     "EdgeCaseOne",
			primeOne: 2,
			primeTwo: 3,
			expected: 1,
		},
		{
			name:     "EdgeCaseTwo",
			primeOne: 2,
			primeTwo: 2,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MostExpensiveItem(tt.primeOne, tt.primeTwo)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMostExpensiveItem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MostExpensiveItem(101, 103)
	}
}

func TestMostExpensiveItemDP(t *testing.T) {
	tests := []struct {
		name     string
		primeOne int
		primeTwo int
		expected int
	}{
		{
			name:     "Basic",
			primeOne: 3,
			primeTwo: 5,
			expected: 7,
		},
		{
			name:     "LargePrimes",
			primeOne: 101,
			primeTwo: 103,
			expected: 10199,
		},
		{
			name:     "EdgeCaseOne",
			primeOne: 2,
			primeTwo: 3,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MostExpensiveItemDP(tt.primeOne, tt.primeTwo)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMostExpensiveItemDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MostExpensiveItem(101, 103)
	}
}

func TestSubArrayRanges(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		{"SingleElement", []int{1}, 0},
		{"TwoElements", []int{1, 2}, 1},
		{"ThreeElements", []int{1, 2, 3}, 4},
		{"AllSameElements", []int{2, 2, 2}, 0},
		{"MixedElements", []int{1, 3, 2, 4}, 12},
		{"DescendingOrder", []int{4, 3, 2, 1}, 10},
		{"EmptyArray", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SubArrayRanges(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSubArrayRanges(b *testing.B) {
	nums := []int{1, 3, 2, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		SubArrayRanges(nums)
	}
}

func TestMinSwapsBinaryString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"EqualZeroAndOne", "010101", 0},
		{"UnequalZeroAndOne", "000111", 1},
		{"UnequalZeroAndOneReturnsNegativeOne", "000011", -1},
		{"AllZeros", "0000", -1},
		{"AllOnes", "1111", -1},
		{"SingleZero", "0", 0},
		{"SingleOne", "1", 0},
		{"EmptyString", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinSwapsBinaryString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinSwapsBinaryString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinSwapsBinaryString("010101")
	}
}

func TestMinimumOperationsArrayZero(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"AllUniqueElements", []int{1, 2, 3, 4, 5}, 5},
		{"AllSameElements", []int{1, 1, 1, 1, 1}, 1},
		{"MixedElements", []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}, 4},
		{"EmptyArray", []int{}, 0},
		{"SingleElement", []int{1}, 1},
		{"ZeroIncluded", []int{0, 1, 2, 3}, 3},
		{"LargeNumbers", []int{100, 99, 98, 97}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinimumOperationsArrayZero(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinimumOperationsArrayZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinimumOperationsArrayZero([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func TestSequentialDigits(t *testing.T) {
	tests := []struct {
		name     string
		low      int
		high     int
		expected []int
	}{
		{"RangeWithinSingleDigit", 10, 12, []int{12}},
		{"RangeWithinDoubleDigits", 10, 100, []int{12, 23, 34, 45, 56, 67, 78, 89}},
		{"RangeCrossingDoubleDigits", 10, 123, []int{12, 23, 34, 45, 56, 67, 78, 89, 123}},
		{"RangeWithinTripleDigits", 100, 300, []int{123, 234}},
		{"RangeWithSingleSequentialDigit", 123, 123, []int{123}},
		{"RangeWithLargeNumbers", 1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := SequentialDigits(tt.low, tt.high)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkSequentialDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialDigits(10, 10000)
	}
}

func TestTriangularSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"SingleElement", []int{5}, 5},
		{"MultipleElements", []int{1, 2, 3, 4, 5}, 8},
		{"AllZeros", []int{0, 0, 0, 0}, 0},
		{"LargeNumbers", []int{9, 9, 9, 9}, 2},
		{"MixedNumbers", []int{3, 6, 1, 8, 2}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := TriangularSum(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkTriangularSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		TriangularSum(nums)
	}
}

func TestGoodDaysToRobBank(t *testing.T) {
	tests := []struct {
		name     string
		security []int
		time     int
		expected []int
	}{
		{"AllDays", []int{1, 2, 3, 4, 5}, 0, []int{0, 1, 2, 3, 4}},
		{"SingleDay", []int{5, 3, 3, 3, 5, 6, 2}, 2, []int{2, 3}},
		{"EdgeCase", []int{1}, 0, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := GoodDaysToRobBank(tt.security, tt.time)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkGoodDaysToRobBank(b *testing.B) {
	security := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	time := 2
	for i := 0; i < b.N; i++ {
		GoodDaysToRobBank(security, time)
	}
}

func TestMinMovesToMakePalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"PalindromeAlready", "racecar", 0},
		{"SingleCharacter", "a", 0},
		{"MultipleMoves", "aabb", 2},
		{"ComplexPalindrome", "abcba", 0},
		{"EmptyString", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinMovesToMakePalindrome(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinMovesToMakePalindrome(b *testing.B) {
	input := "aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz"
	for i := 0; i < b.N; i++ {
		MinMovesToMakePalindrome(input)
	}
}

func TestRacecar(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		expected int
	}{
		{"RacecarReachesTarget", 3, 2},
		{"RacecarReachesTargetWithHigherValue", 6, 5},
		{"RacecarReachesTargetWithZero", 0, 0},
		{"RacecarReachesTargetWithNegativeValue", -1, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Racecar(tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkRacecar(b *testing.B) {
	target := 10000
	for i := 0; i < b.N; i++ {
		Racecar(target)
	}
}

func TestAppealSum(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int64
	}{
		{"SingleCharacter", "a", 1},
		{"AllUniqueCharacters", "abc", 10},
		{"RepeatingCharacters", "abbca", 28},
		{"EmptyString", "", 0},
		{"LongString", "abcdefghijklmnopqrstuvwxyz", 3276},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := AppealSum(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkAppealSum(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		AppealSum(s)
	}
}

func TestMinimumHealth(t *testing.T) {
	tests := []struct {
		name     string
		damage   []int
		armor    int
		expected int64
	}{
		{"SingleDamage", []int{10}, 5, 6},
		{"MultipleDamage", []int{5, 10, 15}, 10, 21},
		{"ArmorGreaterThanMaxDamage", []int{5, 10, 15}, 20, 16},
		{"ArmorZero", []int{5, 10, 15}, 0, 31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := MinimumHealth(tt.damage, tt.armor)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinimumHealth(b *testing.B) {
	damage := []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}
	armor := 25
	for i := 0; i < b.N; i++ {
		MinimumHealth(damage, armor)
	}
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{"EvaluateSingleNumber", []string{"3"}, 3},
		{"EvaluateSimpleAddition", []string{"2", "3", "+"}, 5},
		{"EvaluateSimpleSubtraction", []string{"5", "3", "-"}, 2},
		{"EvaluateSimpleMultiplication", []string{"2", "3", "*"}, 6},
		{"EvaluateSimpleDivision", []string{"6", "3", "/"}, 2},
		{"EvaluateComplexExpression", []string{"2", "1", "+", "3", "*"}, 9},
		{"EvaluateExpressionWithNegativeResult", []string{"4", "13", "5", "/", "+"}, 6},
		{"EvaluateExpressionWithMultipleOperators", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := EvalRPN(tt.tokens)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkEvalRPN(b *testing.B) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	for i := 0; i < b.N; i++ {
		EvalRPN(tokens)
	}
}

func TestNthUglyNumber(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"FirstUglyNumber", 1, 1},
		{"TenthUglyNumber", 10, 12},
		{"FiftiethUglyNumber", 50, 243},
		{"HundredthUglyNumber", 100, 1536},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := NthUglyNumber(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkNthUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NthUglyNumber(1000)
	}
}

func TestInsertIntervals(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			name:        "InsertNewIntervalAtBeginning",
			intervals:   [][]int{{3, 5}, {7, 9}},
			newInterval: []int{1, 2},
			expected:    [][]int{{1, 2}, {3, 5}, {7, 9}},
		},
		{
			name:        "InsertNewIntervalAtEnd",
			intervals:   [][]int{{1, 2}, {3, 5}},
			newInterval: []int{6, 7},
			expected:    [][]int{{1, 2}, {3, 5}, {6, 7}},
		},
		{
			name:        "InsertNewIntervalOverlapping",
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			name:        "InsertNewIntervalMergingMultiple",
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:        "InsertNewIntervalEmptyIntervals",
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
		},
		{
			name:        "InsertNewIntervalSingleInterval",
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 3},
			expected:    [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := Insert(tt.intervals, tt.newInterval)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkInsertIntervals(b *testing.B) {
	intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval := []int{4, 8}
	for i := 0; i < b.N; i++ {
		Insert(intervals, newInterval)
	}
}
