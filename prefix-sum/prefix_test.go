package prefixsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVowelStrings(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		queries  [][]int
		expected []int
	}{
		{"ValidInput", []string{"apple", "orange", "banana", "umbrella"}, [][]int{{0, 1}, {1, 3}, {0, 3}}, []int{2, 2, 3}},
		{"SingleWord", []string{"apple"}, [][]int{{0, 0}}, []int{1}},
		{"NoVowelWords", []string{"bcd", "fgh", "jkl"}, [][]int{{0, 2}}, []int{0}},
		{"MixedVowelWords", []string{"apple", "banana", "ice", "orange"}, [][]int{{0, 3}, {1, 2}, {2, 3}}, []int{3, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := VowelStrings(tt.words, tt.queries)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkVowelStrings(b *testing.B) {
	words := []string{"apple", "orange", "banana", "umbrella"}
	queries := [][]int{{0, 1}, {1, 3}, {0, 3}}
	for i := 0; i < b.N; i++ {
		VowelStrings(words, queries)
	}
}
