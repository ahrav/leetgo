package leetgo

import (
	"math"
	"sort"
	"strings"
)

func MaxArea(heights []int) int {
	var total int

	left, right := 0, len(heights)-1
	for left < right {
		hl, hr := heights[left], heights[right]
		h := min(hl, hr)
		area := h * (right - left)

		if area > total {
			area = total
		}

		if hl < hr {
			left++
		} else {
			right--
		}
	}

	return total
}

func IntToRoman(number int) string {
	romanNumerals := []struct {
		symbol string
		value  int
	}{
		{symbol: "M", value: 1000},
		{symbol: "CM", value: 900},
		{symbol: "D", value: 500},
		{symbol: "CD", value: 400},
		{symbol: "C", value: 100},
		{symbol: "XC", value: 90},
		{symbol: "L", value: 50},
		{symbol: "XL", value: 40},
		{symbol: "X", value: 10},
		{symbol: "IX", value: 9},
		{symbol: "V", value: 5},
		{symbol: "IV", value: 4},
		{symbol: "I", value: 1},
	}

	var res strings.Builder
	for _, rn := range romanNumerals {
		for number >= rn.value {
			res.WriteString(rn.symbol)
			number -= rn.value
		}
	}

	return res.String()
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, idx}
		}
		m[num] = idx
	}

	return nil
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}

	return
}

func MinEatingSpeed(piles []int, h int) int {
	canFinish := func(rate int) bool {
		var hours int
		for _, pile := range piles {
			hours += (pile + rate - 1) / rate
		}
		return hours <= h
	}

	minV, maxV := 1, 0
	for _, pile := range piles {
		if pile > maxV {
			maxV = pile
		}
	}

	for minV < maxV {
		mid := minV + (maxV-minV)/2
		if canFinish(mid) {
			maxV = mid
		} else {
			minV = mid + 1
		}
	}

	return minV
}

func CountGroups(related []string) int {
	visited := make([]bool, len(related))

	var dfs func(int)
	dfs = func(user int) {
		visited[user] = true
		for idx := range related {
			if related[user][idx] == '1' && !visited[idx] {
				dfs(idx)
			}
		}
	}

	var groups int

	for idx := range related {
		if !visited[idx] {
			dfs(idx)
			groups++
		}
	}

	return groups
}

func RomanToInteger(s string) int {
	numerals := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	var total, prev int

	for i := len(s) - 1; i >= 0; i-- {
		curr, _ := numerals[s[i]]
		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}

	return total
}

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	len1, len2 := len(nums1), len(nums2)
	left, right := 0, len1

	for left <= right {
		partX := (left + right + 1) / 2
		partY := ((len1 + len2 + 1) / 2) - partX

		var maxPartLeftX, minPartRightX int
		if partX == 0 {
			maxPartLeftX = math.MinInt64
		} else {
			maxPartLeftX = nums1[partX-1]
		}

		if partX == len1 {
			minPartRightX = math.MaxInt64
		} else {
			minPartRightX = nums1[partX]
		}

		var maxPartLeftY, minPartRightY int
		if partY == 0 {
			maxPartLeftY = math.MinInt64
		} else {
			maxPartLeftY = nums2[partY-1]
		}

		if partY == len2 {
			minPartRightY = math.MaxInt64
		} else {
			minPartRightY = nums2[partY]
		}

		if maxPartLeftX <= minPartRightX && maxPartLeftX <= minPartRightY {
			if (len1+len2)%2 == 0 {
				return (float64(max(maxPartLeftX, maxPartLeftY)) + float64(min(minPartRightX, minPartRightY))) / float64(2)
			}
			return float64(max(maxPartLeftY, maxPartLeftX))
		}

		if maxPartLeftX > maxPartLeftY {
			right = partX - 1
		} else {
			left = partX + 1
		}
	}

	return 0.0
}

func MinimumDifference(nums []int) int {
	const threshold = 4
	if len(nums) <= threshold {
		return 0
	}

	top, bottom := make([]int, threshold), make([]int, threshold)
	copy(top, nums[:threshold])
	copy(bottom, nums[:threshold])

	sort.Ints(top)
	sort.Ints(bottom)

	for _, num := range nums[threshold:] {
		if num < bottom[3] {
			bottom[3] = num
			for i := 3; i > 0; i-- {
				if bottom[i] < bottom[i-1] {
					bottom[i], bottom[i-1] = bottom[i-1], bottom[i]
				}
			}
		}

		if num > top[0] {
			top[0] = num
			for i := 0; i < 3; i++ {
				if top[i] > top[i+1] {
					top[i], top[i+1] = top[i+1], top[i]
				}
			}
		}
	}

	return min(
		top[0]-bottom[0],
		top[1]-bottom[1],
		top[2]-bottom[2],
		top[3]-bottom[3],
	)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseSList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func IsPalindromeSList(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	slow, mid, fast := head, head, head.Next
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}

	mid = ReverseSList(mid.Next)

	for mid != nil {
		if mid.Val != slow.Val {
			return false
		}
		mid = mid.Next
		slow = slow.Next
	}

	return true
}

func AddTwoNumbersLong(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	addToTail := func(node *ListNode) {
		if head == nil {
			head = node
			return
		}

		curr := head
		for curr != nil && curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
		return
	}

	var carry bool
	l1Curr, l2Curr := l1, l2
	for l1Curr != nil && l2Curr != nil {
		sum := l1Curr.Val + l2Curr.Val
		if carry {
			sum += 1
		}

		if sum >= 10 {
			carry = true
		} else {
			carry = false
		}

		sum = sum % 10
		addToTail(&ListNode{Val: sum})
		l1Curr = l1Curr.Next
		l2Curr = l2Curr.Next
	}

	for l1Curr != nil {
		sum := l1Curr.Val
		if carry {
			sum += 1
		}

		if sum >= 10 {
			carry = true
		} else {
			carry = false
		}

		sum = sum % 10
		addToTail(&ListNode{Val: sum})
		l1Curr = l1Curr.Next
	}

	for l2Curr != nil {
		sum := l2Curr.Val
		if carry {
			sum += 1
		}

		if sum >= 10 {
			carry = true
		} else {
			carry = false
		}

		sum = sum % 10
		addToTail(&ListNode{Val: sum})
		l2Curr = l2Curr.Next
	}

	if carry {
		addToTail(&ListNode{Val: 1})
	}

	return head
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	var carry int8

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += int8(l1.Val)
			l1 = l1.Next
		}

		if l2 != nil {
			sum += int8(l2.Val)
			l2 = l2.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: int(sum % 10)}
		curr = curr.Next
	}

	return head.Next
}

func GroupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}
		for _, char := range str {
			count[char-'a']++
		}

		m[count] = append(m[count], str)
	}

	res := make([][]string, 0, len(m))
	for _, val := range m {
		res = append(res, val)
	}

	return res
}

func MaxProfit(prices []int) int {
	maxProfit, minBuyPrice := 0, prices[0]

	for _, price := range prices[1:] {
		if price < minBuyPrice {
			minBuyPrice = price
		} else if price-minBuyPrice > maxProfit {
			maxProfit = price - minBuyPrice
		}
	}

	return maxProfit
}

type DoubleLinkNode struct {
	Key  int
	Val  int
	Next *DoubleLinkNode
	Prev *DoubleLinkNode
}
type DoublyLinkedList struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
	size int
}

func (dll *DoublyLinkedList) RemoveNode(item *DoubleLinkNode) {
	if item == dll.head {
		dll.head = item.Next
	}
	if item == dll.tail {
		dll.tail = item.Prev
	}

	if item.Prev != nil {
		item.Prev.Next = item.Next
	}
	if item.Next != nil {
		item.Next.Prev = item.Prev
	}
	item.Prev = nil
	item.Next = nil
	dll.size--

	if dll.size == 0 {
		dll.head = nil
		dll.tail = nil
	}
}

func (dll *DoublyLinkedList) AddToHead(item *DoubleLinkNode) {
	item.Prev = nil
	item.Next = dll.head

	if dll.head == nil {
		dll.tail = item
	} else {
		dll.head.Prev = item
	}

	dll.head = item
	dll.size++
}

func (dll *DoublyLinkedList) RemoveTail() *DoubleLinkNode {
	prevTail := dll.tail
	dll.tail = prevTail.Prev

	if dll.tail == nil {
		dll.head = nil
	} else {
		dll.tail.Next = nil
	}

	prevTail.Prev = nil
	prevTail.Next = nil
	dll.size--

	return prevTail
}

type LRUCache struct {
	lut    map[int]*DoubleLinkNode
	lst    *DoublyLinkedList
	cap    int
	size   int
	isFull bool
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		lut: make(map[int]*DoubleLinkNode, capacity),
		lst: new(DoublyLinkedList),
		cap: 5,
	}

	return cache
}

func (c *LRUCache) Get(key int) int {
	node, exists := c.lut[key]
	if !exists {
		return -1
	}

	c.lst.RemoveNode(node)
	c.lst.AddToHead(node)
	return node.Val
}

func (c *LRUCache) Put(key, value int) {
	if node, exists := c.lut[key]; exists {
		c.lst.RemoveNode(node)
		c.lst.AddToHead(node)
		if node.Val != value {
			node.Val = value
			c.lut[key] = node
		}
		return
	}

	node := &DoubleLinkNode{Val: value, Key: key}
	if !c.isFull {
		c.size++
		if c.size == c.cap {
			c.isFull = true
		}
	} else {
		prev := c.lst.RemoveTail()
		delete(c.lut, prev.Key)
	}

	c.lst.AddToHead(node)
	c.lut[key] = node

	return
}

type MatrixGraph struct {
	vertices int
	matrix   [][]int
}

func (g *MatrixGraph) AddEdge(v, w int) {
	g.matrix[v][w] = 1
	g.matrix[w][v] = 1
}

func (g *MatrixGraph) BFS(start int) {
	visited := make([]bool, g.vertices)
	queue := []int{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if !visited[curr] {
			visited[curr] = true

			for i := 0; i < g.vertices; i++ {
				if g.matrix[curr][i] == 1 && !visited[i] {
					queue = append(queue, i)
				}
			}
		}
	}
}

func (g *MatrixGraph) DFS(start int) {
	visited := make([]bool, g.vertices)

	var dfs func(int)
	dfs = func(v int) {
		if visited[v] {
			return
		}

		visited[v] = true
		for i := 0; i < g.vertices; i++ {
			if g.matrix[v][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	dfs(start)
}

type AdjListGraph struct {
	vertices int
	adjList  map[int][]int
}

func (g *AdjListGraph) AddEdge(v, w int) {
	g.adjList[v] = append(g.adjList[v], w)
	g.adjList[w] = append(g.adjList[w], v)
}

func (g *AdjListGraph) BFS(start int) {
	visited := make(map[int]bool, g.vertices)
	queue := []int{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if !visited[curr] {
			visited[curr] = true

			for _, neighbor := range g.adjList[curr] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
}

func (g *AdjListGraph) DFS(start int) {
	visited := make(map[int]bool, g.vertices)

	var dfs func(int)
	dfs = func(v int) {
		if visited[v] {
			return
		}

		visited[v] = true
		for _, neighbor := range g.adjList[v] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
}

func NumIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	if rows == 0 && cols == 0 {
		return 0
	}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeBFS(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
}

func TreeDFS(root *TreeNode) {
	if root == nil {
		return
	}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			dfs(node.Left)
		}

		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}

	return nil
}

func ProductExceptSelf(nums []int) []int {
	res := make([]int, 0, len(nums))

	ptr := 1
	for _, num := range nums {
		res = append(res, ptr)
		ptr *= num
	}

	ptr = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= ptr
		ptr *= nums[i]
	}

	return res
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var charCount [26]int
	for i := 0; i < len(s); i++ {
		charCount[s[i]-'a']++
		charCount[t[i]-'a']--
	}

	for _, cnt := range charCount {
		if cnt != 0 {
			return false
		}
	}

	return true
}
