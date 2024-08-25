package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{"NoPoints", [][]int{}, 0},
		{"SinglePoint", [][]int{{0, 0}}, 0},
		{"TwoPoints", [][]int{{0, 0}, {1, 1}}, 2},
		{"ThreePoints", [][]int{{0, 0}, {2, 2}, {3, 10}}, 13},
		{"MultiplePoints", [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := MinCostConnectPoints(tt.points)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func BenchmarkMinCostConnectPoints(b *testing.B) {
	tests := []struct {
		name   string
		points [][]int
	}{
		{"NoPoints", [][]int{}},
		{"SinglePoint", [][]int{{0, 0}}},
		{"TwoPoints", [][]int{{0, 0}, {1, 1}}},
		{"ThreePoints", [][]int{{0, 0}, {2, 2}, {3, 10}}},
		{"MultiplePoints", [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = MinCostConnectPoints(tt.points)
			}
		})
	}
}
