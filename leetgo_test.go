package leetgo

import (
	"reflect"
	"sort"
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
