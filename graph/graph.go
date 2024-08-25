package graph

import "container/heap"

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
