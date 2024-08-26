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

	h := new(EdgeMinHeap)
	heap.Init(h)

	heap.Push(h, Edge{to: 0, weight: 0})

	n := len(points)
	if n == 0 {
		return 0
	}
	visited := make([]bool, n)

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
	if len(connections) < n-1 {
		return -1
	}

	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	uf := NewUnionFind(n)
	totalCost, edgesAccessed := 0, 0
	for _, conn := range connections {
		x, y, cost := conn[0]-1, conn[1]-1, conn[2]
		if uf.Union(x, y) {
			totalCost += cost
			edgesAccessed++
			if edgesAccessed == n-1 {
				return totalCost
			}
		}
	}

	return -1
}

// LadderLength - https://leetcode.com/problems/word-ladder/?envType=problem-list-v2&envId=954v5ops
func LadderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	wordSet := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	type WordDepth struct {
		word  string
		depth int
	}

	queue := []WordDepth{{word: beginWord, depth: 1}}
	for len(queue) > 0 {
		wordD := queue[0]
		queue = queue[1:]

		word, depth := wordD.word, wordD.depth
		for i := 0; i < len(word); i++ {
			for char := 'a'; char <= 'z'; char++ {
				newWord := word[:i] + string(char) + word[i+1:]

				if newWord == endWord {
					return depth + 1
				}

				if _, ok := wordSet[newWord]; ok {
					queue = append(queue, WordDepth{word: newWord, depth: depth + 1})
					delete(wordSet, newWord)
				}
			}
		}
	}

	return 0
}

// LadderLength - https://leetcode.com/problems/word-ladder/?envType=problem-list-v2&envId=954v5ops
func LadderLengthBirectional(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	wordSet := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	fwdQueue := map[string]int{beginWord: 1}
	bckQueue := map[string]int{endWord: 1}

	for len(fwdQueue) > 0 && len(bckQueue) > 0 {
		if len(fwdQueue) > len(bckQueue) {
			fwdQueue, bckQueue = bckQueue, fwdQueue
		}

		nextQueue := make(map[string]int)
		for word, depth := range fwdQueue {
			for i := 0; i < len(word); i++ {
				for char := 'a'; char <= 'z'; char++ {
					newWord := word[:i] + string(char) + word[i+1:]

					if bckDepth, ok := bckQueue[newWord]; ok {
						return bckDepth + depth
					}

					if _, ok := wordSet[newWord]; ok {
						nextQueue[newWord] = depth + 1
						delete(wordSet, newWord)
					}
				}
			}
		}
		fwdQueue = nextQueue
	}

	return 0
}

// FindLadders - https://leetcode.com/problems/word-ladder-ii/?envType=problem-list-v2&envId=954v5ops
func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return [][]string{}
	}

	queue := [][]string{{beginWord}}
	visited := map[string]bool{beginWord: true}
	found := false
	allLevels := make(map[string][]string)

	for len(queue) > 0 && !found {
		levelVisited := make(map[string]bool)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			path := queue[0]
			queue = queue[1:]
			lastWord := path[len(path)-1]

			for i := 0; i < len(lastWord); i++ {
				for c := 'a'; c <= 'z'; c++ {
					nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
					if wordSet[nextWord] && !visited[nextWord] {
						if nextWord == endWord {
							found = true
						}

						if !levelVisited[nextWord] {
							levelVisited[nextWord] = true
							queue = append(queue, append([]string{}, append(path, nextWord)...))
						}

						allLevels[nextWord] = append(allLevels[nextWord], lastWord)
					}
				}
			}
		}

		for word := range levelVisited {
			visited[word] = true
		}
	}

	if !found {
		return [][]string{}
	}

	reverse := func(arr []string) []string {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}

	var result [][]string

	var backtrack func(word string, path []string)
	backtrack = func(word string, path []string) {
		if word == beginWord {
			tmp := make([]string, len(path))
			copy(tmp, path)
			result = append(result, reverse(tmp))
			return
		}

		for _, prevWord := range allLevels[word] {
			backtrack(prevWord, append(path, prevWord))
		}
	}

	backtrack(endWord, []string{endWord})

	return result
}
