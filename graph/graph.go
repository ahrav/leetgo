package graph

import (
	"container/heap"
	"sort"
)

type Edge struct{ to, weight int }

type EdgeMinHeap []Edge

func (h EdgeMinHeap) Len() int           { return len(h) }
func (h EdgeMinHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h EdgeMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeMinHeap) Push(x any) { *h = append(*h, x.(Edge)) }

func (h *EdgeMinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

// MinCostConnectPoints - https://leetcode.com/problems/min-cost-to-connect-all-points/
func MinCostConnectPoints(points [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	distance := func(p1, p2 []int) int {
		return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
	}

	n := len(points)
	if n <= 1 {
		return 0
	}

	h := new(EdgeMinHeap)
	heap.Init(h)

	visited := make([]bool, n)

	heap.Push(h, Edge{to: 0, weight: 0})

	totalCost := 0

	for h.Len() > 0 {
		edge := heap.Pop(h).(Edge)

		if visited[edge.to] {
			continue
		}

		visited[edge.to] = true
		totalCost += edge.weight

		for i := 0; i < n; i++ {
			if !visited[i] {
				dist := distance(points[edge.to], points[i])
				heap.Push(h, Edge{to: i, weight: dist})
			}
		}
	}

	return totalCost
}

type UnionFind struct {
	parents []int
	ranks   []int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	ranks := make([]int, n)
	for i := range parents {
		parents[i] = i
	}

	return &UnionFind{parents: parents, ranks: ranks}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parents[x] != x {
		uf.parents[x] = uf.Find(uf.parents[x])
	}
	return uf.parents[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX, rootY := uf.Find(x), uf.Find(y)

	if rootX != rootY {
		if uf.ranks[rootX] > uf.ranks[rootY] {
			uf.parents[rootY] = rootX
		} else if uf.ranks[rootX] < uf.ranks[rootY] {
			uf.parents[rootX] = rootY
		} else {
			uf.parents[rootY] = rootX
			uf.ranks[rootX]++
		}
		return true
	}

	return false
}

// MinimumCost - https://leetcode.com/problems/connecting-cities-with-minimum-cost/
func MinimumCost(n int, connections [][]int) int {
	if len(connections) < n-1 { // Not enough cities to connect
		return -1
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	uf := NewUnionFind(n)

	totalCost, edgesUsed := 0, 0
	for _, conn := range connections {
		x, y, cost := conn[0]-1, conn[1]-1, conn[2]
		if uf.Union(x, y) {
			totalCost += cost
			edgesUsed++
			if edgesUsed == n-1 {
				return totalCost
			}
		}
	}

	return -1
}
