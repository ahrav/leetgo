package seventyfive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
