package top50

import (
	"testing"

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
