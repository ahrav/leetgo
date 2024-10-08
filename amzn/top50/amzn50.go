package top50

import (
	"container/heap"
	"sort"
	"strings"
)

// NumberToWords - https://leetcode.com/problems/integer-to-english-words/
func NumberToWords(num int) string {
	// Handle zero edge case.
	if num == 0 {
		return "Zero"
	}

	// Setup arrays in order to quickly determine word for a place value. (ones, hundreds, etc..)
	var (
		ones      = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
		tens      = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		thousands = []string{"", "Thousand", "Million", "Billion"}
	)

	var helper func(num int) string
	helper = func(num int) string {
		switch {
		case num == 0:
			return ""
		case num < 10:
			return ones[num] + " "
		case num < 20:
			return teens[num-10] + " "
		case num < 100:
			return tens[num/10] + " " + helper(num%10)
		default:
			return ones[num/100] + " Hundred " + helper(num%100)
		}
	}

	var result []string
	for i := 0; num > 0; i++ {
		rem := num % 1000
		if rem != 0 {
			result = append([]string{helper(rem) + thousands[i]}, result...)
		}
		num /= 1000
	}

	return strings.TrimSpace(strings.Join(result, " "))
}

// Trap - https://leetcode.com/problems/trapping-rain-water/
func Trap(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	leftMax, rightMax := arr[0], arr[n-1]
	left, right := 0, n-1

	water := 0
	for left < right {
		if leftMax < rightMax {
			left++

			if arr[left] < leftMax {
				water += leftMax - arr[left]
			} else {
				leftMax = arr[left]
			}
		} else {
			right--

			if arr[right] < rightMax {
				water += rightMax - arr[right]
			} else {
				rightMax = arr[right]
			}
		}
	}

	return water
}

// NumIslands - https://leetcode.com/problems/number-of-islands/
func NumIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == '0' {
			return
		}

		grid[x][y] = '0'
		for _, dir := range directions {
			nr, nc := x+dir[0], y+dir[1]
			dfs(nr, nc)
		}
	}

	numIslands := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == '1' {
				dfs(i, j)
				numIslands++
			}
		}
	}

	return numIslands
}

// GroupAnagrams - https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	freq := make(map[[26]int][]string)

	for _, word := range strs {
		var cnt [26]int
		for _, char := range word {
			cnt[char-'a']++
		}

		freq[cnt] = append(freq[cnt], word)
	}

	result := make([][]string, 0, len(freq))
	for _, words := range freq {
		result = append(result, words)
	}

	return result
}

// MinimumSwaps - https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/
func MinimumSwaps(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	minIdx, maxIdx := 0, 0
	for i, num := range nums {
		if num < nums[minIdx] {
			minIdx = i
		}
		if num >= nums[maxIdx] {
			maxIdx = i
		}
	}

	totalSwaps := n - 1 - maxIdx + minIdx
	if minIdx < maxIdx {
		return totalSwaps
	}

	return totalSwaps - 1
}

// MinMeetingRooms - https://leetcode.com/problems/meeting-rooms-ii/
func MinMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	if n == 1 {
		return 1
	}

	starts, ends := make([]int, n), make([]int, n)
	for i := range n {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	startPtr, endPtr := 0, 0
	rooms, maxRooms := 0, 1

	for startPtr < n {
		if starts[startPtr] < ends[endPtr] {
			rooms++
			startPtr++
		} else {
			rooms--
			endPtr++
		}

		if rooms > maxRooms {
			maxRooms = rooms
		}
	}

	return maxRooms
}

type CharFreq struct {
	freq int
	char byte
}

type MaxCharFreqHeap []CharFreq

func (h MaxCharFreqHeap) Len() int           { return len(h) }
func (h MaxCharFreqHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxCharFreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxCharFreqHeap) Push(x any) { *h = append(*h, x.(CharFreq)) }
func (h *MaxCharFreqHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ReorganizeString - https://leetcode.com/problems/reorganize-string/
func ReorganizeString(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	half := (n + 1) / 2
	var charFreq [26]int
	for _, char := range s {
		charFreq[char-'a']++
		if charFreq[char-'a'] > half {
			return ""
		}
	}

	h := new(MaxCharFreqHeap)
	heap.Init(h)

	for idx, freq := range charFreq {
		if freq > 0 {
			heap.Push(h, CharFreq{char: byte(idx) + 'a', freq: freq})
		}
	}

	var builder strings.Builder
	var prev CharFreq

	for h.Len() > 0 {
		cf := heap.Pop(h).(CharFreq)
		builder.WriteByte(cf.char)

		if prev.freq > 0 {
			heap.Push(h, prev)
		}

		cf.freq--
		prev = cf
	}

	return builder.String()
}

// OrangesRotting - https://leetcode.com/problems/rotting-oranges/
func OrangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var queue [][2]int
	freshCnt := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 {
				freshCnt++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	if freshCnt == 0 {
		return 0
	}

	if len(queue) == 0 {
		return -1
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	minutes := -1
	for len(queue) > 0 {
		level := len(queue)

		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				nr, nc := curr[0]+dir[0], curr[1]+dir[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == 0 {
					continue
				}
				if grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					freshCnt--
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		minutes++
	}

	if freshCnt == 0 {
		return minutes
	}
	return -1
}

// MaxSlidingWindow - https://leetcode.com/problems/sliding-window-maximum/
func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	if n == 1 || k == 1 {
		return nums
	}

	var deque []int
	result := make([]int, 0, n-k+1)

	for i, num := range nums {
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

// CopyRandomList - https://leetcode.com/problems/copy-list-with-random-pointer/
func CopyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	curr := dummy

	// Add a copy of each node inbetwen existing nodes.
	for curr != nil {
		copyNode := &ListNode{Val: curr.Val}
		copyNode.Next = curr.Next
		curr.Next = copyNode
		curr = copyNode.Next
	}

	curr = head
	// Set random pointers for new copied nodes using the random pointers from the original nodes.
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// Remove the original nodes and return the copies.
	copyCurr, curr := dummy, head
	for curr != nil {
		copyCurr.Next = curr.Next
		curr.Next = curr.Next.Next
		copyCurr = copyCurr.Next
		curr = curr.Next
	}

	return dummy.Next
}

// LadderLength - https://leetcode.com/problems/word-ladder/
func LadderLength(beginWord, endWord string, wordList []string) int {
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
		// Always explore the smaller number of options.
		if len(fwdQueue) > len(bckQueue) {
			fwdQueue, bckQueue = bckQueue, fwdQueue
		}

		nextQueue := make(map[string]int)
		for word, depth := range fwdQueue {
			for i := 0; i < len(word); i++ {
				for c := 'a'; c <= 'z'; c++ {
					newW := word[:i] + string(c) + word[i+1:]

					if bckDepth, ok := bckQueue[newW]; ok {
						return bckDepth + depth
					}

					if _, ok := wordSet[newW]; ok {
						delete(wordSet, newW)
						nextQueue[newW] = depth + 1
					}
				}
			}
		}
		fwdQueue = nextQueue
	}

	return 0
}

// SmallestDistancePair - https://leetcode.com/problems/find-k-th-smallest-pair-distance/
func SmallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)

	countPairs := func(mid int) int {
		count, j := 0, 0
		for i := 0; i < n; i++ {
			for j < n && nums[j]-nums[i] <= mid {
				j++
			}
			count += j - i - 1
		}
		return count
	}

	left, right := 0, nums[n-1]-nums[0]
	for left < right {
		mid := left + (right-left)/2
		if countPairs(mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// PlatesBetweenCandles - https://leetcode.com/problems/plates-between-candles/
func PlatesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	if n == 0 {
		return nil
	}

	prefix := make([]int, n+1)
	prev, next := make([]int, n), make([]int, n)

	lastCandle := -1
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i]
		if s[i] == '*' {
			prefix[i+1]++
		} else {
			lastCandle = i
		}
		prev[i] = lastCandle
	}

	lastCandle = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			lastCandle = i
		}
		next[i] = lastCandle
	}

	result := make([]int, len(queries))
	for idx, query := range queries {
		left, right := query[0], query[1]
		start, end := next[left], prev[right]

		if start != -1 && end != -1 && start < end {
			result[idx] = prefix[end] - prefix[start]
		}
	}

	return result
}

// CanFinish - https://leetcode.com/problems/course-schedule/
func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		course, req := pair[0], pair[1]
		graph[req] = append(graph[req], course)
	}

	visited := make([]bool, numCourses)
	recurStack := make([]bool, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if recurStack[course] {
			return true
		}

		if visited[course] {
			return false
		}

		recurStack[course] = true
		visited[course] = true

		for _, req := range graph[course] {
			if hasCycle(req) {
				return true
			}
		}

		recurStack[course] = false
		return false
	}

	for i := range numCourses {
		if !visited[i] {
			if hasCycle(i) {
				return false
			}
		}
	}

	return true
}

// WordBreakII - https://leetcode.com/problems/word-break-ii/
func WordBreakII(s string, wordDict []string) []string {
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	memo := make(map[string][]string)

	var backtrack func(s string) []string
	backtrack = func(s string) []string {
		if res, ok := memo[s]; ok {
			return res
		}

		var result []string
		if _, ok := wordSet[s]; ok {
			result = append(result, s)
		}

		for i := 1; i < len(s); i++ {
			prefix := s[:i]
			if _, ok := wordSet[prefix]; ok {
				suffix := s[i:]
				for _, sentence := range backtrack(suffix) {
					result = append(result, prefix+" "+sentence)
				}
			}
		}

		memo[s] = result
		return result
	}

	return backtrack(s)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BoundaryOfBinaryTree - https://leetcode.com/problems/boundary-of-binary-tree/
func BoundaryOfBinaryTree(root *TreeNode) []int {
	boundary := []int{root.Val}

	isLeaf := func(node *TreeNode) bool {
		return node.Left == nil && node.Right == nil
	}
	if isLeaf(root) {
		return boundary
	}

	var leftB, rightB, leaves func(node *TreeNode)

	leftB = func(node *TreeNode) {
		if node == nil {
			return
		}

		if isLeaf(node) {
			return
		}

		boundary = append(boundary, node.Val)
		if node.Left != nil {
			leftB(node.Left)
		} else {
			leftB(node.Right)
		}
	}

	rightB = func(node *TreeNode) {
		if node == nil {
			return
		}

		if isLeaf(node) {
			return
		}

		if node.Right != nil {
			rightB(node.Right)
		} else {
			rightB(node.Left)
		}
		boundary = append(boundary, node.Val)
	}

	leaves = func(node *TreeNode) {
		if node == nil {
			return
		}

		if isLeaf(node) {
			boundary = append(boundary, node.Val)
		}

		leaves(node.Left)
		leaves(node.Right)
	}

	leftB(root.Left)
	leaves(root)
	rightB(root.Right)

	return boundary
}

// FindAllConcatenatedWordsInADict - https://leetcode.com/problems/concatenated-words/
func FindAllConcatenatedWordsInADict(words []string) []string {
	n := len(words)
	wordSet := make(map[string]struct{}, n)
	for _, word := range words {
		wordSet[word] = struct{}{}
	}

	canConcat := func(w string) bool {
		n := len(w)
		dp := make([]bool, n+1)
		dp[0] = true

		for i := 1; i <= n; i++ {
			for j := 0; j < i; j++ {
				if _, ok := wordSet[w[j:i]]; ok && dp[j] {
					dp[i] = true
					break
				}
			}
		}

		return dp[n]
	}

	var result []string
	for _, word := range words {
		delete(wordSet, word)
		if canConcat(word) {
			result = append(result, word)
		}
		wordSet[word] = struct{}{}
	}

	return result
}

// SequentialDigits - https://leetcode.com/problems/sequential-digits/
func SequentialDigits(low, high int) []int {
	const maxLen = 9
	digits := "123456789"

	var result []int
	for length := 2; length <= maxLen; length++ {
		for j := 0; j+length <= maxLen; j++ {
			num := digits[j : j+length]

			seqDigits := 0
			for _, d := range num {
				seqDigits = seqDigits*10 + int(d-'0')
			}

			if seqDigits >= low && seqDigits <= high {
				result = append(result, seqDigits)
			}
		}
	}

	return result
}
