package top50

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

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

			result := MinMeetingRooms(tt.intervals)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMinMeetingRooms(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}, {25, 30}, {30, 35}})
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

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name     string
		begin    string
		end      string
		wordList []string
		expected int
	}{
		{
			name:     "NoTransformationNeeded",
			begin:    "hit",
			end:      "hit",
			wordList: []string{"hit"},
			expected: 1,
		},
		{
			name:     "SingleStepTransformation",
			begin:    "hit",
			end:      "hot",
			wordList: []string{"hot"},
			expected: 2,
		},
		{
			name:     "MultipleStepsTransformation",
			begin:    "hit",
			end:      "cog",
			wordList: []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected: 5,
		},
		{
			name:     "NoPossibleTransformation",
			begin:    "hit",
			end:      "cog",
			wordList: []string{"hot", "dot", "dog", "lot", "log"},
			expected: 0,
		},
		{
			name:     "EmptyWordList",
			begin:    "hit",
			end:      "cog",
			wordList: []string{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := LadderLength(tt.begin, tt.end, tt.wordList)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLadderLength(b *testing.B) {
	tests := []struct {
		name     string
		begin    string
		end      string
		wordList []string
	}{
		{
			name:     "NoTransformationNeeded",
			begin:    "hit",
			end:      "hit",
			wordList: []string{"hit"},
		},
		{
			name:     "SingleStepTransformation",
			begin:    "hit",
			end:      "hot",
			wordList: []string{"hot"},
		},
		{
			name:     "MultipleStepsTransformation",
			begin:    "hit",
			end:      "cog",
			wordList: []string{"hot", "dot", "dog", "lot", "log", "cog"},
		},
		{
			name:     "NoPossibleTransformation",
			begin:    "hit",
			end:      "cog",
			wordList: []string{"hot", "dot", "dog", "lot", "log"},
		},
		{
			name:     "EmptyWordList",
			begin:    "hit",
			end:      "cog",
			wordList: []string{},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = LadderLength(tt.begin, tt.end, tt.wordList)
			}
		})
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
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SmallestDistancePair([]int{1, 3, 1}, 1)
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
