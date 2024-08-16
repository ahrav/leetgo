package seventyfive

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestGcd(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Two equal numbers",
			a:        5,
			b:        5,
			expected: 5,
		},
		{
			name:     "Two different numbers",
			a:        5,
			b:        10,
			expected: 5,
		},
		{
			name:     "Two prime numbers",
			a:        7,
			b:        11,
			expected: 1,
		},
		{
			name:     "Two large prime numbers",
			a:        7919,
			b:        7907,
			expected: 1,
		},
		{
			name:     "Two large numbers with common factor",
			a:        7919,
			b:        15838,
			expected: 7919,
		},
		{
			name:     "Two large numbers with no common factor",
			a:        7919,
			b:        15839,
			expected: 1,
		},
		{
			name:     "Two large numbers with common factor",
			a:        15838,
			b:        7919,
			expected: 7919,
		},
		{
			name:     "Two large numbers with no common factor",
			a:        15839,
			b:        7919,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Gcd(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkGcd(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Gcd(15838, 7919)
	}
}

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		name     string
		str1     string
		str2     string
		expected string
	}{
		{
			name:     "Two empty strings",
			str1:     "",
			str2:     "",
			expected: "",
		},
		{
			name:     "Single character strings",
			str1:     "A",
			str2:     "A",
			expected: "A",
		},
		{
			name:     "Two different single character strings",
			str1:     "A",
			str2:     "B",
			expected: "",
		},
		{
			name:     "Two equal strings",
			str1:     "ABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			name:     "Two different strings",
			str1:     "ABC",
			str2:     "DEF",
			expected: "",
		},
		{
			name:     "Two strings with common factor",
			str1:     "ABCABC",
			str2:     "ABC",
			expected: "ABC",
		},
		{
			name:     "Two strings with no common factor",
			str1:     "ABCDEF",
			str2:     "ABC",
			expected: "",
		},
		{
			name:     "Two strings with no common factor",
			str1:     "ABCDEF",
			str2:     "DEF",
			expected: "",
		},
		{
			name:     "Two strings with common factor",
			str1:     "ABCDEFABCDEF",
			str2:     "ABCDEF",
			expected: "ABCDEF",
		},
		{
			name:     "Two strings with common factor",
			str1:     "ABCDEFABCDEF",
			str2:     "DEF",
			expected: "",
		},
		{
			name:     "Two strings with common factor",
			str1:     "ABCDEFABCDEF",
			str2:     "ABC",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := GcdOfStrings(tt.str1, tt.str2)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkGcdOfStrings(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = GcdOfStrings("ABCDEFABCDEF", "ABCDEF")
	}
}

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			name:         "Single kid with all candies",
			candies:      []int{5},
			extraCandies: 5,
			expected:     []bool{true},
		},
		{
			name:         "Two kids with equal candies",
			candies:      []int{5, 5},
			extraCandies: 5,
			expected:     []bool{true, true},
		},
		{
			name:         "Two kids with different candies",
			candies:      []int{5, 10},
			extraCandies: 5,
			expected:     []bool{true, true},
		},
		{
			name:         "Three kids with different candies",
			candies:      []int{2, 3, 5},
			extraCandies: 1,
			expected:     []bool{false, false, true},
		},
		{
			name:         "Four kids with different candies",
			candies:      []int{2, 3, 5, 1},
			extraCandies: 3,
			expected:     []bool{true, true, true, false},
		},
		{
			name:         "Five kids with different candies",
			candies:      []int{2, 3, 5, 1, 7},
			extraCandies: 2,
			expected:     []bool{false, false, true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := KidsWithCandies(tt.candies, tt.extraCandies)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkKidsWithCandies(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = KidsWithCandies([]int{2, 3, 5, 1, 7}, 2)
	}
}

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			name:      "Empty flowerbed",
			flowerbed: []int{},
			n:         1,
			expected:  true,
		},
		{
			name:      "Single flowerbed",
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Two flowerbeds",
			flowerbed: []int{0, 0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Three flowerbeds",
			flowerbed: []int{0, 0, 0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Four flowerbeds",
			flowerbed: []int{0, 0, 0, 0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Five flowerbeds",
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         1,
			expected:  true,
		},
		{
			name:      "Six flowerbeds",
			flowerbed: []int{0, 0, 0, 0, 0, 0},
			n:         1,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := CanPlaceFlowers(tt.flowerbed, tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCanPlaceFlowers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CanPlaceFlowers([]int{0, 0, 0, 0, 0}, 1)
	}
}

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Empty string",
			s:        "",
			expected: "",
		},
		{
			name:     "Single character string",
			s:        "A",
			expected: "A",
		},
		{
			name:     "Two character string",
			s:        "AB",
			expected: "AB",
		},
		{
			name:     "Three character string",
			s:        "ABC",
			expected: "ABC",
		},
		{
			name:     "Three character string with vowels",
			s:        "ABC",
			expected: "ABC",
		},
		{
			name:     "Three character string with vowels",
			s:        "ABE",
			expected: "EBA",
		},
		{
			name:     "Four character string with vowels",
			s:        "ABCE",
			expected: "EBCA",
		},
		{
			name:     "Five character string with vowels",
			s:        "ABCDE",
			expected: "EBCDA",
		},
		{
			name:     "Six character string with vowels",
			s:        "ABCDEF",
			expected: "EBCDAF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ReverseVowels(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkReverseVowels(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ReverseVowels("ABCDEF")
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Empty string",
			s:        "",
			expected: "",
		},
		{
			name:     "Single word string",
			s:        "Hello",
			expected: "Hello",
		},
		{
			name:     "Two word string",
			s:        "Hello World",
			expected: "World Hello",
		},
		{
			name:     "Three word string",
			s:        "The quick   brown fox",
			expected: "fox brown quick The",
		},
		{
			name:     "Four word string",
			s:        "  The quick brown fox jumps  ",
			expected: "jumps fox brown quick The",
		},
		{
			name:     "Five word string",
			s:        "  The quick brown fox jumps over",
			expected: "over jumps fox brown quick The",
		},
		{
			name:     "Six word string",
			s:        "The quick brown fox jumps over the  ",
			expected: "the over jumps fox brown quick The",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ReverseWords(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkReverseWords(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = ReverseWords("The quick brown fox jumps over the")
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

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "Empty s string",
			s:        "",
			t:        "abc",
			expected: true,
		},
		{
			name:     "Empty t string",
			s:        "abc",
			t:        "",
			expected: false,
		},
		{
			name:     "Single character strings",
			s:        "a",
			t:        "a",
			expected: true,
		},
		{
			name:     "Two character strings",
			s:        "ab",
			t:        "abc",
			expected: true,
		},
		{
			name:     "Two character strings",
			s:        "ac",
			t:        "abc",
			expected: true,
		},
		{
			name:     "Two character strings",
			s:        "bc",
			t:        "abc",
			expected: true,
		},
		{
			name:     "Two character strings",
			s:        "cb",
			t:        "abc",
			expected: false,
		},
		{
			name:     "Three character strings",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "Three character strings",
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			name:     "Three character strings",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "Three character strings",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "Four character strings",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := IsSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkIsSubsequence(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = IsSubsequence("abc", "ahbgdc")
	}
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			k:        1,
			expected: 1.0,
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			k:        1,
			expected: 2.0,
		},
		{
			name:     "Three elements",
			nums:     []int{1, 2, 3},
			k:        1,
			expected: 3.0,
		},
		{
			name:     "Four elements",
			nums:     []int{1, 2, 3, 4},
			k:        1,
			expected: 4.0,
		},
		{
			name:     "Five elements",
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: 5.0,
		},
		{
			name:     "Six elements",
			nums:     []int{1, 2, 3, 4, 5, 6},
			k:        1,
			expected: 6.0,
		},
		{
			name:     "Seven elements",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        1,
			expected: 7.0,
		},
		{
			name:     "Eight elements",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:        1,
			expected: 8.0,
		},
		{
			name:     "Nine elements",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			k:        1,
			expected: 9.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := FindMaxAverage(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindMaxAverage(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FindMaxAverage([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1)
	}
}

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

func TestMaxOperationsTwoPointer(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 3, 4, 3},
			target:   6,
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 1, 3, 4, 3},
			target:   7,
			expected: 1,
		},
		{
			name:     "Example 4",
			nums:     []int{2, 2, 2, 3, 3},
			target:   4,
			expected: 1,
		},
		{
			name:     "Example 5",
			nums:     []int{1, 1, 1, 1},
			target:   2,
			expected: 2,
		},
		{
			name:     "Example 6",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 7",
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: 2,
		},
		{
			name:     "Example 8",
			nums:     []int{1, 2, 3, 4, 5},
			target:   7,
			expected: 2,
		},
		{
			name:     "Example 9",
			nums:     []int{1, 2, 3, 4, 5},
			target:   8,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MaxOperationsTwoPointer(tt.nums, tt.target)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaxOperationsTwoPointer(b *testing.B) {
	b.ReportAllocs()

	large := make([]int, 10_000)
	for i := range large {
		large[i] = rand.Intn(1_000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MaxOperationsTwoPointer(large, 500)
	}
}

func TestMaxOperationsComplement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 3, 4, 3},
			target:   6,
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{3, 1, 3, 4, 3},
			target:   7,
			expected: 1,
		},
		{
			name:     "Example 4",
			nums:     []int{2, 2, 2, 3, 3},
			target:   4,
			expected: 1,
		},
		{
			name:     "Example 5",
			nums:     []int{1, 1, 1, 1},
			target:   2,
			expected: 2,
		},
		{
			name:     "Example 6",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 2,
		},
		{
			name:     "Example 7",
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: 2,
		},
		{
			name:     "Example 8",
			nums:     []int{1, 2, 3, 4, 5},
			target:   7,
			expected: 2,
		},
		{
			name:     "Example 9",
			nums:     []int{1, 2, 3, 4, 5},
			target:   8,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MaxOperationsComplement(tt.nums, tt.target)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaxOperationsComplement(b *testing.B) {
	b.ReportAllocs()

	large := make([]int, 10_000)
	for i := range large {
		large[i] = rand.Intn(1_000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MaxOperationsComplement(large, 500)
	}
}

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
		{
			name:     "Example 4",
			nums:     []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
			expected: false,
		},
		{
			name:     "Example 5",
			nums:     []int{1, 2, 3, 1, 2, 1, 2, 1, 2, 1},
			expected: true,
		},
		{
			name:     "Example 6",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: true,
		},
		{
			name:     "Example 7",
			nums:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Example 8",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: true,
		},
		{
			name:     "Example 9",
			nums:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := IncreasingTriplet(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkIncreasingTriplet(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = IncreasingTriplet([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func TestIncreasingTripletDP(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "Example 2",
			nums:     []int{5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Example 3",
			nums:     []int{2, 1, 5, 0, 4, 6},
			expected: true,
		},
		{
			name:     "Example 4",
			nums:     []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
			expected: false,
		},
		{
			name:     "Example 5",
			nums:     []int{1, 2, 3, 1, 2, 1, 2, 1, 2, 1},
			expected: true,
		},
		{
			name:     "Example 6",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: true,
		},
		{
			name:     "Example 7",
			nums:     []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: false,
		},
		{
			name:     "Example 8",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: true,
		},
		{
			name:     "Example 9",
			nums:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := IncreasingTripletDP(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkIncreasingTripletDP(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = IncreasingTripletDP([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		chars    []byte
		expected int
	}{
		{
			name:     "Single character",
			chars:    []byte{'a'},
			expected: 1,
		},
		{
			name:     "Two different characters",
			chars:    []byte{'a', 'b'},
			expected: 2,
		},
		{
			name:     "Three repeated characters",
			chars:    []byte{'a', 'a', 'a'},
			expected: 2,
		},
		{
			name:     "Mixed repeated and single characters",
			chars:    []byte{'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 5,
		},
		{
			name:     "Long sequence of repeated characters",
			chars:    []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			expected: 3,
		},
		{
			name:     "Alternating characters",
			chars:    []byte{'a', 'b', 'a', 'b', 'a', 'b'},
			expected: 6,
		},
		{
			name:     "Empty input",
			chars:    []byte{},
			expected: 0,
		},
		{
			name:     "Multiple groups of repeated characters",
			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd', 'd', 'd', 'd'},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Compress(tt.chars)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCompress(b *testing.B) {
	b.ReportAllocs()

	input := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd', 'd', 'd', 'd', 'e', 'e', 'e', 'e', 'e'}

	for i := 0; i < b.N; i++ {
		_ = Compress(input)
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
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			k:        0,
			expected: 5,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			k:        3,
			expected: 3,
		},
		{
			name:     "Alternating ones and zeros",
			nums:     []int{1, 0, 1, 0, 1},
			k:        1,
			expected: 3,
		},
		{
			name:     "Multiple groups of ones",
			nums:     []int{1, 1, 0, 0, 1, 1, 1, 0, 1},
			k:        2,
			expected: 7,
		},
		{
			name:     "K greater than number of zeros",
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        10,
			expected: 19,
		},
		{
			name:     "K equals zero",
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        0,
			expected: 4,
		},
		{
			name:     "Single element array",
			nums:     []int{0},
			k:        1,
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := LongestOnes(tt.nums, tt.k)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkLongestOnes(b *testing.B) {
	b.ReportAllocs()

	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2

	for i := 0; i < b.N; i++ {
		_ = LongestOnes(nums, k)
	}
}

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			name:     "No difference",
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2, 3},
			expected: [][]int{{}, {}},
		},
		{
			name:     "All elements different",
			nums1:    []int{1, 2, 3},
			nums2:    []int{4, 5, 6},
			expected: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:     "Some elements different",
			nums1:    []int{1, 2, 3, 4},
			nums2:    []int{3, 4, 5, 6},
			expected: [][]int{{1, 2}, {5, 6}},
		},
		{
			name:     "Empty nums1",
			nums1:    []int{},
			nums2:    []int{1, 2, 3},
			expected: [][]int{{}, {1, 2, 3}},
		},
		{
			name:     "Empty nums2",
			nums1:    []int{1, 2, 3},
			nums2:    []int{},
			expected: [][]int{{1, 2, 3}, {}},
		},
		{
			name:     "Both empty",
			nums1:    []int{},
			nums2:    []int{},
			expected: [][]int{{}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := FindDifference(tt.nums1, tt.nums2)
			assert.ElementsMatch(t, tt.expected[0], actual[0])
			assert.ElementsMatch(t, tt.expected[1], actual[1])
		})
	}
}

func BenchmarkFindDifference(b *testing.B) {
	b.ReportAllocs()

	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{3, 4, 5, 6}

	for i := 0; i < b.N; i++ {
		_ = FindDifference(nums1, nums2)
	}
}

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "All vowels",
			s:        "aeiou",
			k:        3,
			expected: 3,
		},
		{
			name:     "No vowels",
			s:        "bcdfg",
			k:        2,
			expected: 0,
		},
		{
			name:     "Mixed characters",
			s:        "abciiidef",
			k:        3,
			expected: 3,
		},
		{
			name:     "Single character",
			s:        "a",
			k:        1,
			expected: 1,
		},
		{
			name:     "Empty string",
			s:        "",
			k:        1,
			expected: 0,
		},
		{
			name:     "K greater than string length",
			s:        "aeiou",
			k:        5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MaxVowels(tt.s, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaxVowels(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MaxVowels("abciiidef", 3)
	}
}

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Single zero",
			nums:     []int{1, 0, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Multiple zeros",
			nums:     []int{1, 0, 1, 0, 1},
			expected: 2,
		},
		{
			name:     "Leading zero",
			nums:     []int{0, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Trailing zero",
			nums:     []int{1, 1, 1, 1, 0},
			expected: 4,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element one",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "Single element zero",
			nums:     []int{0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := LongestSubarray(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLongestSubarray(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LongestSubarray([]int{1, 0, 1, 0, 1})
	}
}

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name     string
		gain     []int
		expected int
	}{
		{
			name:     "All positive gains",
			gain:     []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "All negative gains",
			gain:     []int{-1, -2, -3, -4},
			expected: 0,
		},
		{
			name:     "Mixed gains",
			gain:     []int{-5, 1, 5, 0, -7},
			expected: 1,
		},
		{
			name:     "Single element positive",
			gain:     []int{5},
			expected: 5,
		},
		{
			name:     "Single element negative",
			gain:     []int{-5},
			expected: 0,
		},
		{
			name:     "Empty gain array",
			gain:     []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := LargestAltitude(tt.gain)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkLargestAltitude(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = LargestAltitude([]int{-5, 1, 5, 0, -7})
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example case",
			nums:     []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			name:     "No pivot index",
			nums:     []int{1, 2, 3},
			expected: -1,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -1, 0, 1, 1},
			expected: 0,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := PivotIndex(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkPivotIndex(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = PivotIndex([]int{1, 7, 3, 6, 5, 6})
	}
}

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "All unique occurrences",
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			name:     "Non-unique occurrences",
			arr:      []int{1, 2, 2, 1, 1, 2},
			expected: false,
		},
		{
			name:     "Single element",
			arr:      []int{1},
			expected: true,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			expected: true,
		},
		{
			name:     "All elements same",
			arr:      []int{1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "Negative numbers",
			arr:      []int{-1, -1, -2, -2, -2},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := UniqueOccurrences(tt.arr)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkUniqueOccurrences(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = UniqueOccurrences([]int{1, 2, 2, 1, 1, 3})
	}
}

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected bool
	}{
		{
			name:     "Identical strings",
			word1:    "abc",
			word2:    "abc",
			expected: true,
		},
		{
			name:     "Anagrams with same frequency",
			word1:    "aabbcc",
			word2:    "ccbbaa",
			expected: true,
		},
		{
			name:     "Different lengths",
			word1:    "abc",
			word2:    "abcd",
			expected: false,
		},
		{
			name:     "Same characters different frequency",
			word1:    "aabbcc",
			word2:    "aabbccc",
			expected: false,
		},
		{
			name:     "Different characters",
			word1:    "abc",
			word2:    "def",
			expected: false,
		},
		{
			name:     "Empty strings",
			word1:    "",
			word2:    "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := CloseStrings(tt.word1, tt.word2)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCloseStrings(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CloseStrings("aabbcc", "ccbbaa")
	}
}

func TestEqualPairs(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name:     "Single element grid",
			grid:     [][]int{{1}},
			expected: 1,
		},
		{
			name:     "All rows and columns equal",
			grid:     [][]int{{1, 2}, {1, 2}},
			expected: 0,
		},
		{
			name:     "No equal rows and columns",
			grid:     [][]int{{1, 2}, {3, 4}},
			expected: 0,
		},
		{
			name:     "Some equal rows and columns",
			grid:     [][]int{{1, 2, 3}, {2, 5, 6}, {3, 2, 3}},
			expected: 1,
		},
		{
			name:     "Larger grid with multiple matches",
			grid:     [][]int{{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := EqualPairs(tt.grid)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkEqualPairs(b *testing.B) {
	b.ReportAllocs()

	grid := [][]int{{1, 2, 3, 4}, {2, 3, 4, 1}, {3, 4, 1, 2}, {4, 1, 2, 3}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = EqualPairs(grid)
	}
}

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No stars",
			input:    "abc",
			expected: "abc",
		},
		{
			name:     "Single star",
			input:    "ab*c",
			expected: "ac",
		},
		{
			name:     "Multiple stars",
			input:    "leet**cod*e",
			expected: "lecoe",
		},
		{
			name:     "Stars at the end",
			input:    "erase*****",
			expected: "",
		},
		{
			name:     "Interleaved stars",
			input:    "a*b*c*d*",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := RemoveStars(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkRemoveStars(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = RemoveStars("leet**cod*e")
	}
}

func TestAsteroidsCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		expected  []int
	}{
		{
			name:      "Equal mass collision",
			asteroids: []int{5, 10, -10},
			expected:  []int{5},
		},
		{
			name:      "All positive",
			asteroids: []int{1, 2, 3, 4, 5},
			expected:  []int{1, 2, 3, 4, 5},
		},
		{
			name:      "All negative",
			asteroids: []int{-1, -2, -3, -4, -5},
			expected:  []int{-1, -2, -3, -4, -5},
		},
		{
			name:      "Mixed collision",
			asteroids: []int{10, 2, -5},
			expected:  []int{10},
		},
		{
			name:      "No collision",
			asteroids: []int{10, -2, -5},
			expected:  []int{10},
		},
		{
			name:      "Multiple collisions",
			asteroids: []int{8, -8, 10, -10, 5, -5},
			expected:  []int{},
		},
		{
			name:      "Large numbers",
			asteroids: []int{1000, -1000, 500, -500},
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := AsteroidCollision(tt.asteroids)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkAsteroidsCollision(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = AsteroidCollision([]int{10, -2, -5})
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single set of brackets",
			input:    "3[a]",
			expected: "aaa",
		},
		{
			name:     "Nested brackets",
			input:    "3[a2[b]]",
			expected: "abbabbabb",
		},
		{
			name:     "Multiple sets of brackets",
			input:    "2[abc]3[cd]ef",
			expected: "abcabccdcdcdef",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "No brackets",
			input:    "abc",
			expected: "abc",
		},
		{
			name:  "Complex nested brackets",
			input: "2[2[y]pq4[2[jk]e1[f]]]ef",
			expected: "yypqjkjkefjkjkefjkjkefjkjkefyy" +
				"pqjkjkefjkjkefjkjkefjkjkefef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := DecodeString(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkDecodeString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = DecodeString("2[2[y]pq4[2[jk]e1[f]]]ef")
	}
}

func TestPredictPartyVictory(t *testing.T) {
	tests := []struct {
		name     string
		senate   string
		expected string
	}{
		{
			name:     "Radiant wins",
			senate:   "RRDDD",
			expected: "Radiant",
		},
		{
			name:     "Dire wins",
			senate:   "RDDDR",
			expected: "Dire",
		},
		{
			name:     "Radiant wins with alternating",
			senate:   "RDRDR",
			expected: "Radiant",
		},
		{
			name:     "Dire wins with alternating",
			senate:   "DRDRD",
			expected: "Dire",
		},
		{
			name:     "Single Radiant",
			senate:   "R",
			expected: "Radiant",
		},
		{
			name:     "Single Dire",
			senate:   "D",
			expected: "Dire",
		},
		{
			name:     "Equal number of Radiant and Dire",
			senate:   "RDRD",
			expected: "Radiant",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := PredictPartyVictory(tt.senate)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkPredictPartyVictory(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = PredictPartyVictory("RRDDD")
	}
}

func TestMaxDepthOfEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	expected := 0

	actual := MaxDepth(root)
	assert.Equal(t, expected, actual)
}

func TestMaxDepthOfSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := 1

	actual := MaxDepth(root)
	assert.Equal(t, expected, actual)
}

func TestMaxDepthOfBalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expected := 3

	actual := MaxDepth(root)
	assert.Equal(t, expected, actual)
}

func BenchmarkMaxDepth(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	for i := 0; i < b.N; i++ {
		_ = MaxDepth(root)
	}
}

func TestMaxDepthOfUnbalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	expected := 4

	actual := MaxDepth(root)
	assert.Equal(t, expected, actual)
}

func TestLeafSimilarTreesWithSameLeaves(t *testing.T) {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	expected := true

	actual := LeafSimilar(root1, root2)
	assert.Equal(t, expected, actual)
}

func TestLeafSimilarTreesWithDifferentLeaves(t *testing.T) {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 10}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	expected := false

	actual := LeafSimilar(root1, root2)
	assert.Equal(t, expected, actual)
}

func TestLeafSimilarEmptyTrees(t *testing.T) {
	root1 := (*TreeNode)(nil)
	root2 := (*TreeNode)(nil)
	expected := true

	actual := LeafSimilar(root1, root2)
	assert.Equal(t, expected, actual)
}

func TestLeafSimilarOneEmptyTree(t *testing.T) {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := (*TreeNode)(nil)
	expected := false

	actual := LeafSimilar(root1, root2)
	assert.Equal(t, expected, actual)
}

func BenchmarkLeafSimilar(b *testing.B) {
	b.ReportAllocs()

	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = LeafSimilar(root1, root2)
	}
}

func TestGoodNodesInEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	expected := 0

	actual := GoodNodes(root)
	assert.Equal(t, expected, actual)
}

func TestGoodNodesInSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := 1

	actual := GoodNodes(root)
	assert.Equal(t, expected, actual)
}

func TestGoodNodesInBalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5},
		},
	}
	expected := 4

	actual := GoodNodes(root)
	assert.Equal(t, expected, actual)
}

func TestGoodNodesInUnbalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 2},
			},
		},
	}
	expected := 3

	actual := GoodNodes(root)
	assert.Equal(t, expected, actual)
}

func TestGoodNodesWithNegativeValues(t *testing.T) {
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:  -2,
			Left: &TreeNode{Val: -3},
		},
		Right: &TreeNode{
			Val:   -4,
			Right: &TreeNode{Val: -5},
		},
	}
	expected := 1

	actual := GoodNodes(root)
	assert.Equal(t, expected, actual)
}

func BenchmarkGoodNodes(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 5},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = GoodNodes(root)
	}
}

func TestPathSumInEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	expected := 0

	actual := PathSum(root, 10)
	assert.Equal(t, expected, actual)
}

func TestPathSumSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 10}
	expected := 1

	actual := PathSum(root, 10)
	assert.Equal(t, expected, actual)
}

func TestPathSumNoMatchingPath(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
	expected := 0

	actual := PathSum(root, 20)
	assert.Equal(t, expected, actual)
}

func TestPathSumMultipleMatchingPaths(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}
	expected := 2

	actual := PathSum(root, 22)
	assert.Equal(t, expected, actual)
}

func TestPathSumNegativeValues(t *testing.T) {
	root := &TreeNode{
		Val: -2,
		Left: &TreeNode{
			Val: -3,
		},
		Right: &TreeNode{
			Val: -1,
		},
	}
	expected := 1

	actual := PathSum(root, -5)
	assert.Equal(t, expected, actual)
}

func BenchmarkPathSum(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
		},
		Right: &TreeNode{
			Val:  8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 1},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = PathSum(root, 22)
	}
}

func TestLowestCommonAncestorInBalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	p := root.Left
	q := root.Right
	expected := root

	actual := LowestCommonAncestor(root, p, q)
	assert.Equal(t, expected, actual)
}

func TestLowestCommonAncestorInUnbalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2},
		},
	}
	p := root.Left
	q := root.Left.Right
	expected := root.Left

	actual := LowestCommonAncestor(root, p, q)
	assert.Equal(t, expected, actual)
}

func TestLowestCommonAncestorWithOneNode(t *testing.T) {
	root := &TreeNode{Val: 1}
	p := root
	q := root
	expected := root

	actual := LowestCommonAncestor(root, p, q)
	assert.Equal(t, expected, actual)
}

func TestLowestCommonAncestorWithNilNodes(t *testing.T) {
	root := &TreeNode{Val: 1}
	p := (*TreeNode)(nil)
	q := (*TreeNode)(nil)

	actual := LowestCommonAncestor(root, p, q)
	assert.Nil(t, actual)
}

func TestLowestCommonAncestorWithNilRoot(t *testing.T) {
	root := (*TreeNode)(nil)
	p := &TreeNode{Val: 1}
	q := &TreeNode{Val: 2}
	expected := (*TreeNode)(nil)

	actual := LowestCommonAncestor(root, p, q)
	assert.Equal(t, expected, actual)
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	p := root.Left
	q := root.Right

	for i := 0; i < b.N; i++ {
		_ = LowestCommonAncestor(root, p, q)
	}
}

func TestLongestZigZagInEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	expected := 0

	actual := LongestZigZag(root)
	assert.Equal(t, expected, actual)
}

func TestLongestZigZagInSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := 0

	actual := LongestZigZag(root)
	assert.Equal(t, expected, actual)
}

func TestLongestZigZagInBalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	expected := 2

	actual := LongestZigZag(root)
	assert.Equal(t, expected, actual)
}

func TestLongestZigZagInUnbalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}
	expected := 1

	actual := LongestZigZag(root)
	assert.Equal(t, expected, actual)
}

func TestLongestZigZagWithNegativeValues(t *testing.T) {
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -2,
			Left:  &TreeNode{Val: -3},
			Right: &TreeNode{Val: -4},
		},
		Right: &TreeNode{
			Val:   -5,
			Left:  &TreeNode{Val: -6},
			Right: &TreeNode{Val: -7},
		},
	}
	expected := 2

	actual := LongestZigZag(root)
	assert.Equal(t, expected, actual)
}

func BenchmarkLongestZigZag(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = LongestZigZag(root)
	}
}

func TestSearchBSTInEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	val := 5
	expected := (*TreeNode)(nil)

	actual := SearchBST(root, val)
	assert.Equal(t, expected, actual)
}

func TestSearchBSTSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 5}
	val := 5
	expected := root

	actual := SearchBST(root, val)
	assert.Equal(t, expected, actual)
}

func TestSearchBSTValueExists(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	val := 3
	expected := root.Left.Right

	actual := SearchBST(root, val)
	assert.Equal(t, expected, actual)
}

func TestSearchBSTValueDoesNotExist(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	val := 5
	expected := (*TreeNode)(nil)

	actual := SearchBST(root, val)
	assert.Equal(t, expected, actual)
}

func TestSearchBSTValueInRightSubtree(t *testing.T) {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	val := 7
	expected := root.Right

	actual := SearchBST(root, val)
	assert.Equal(t, expected, actual)
}

func BenchmarkSearchBST(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}

	for i := 0; i < b.N; i++ {
		_ = SearchBST(root, 3)
	}
}

func TestRightSideViewOfEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	expected := []int(nil)

	actual := RightSideView(root)
	assert.Equal(t, expected, actual)
}

func TestRightSideViewOfSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := []int{1}

	actual := RightSideView(root)
	assert.Equal(t, expected, actual)
}

func TestRightSideViewOfBalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}
	expected := []int{1, 3, 6}

	actual := RightSideView(root)
	assert.Equal(t, expected, actual)
}

func TestRightSideViewOfUnbalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 4},
			},
		},
	}
	expected := []int{1, 2, 3, 4}

	actual := RightSideView(root)
	assert.Equal(t, expected, actual)
}

func TestRightSideViewWithNilNodes(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}
	expected := []int{1, 3, 4}

	actual := RightSideView(root)
	assert.Equal(t, expected, actual)
}

func BenchmarkRightSideView(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = RightSideView(root)
	}
}

func TestMaxLevelSumInEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	expected := 0

	actual := MaxLevelSum(root)
	assert.Equal(t, expected, actual)
}

func TestMaxLevelSumInSingleNodeTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	expected := 1

	actual := MaxLevelSum(root)
	assert.Equal(t, expected, actual)
}

func TestMaxLevelSumInBalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}
	expected := 3

	actual := MaxLevelSum(root)
	assert.Equal(t, expected, actual)
}

func TestMaxLevelSumInUnbalancedTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 8},
			},
		},
		Right: &TreeNode{Val: 3},
	}
	expected := 4

	actual := MaxLevelSum(root)
	assert.Equal(t, expected, actual)
}

func TestMaxLevelSumWithNegativeValues(t *testing.T) {
	root := &TreeNode{
		Val: -1,
		Left: &TreeNode{
			Val:   -2,
			Left:  &TreeNode{Val: -4},
			Right: &TreeNode{Val: -5},
		},
		Right: &TreeNode{
			Val:   -3,
			Left:  &TreeNode{Val: -6},
			Right: &TreeNode{Val: -7},
		},
	}
	expected := 1

	actual := MaxLevelSum(root)
	assert.Equal(t, expected, actual)
}

func BenchmarkMaxLevelSum(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = MaxLevelSum(root)
	}
}

func TestDeleteNodeInEmptyTree(t *testing.T) {
	root := (*TreeNode)(nil)
	key := 5
	expected := (*TreeNode)(nil)

	actual := DeleteNode(root, key)
	assert.Equal(t, expected, actual)
}

func TestDeleteNodeLeafNode(t *testing.T) {
	root := &TreeNode{Val: 5}
	key := 5
	expected := (*TreeNode)(nil)

	actual := DeleteNode(root, key)
	assert.Equal(t, expected, actual)
}

func TestDeleteNodeSingleChild(t *testing.T) {
	root := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 3},
	}
	key := 5
	expected := root.Left

	actual := DeleteNode(root, key)
	assert.Equal(t, expected, actual)
}

func TestDeleteNodeTwoChildren(t *testing.T) {
	root := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{
			Val:  7,
			Left: &TreeNode{Val: 6},
		},
	}
	key := 5
	expected := &TreeNode{
		Val:   6,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 7},
	}

	actual := DeleteNode(root, key)
	assert.Equal(t, expected, actual)
}

func TestDeleteNodeNonExistentKey(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 7},
	}
	key := 10
	expected := root

	actual := DeleteNode(root, key)
	assert.Equal(t, expected, actual)
}

func BenchmarkDeleteNode(b *testing.B) {
	b.ReportAllocs()

	root := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 3},
		Right: &TreeNode{
			Val:  7,
			Left: &TreeNode{Val: 6},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 9},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		_ = DeleteNode(root, 5)
	}
}

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name     string
		rooms    [][]int
		expected bool
	}{
		{
			name:     "EmptyRooms",
			rooms:    [][]int{},
			expected: true,
		},
		{
			name:     "SingleRoom",
			rooms:    [][]int{{}},
			expected: true,
		},
		{
			name:     "AllKeys",
			rooms:    [][]int{{1}, {2}, {3}, {}},
			expected: true,
		},
		{
			name:     "MissingKey",
			rooms:    [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			expected: false,
		},
		{
			name:     "CircularKeys",
			rooms:    [][]int{{1}, {2}, {0}},
			expected: true,
		},
		{
			name:     "Disconnected",
			rooms:    [][]int{{1}, {2}, {}, {3}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := CanVisitAllRooms(tt.rooms)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkCanVisitAllRooms(b *testing.B) {
	b.ReportAllocs()

	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = CanVisitAllRooms(rooms)
	}
}

func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		name        string
		isConnected [][]int
		expected    int
	}{
		{
			name:        "NoProvincesInEmptyGraph",
			isConnected: [][]int{},
			expected:    0,
		},
		{
			name:        "SingleProvinceInSingleCity",
			isConnected: [][]int{{1}},
			expected:    1,
		},
		{
			name: "MultipleProvincesInDisconnectedGraph",
			isConnected: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			expected: 3,
		},
		{
			name: "SingleProvinceInFullyConnectedGraph",
			isConnected: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 1,
		},
		{
			name: "MultipleProvincesInPartiallyConnectedGraph",
			isConnected: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := FindCircleNum(tt.isConnected)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFindCircleNum(b *testing.B) {
	b.ReportAllocs()

	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FindCircleNum(isConnected)
	}
}

func TestMinReorder(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		expected    int
	}{
		{
			name:        "AllEdgesReordered",
			n:           3,
			connections: [][]int{{0, 1}, {1, 2}},
			expected:    2,
		},
		{
			name:        "NoEdgesReordered",
			n:           3,
			connections: [][]int{{1, 0}, {2, 1}},
			expected:    0,
		},
		{
			name:        "ReorderMixedEdges",
			n:           4,
			connections: [][]int{{0, 1}, {2, 0}, {3, 2}},
			expected:    1,
		},
		{
			name:        "ReorderSingleNode",
			n:           1,
			connections: [][]int{},
			expected:    0,
		},
		{
			name:        "ReorderDisconnectedGraph",
			n:           4,
			connections: [][]int{{0, 1}, {2, 3}},
			expected:    1,
		},
		{
			name: "ReorderThreeNodes",
			n:    6,
			connections: [][]int{
				{0, 1},
				{1, 3},
				{2, 3},
				{4, 0},
				{4, 5},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MinReorder(tt.n, tt.connections)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinReorder(b *testing.B) {
	b.ReportAllocs()

	connections := [][]int{
		{0, 1},
		{1, 3},
		{2, 3},
		{4, 0},
		{4, 5},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = MinReorder(6, connections)
	}
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "All fresh oranges",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: -1,
		},
		{
			name: "All rotten oranges",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "Mixed fresh and rotten oranges",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "No oranges",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "Single fresh orange",
			grid: [][]int{
				{1},
			},
			expected: -1,
		},
		{
			name: "Single rotten orange",
			grid: [][]int{
				{2},
			},
			expected: 0,
		},
		{
			name: "Fresh oranges surrounded by empty cells",
			grid: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 1, 0},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := OrangesRotting(tt.grid)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkOrangesRotting(b *testing.B) {
	b.ReportAllocs()

	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = OrangesRotting(grid)
	}
}

func TestNearestExit(t *testing.T) {
	tests := []struct {
		name     string
		maze     [][]byte
		entrance []int
		expected int
	}{
		{
			name: "SinglePathToExit",
			maze: [][]byte{
				{'+', '+', '.', '+'},
				{'.', '.', '.', '+'},
				{'+', '+', '+', '.'},
			},
			entrance: []int{1, 0},
			expected: 3,
		},
		{
			name: "NoExit",
			maze: [][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			entrance: []int{1, 0},
			expected: 2,
		},
		{
			name: "MultiplePathsToExit",
			maze: [][]byte{
				{'+', '.', '+', '+', '+', '+', '+'},
				{'+', '.', '.', '.', '.', '.', '+'},
				{'+', '+', '+', '+', '.', '+', '+'},
			},
			entrance: []int{1, 1},
			expected: 1,
		},
		{
			name: "ExitOnEdge",
			maze: [][]byte{
				{'+', '+', '+'},
				{'.', '.', '.'},
				{'+', '+', '+'},
			},
			entrance: []int{1, 2},
			expected: 2,
		},
		{
			name: "LargeMaze",
			maze: [][]byte{
				{'+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
			},
			entrance: []int{1, 0},
			expected: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := NearestExit(tt.maze, tt.entrance)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkNearestExit(b *testing.B) {
	b.ReportAllocs()

	maze := [][]byte{
		{'+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
	}
	entrance := []int{1, 1}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = NearestExit(maze, entrance)
	}
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "kth largest in sorted array",
			nums:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: 4,
		},
		{
			name:     "kth largest in reverse sorted array",
			nums:     []int{5, 4, 3, 2, 1},
			k:        3,
			expected: 3,
		},
		{
			name:     "kth largest in array with duplicates",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			name:     "kth largest in array with negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        1,
			expected: -1,
		},
		{
			name:     "kth largest in array with mixed numbers",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "kth largest in single element array",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := FindKthLargest(tt.nums, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	b.ReportAllocs()

	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FindKthLargest(nums, k)
	}
}

func TestSmallestInfiniteSetConstructorCreatesSetWithAllNumbers(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	for i := 1; i <= 1000; i++ {
		assert.Equal(t, 1, s.arr[i])
	}
}

func TestPopSmallestReturnsSmallestNumber(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	assert.Equal(t, 1, s.PopSmallest())
	assert.Equal(t, 0, s.arr[1])
}

func TestPopSmallestRemovesNumberFromSet(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	s.PopSmallest()
	assert.Equal(t, 0, s.arr[1])
}

func TestAddBackReinsertsNumberIntoSet(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	s.PopSmallest()
	s.AddBack(1)
	assert.Equal(t, 1, s.arr[1])
}

func TestAddBackDoesNotReinsertExistingNumber(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	s.AddBack(1)
	assert.Equal(t, 1, s.arr[1])
}

func TestPopSmallestAfterAddBack(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	s.PopSmallest()
	s.AddBack(1)
	assert.Equal(t, 1, s.PopSmallest())
}

func TestPopSmallestMultipleTimes(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	assert.Equal(t, 1, s.PopSmallest())
	assert.Equal(t, 2, s.PopSmallest())
	assert.Equal(t, 3, s.PopSmallest())
}

func TestAddBackAndPopSmallest(t *testing.T) {
	t.Parallel()

	s := SmallestInfiniteSetConstructor()
	s.PopSmallest()
	s.PopSmallest()
	s.AddBack(1)
	assert.Equal(t, 1, s.PopSmallest())
	assert.Equal(t, 3, s.PopSmallest())
}

func TestMaxScore(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		k        int
		expected int64
	}{
		{
			name:     "equal length arrays",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{5, 4, 3, 2, 1},
			k:        3,
			expected: int64(18),
		},
		{
			name:     "single element arrays",
			nums1:    []int{1},
			nums2:    []int{1},
			k:        1,
			expected: int64(1),
		},
		{
			name:     "k equal to zero",
			nums1:    []int{1, 2, 3},
			nums2:    []int{3, 2, 1},
			k:        0,
			expected: int64(0),
		},
		{
			name:     "negative values",
			nums1:    []int{-1, -2, -3},
			nums2:    []int{-3, -2, -1},
			k:        2,
			expected: int64(10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MaxScore(tt.nums1, tt.nums2, tt.k)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMaxScore(b *testing.B) {
	b.ReportAllocs()

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 4, 3, 2, 1}
	k := 3
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = MaxScore(nums1, nums2, k)
	}
}

func TestTotalCost(t *testing.T) {
	tests := []struct {
		name       string
		costs      []int
		k          int
		candidates int
		expected   int64
	}{
		{
			name:       "Equal costs",
			costs:      []int{5, 5, 5, 5, 5},
			k:          3,
			candidates: 2,
			expected:   15,
		},
		{
			name:       "Different costs",
			costs:      []int{1, 2, 3, 4, 5},
			k:          3,
			candidates: 2,
			expected:   6,
		},
		{
			name:       "Single candidate",
			costs:      []int{1, 2, 3, 4, 5},
			k:          1,
			candidates: 1,
			expected:   1,
		},
		{
			name:       "More candidates than costs",
			costs:      []int{1, 2, 3},
			k:          2,
			candidates: 5,
			expected:   3,
		},
		{
			name:       "Zero candidates",
			costs:      []int{1, 2, 3, 4, 5},
			k:          3,
			candidates: 0,
			expected:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := TotalCost(tt.costs, tt.k, tt.candidates)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkTotalCost(b *testing.B) {
	b.ReportAllocs()

	costs := []int{1, 2, 3, 4, 5}
	k := 3
	candidates := 2
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = TotalCost(costs, k, candidates)
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
			name:     "Single pile, exact hours",
			piles:    []int{5},
			h:        5,
			expected: 1,
		},
		{
			name:     "Single pile, more hours",
			piles:    []int{5},
			h:        10,
			expected: 1,
		},
		{
			name:     "Single pile, less hours",
			piles:    []int{5},
			h:        1,
			expected: 5,
		},
		{
			name:     "Multiple piles, exact hours",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Empty piles",
			piles:    []int{},
			h:        5,
			expected: 1,
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

func TestSuccessfulPairs(t *testing.T) {
	tests := []struct {
		name     string
		spells   []int
		potions  []int
		success  int64
		expected []int
	}{
		{
			name:     "All pairs successful",
			spells:   []int{10, 20},
			potions:  []int{2, 3, 4},
			success:  20,
			expected: []int{3, 3},
		},
		{
			name:     "No pairs successful",
			spells:   []int{1, 2},
			potions:  []int{1, 1, 1},
			success:  10,
			expected: []int{0, 0},
		},
		{
			name:     "Mixed success",
			spells:   []int{5, 1},
			potions:  []int{1, 2, 3},
			success:  5,
			expected: []int{3, 0},
		},
		{
			name:     "Empty spells",
			spells:   []int{},
			potions:  []int{1, 2, 3},
			success:  5,
			expected: []int{},
		},
		{
			name:     "Empty potions",
			spells:   []int{1, 2, 3},
			potions:  []int{},
			success:  5,
			expected: []int{0, 0, 0},
		},
		{
			name:     "Large numbers",
			spells:   []int{100000, 200000},
			potions:  []int{100000, 200000},
			success:  10000000000,
			expected: []int{2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := SuccessfulPairs(tt.spells, tt.potions, tt.success)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkSuccessfulPairs(b *testing.B) {
	b.ReportAllocs()

	spells := []int{10, 20}
	potions := []int{2, 3, 4}
	success := int64(20)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = SuccessfulPairs(spells, potions, success)
	}
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected any // int or []int
	}{
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "Multiple peaks",
			nums:     []int{1, 3, 2, 4, 1},
			expected: []int{1, 3}, // Both 1 and 3 are valid peaks
		},
		{
			name:     "Descending order",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Ascending order",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Plateau",
			nums:     []int{1, 2, 2, 3, 3, 2, 1},
			expected: []int{3, 4}, // Both 3 and 4 are valid peaks
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FindPeakElement(tt.nums)
			switch expected := tt.expected.(type) {
			case int:
				assert.Equal(t, expected, actual)
			case []int:
				assert.Contains(t, expected, actual)
			}
		})
	}
}

func BenchmarkFindPeakElement(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = FindPeakElement([]int{1, 3, 2, 4, 1})
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

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		n        int
		expected [][]int
	}{
		{
			name:     "SingleCombination",
			k:        3,
			n:        7,
			expected: [][]int{{1, 2, 4}},
		},
		{
			name:     "MultipleCombinations",
			k:        3,
			n:        9,
			expected: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			name:     "NoCombinations",
			k:        4,
			n:        1,
			expected: [][]int{},
		},
		{
			name:     "SingleElementCombination",
			k:        1,
			n:        9,
			expected: [][]int{{9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := CombinationSum3(tt.k, tt.n)
			assert.ElementsMatch(t, tt.expected, actual)
		})
	}
}

func BenchmarkCombinationSum3(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = CombinationSum3(3, 9)
	}
}

func TestTribonacciDP(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Zero",
			n:        0,
			expected: 0,
		},
		{
			name:     "One",
			n:        1,
			expected: 1,
		},
		{
			name:     "Two",
			n:        2,
			expected: 1,
		},
		{
			name:     "Three",
			n:        3,
			expected: 2,
		},
		{
			name:     "Four",
			n:        4,
			expected: 4,
		},
		{
			name:     "Ten",
			n:        10,
			expected: 149,
		},
		{
			name:     "Thirty",
			n:        30,
			expected: 29249425,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := TribonacciDP(tt.n)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkTribonacciDP(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = TribonacciDP(30)
	}
}

func TestTribonacciMemo(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Zero",
			n:        0,
			expected: 0,
		},
		{
			name:     "One",
			n:        1,
			expected: 1,
		},
		{
			name:     "Two",
			n:        2,
			expected: 1,
		},
		{
			name:     "Three",
			n:        3,
			expected: 2,
		},
		{
			name:     "Four",
			n:        4,
			expected: 4,
		},
		{
			name:     "Ten",
			n:        10,
			expected: 149,
		},
		{
			name:     "Thirty",
			n:        30,
			expected: 29249425,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := TribonacciMemo(tt.n)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkTribonacciMemo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = TribonacciMemo(30)
	}
}

func TestMinCostClimbingStairsDP(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "TwoSteps",
			cost:     []int{10, 15},
			expected: 10,
		},
		{
			name:     "ThreeSteps",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "MultipleSteps",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
		{
			name:     "AllEqualSteps",
			cost:     []int{5, 5, 5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "IncreasingSteps",
			cost:     []int{1, 2, 3, 4, 5, 6},
			expected: 9,
		},
		{
			name:     "DecreasingSteps",
			cost:     []int{6, 5, 4, 3, 2, 1},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MinCostClimbingStairsDP(tt.cost)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinCostClimbingStairsDP(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinCostClimbingStairsDP([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	}
}

func TestMinCostClimbingStairsMemo(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "TwoSteps",
			cost:     []int{10, 15},
			expected: 10,
		},
		{
			name:     "ThreeSteps",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "MultipleSteps",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
		{
			name:     "AllEqualSteps",
			cost:     []int{5, 5, 5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "IncreasingSteps",
			cost:     []int{1, 2, 3, 4, 5, 6},
			expected: 9,
		},
		{
			name:     "DecreasingSteps",
			cost:     []int{6, 5, 4, 3, 2, 1},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := MinCostClimbingStairsMemo(tt.cost)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinCostClimbingStairsMemo(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = MinCostClimbingStairsMemo([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	}
}

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "SingleHouse",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "TwoHouses",
			nums:     []int{5, 10},
			expected: 10,
		},
		{
			name:     "MultipleHouses",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "NonAdjacentHouses",
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			name:     "AllHousesWithSameValue",
			nums:     []int{4, 4, 4, 4, 4},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := Rob(tt.nums)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkRob(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Rob([]int{1, 2, 3, 1, 11, 8, 4, 3})
	}
}
