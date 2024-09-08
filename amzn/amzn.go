package amzn

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"
	"strconv"
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

func LRUCacheConstructor(capacity int) LRUCache {
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

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{lut: make(map[int]int), vals: make([]int, 0)}
}

type RandomizedSet struct {
	lut  map[int]int
	vals []int
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.lut[val]; ok {
		return false
	}

	s.lut[val] = len(s.vals)
	s.vals = append(s.vals, val)

	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	idx, ok := s.lut[val]
	if !ok {
		return false
	}

	lastIdx := len(s.vals) - 1
	lastVal := s.vals[lastIdx]

	s.vals[idx] = lastVal
	s.lut[lastVal] = idx

	s.vals = s.vals[:lastIdx]
	delete(s.lut, val)

	return true
}

func (s *RandomizedSet) GetRandom() int { return s.vals[rand.Intn(len(s.vals))] }

func FizzBuzz(n int) []string {
	res := make([]string, n)

	const (
		fizzBuzz = "FizzBuzz"
		fizz     = "Fizz"
		buzz     = "Buzz"
	)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res[i-1] = fizzBuzz
			continue
		}

		if i%3 == 0 {
			res[i-1] = fizz
			continue
		}
		if i%5 == 0 {
			res[i-1] = buzz
			continue
		}

		res[i-1] = strconv.Itoa(i)
	}

	return res
}

func LongestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func CoinChange(coins []int, total int) int {
	dp := make([]int, total+1)
	dp[0] = 0
	for i := 1; i <= total; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i <= total; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	res := dp[total]
	if res == math.MaxInt {
		return -1
	}
	return res
}

func BoundaryOfBinaryTree(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	boundary := []int{root.Val}

	isLeaf := func(node *TreeNode) bool {
		return node.Left == nil && node.Right == nil
	}

	var leftBoundary, leaves, rightBoundary func(node *TreeNode)

	leftBoundary = func(node *TreeNode) {
		if node == nil {
			return
		}
		if isLeaf(node) {
			return
		}

		boundary = append(boundary, node.Val)
		if node.Left != nil {
			leftBoundary(node.Left)
		} else {
			leftBoundary(node.Right)
		}
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

	rightBoundary = func(node *TreeNode) {
		if node == nil {
			return
		}

		if isLeaf(node) {
			return
		}

		if node.Right != nil {
			rightBoundary(node.Right)
		} else {
			rightBoundary(node.Left)
		}
		boundary = append(boundary, node.Val)
	}

	leftBoundary(root.Left)
	leaves(root)
	rightBoundary(root.Right)

	return boundary
}

func PreOrderTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	PreOrderTree(root.Left)
	PreOrderTree(root.Right)
}

func InOrderTree(root *TreeNode) {
	if root == nil {
		return
	}

	InOrderTree(root.Left)
	fmt.Println(root.Val)
	InOrderTree(root.Right)
}

func PostOrderTree(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrderTree(root.Left)
	PostOrderTree(root.Right)
	fmt.Println(root.Val)
}

func NumberOfDistinctIslands(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	m := make(map[string]struct{})

	var dfs func(int, int, string) string
	dfs = func(i, j int, direction string) string {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] != 1 {
			return ""
		}

		grid[i][j] = 0
		return direction +
			dfs(i+1, j, "U") +
			dfs(i-1, j, "D") +
			dfs(i, j+1, "R") +
			dfs(i, j-1, "L") +
			"0"
	}

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				dir := dfs(i, j, "")
				m[dir] = struct{}{}
			}
		}
	}

	return len(m)
}

func BinaryTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := BinaryTreeHeight(root.Left)
	rh := BinaryTreeHeight(root.Right)

	if lh > rh {
		return lh + 1
	}

	return rh + 1
}

func BinaryTreeDiameter(root *TreeNode) int {
	maxPath := 0
	if root == nil {
		return maxPath
	}

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := dfs(node.Left)
		rh := dfs(node.Right)

		maxPath = max(maxPath, lh+rh)
		return max(lh, rh) + 1
	}

	return dfs(root)
}

func DistanceK(root *TreeNode, target *TreeNode, k int) []int {
	if root == nil {
		return nil
	}

	parents := make(map[*TreeNode]*TreeNode)

	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node *TreeNode, parent *TreeNode) {
		if node == nil {
			return
		}

		parents[node] = parent
		dfs(node.Left, node)
		dfs(node.Right, node)
	}

	dfs(root, nil)

	visited := make(map[*TreeNode]bool)
	visited[target] = true
	dist := 0
	queue := []*TreeNode{target}

	for len(queue) > 0 && dist < k {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.Left != nil && !visited[curr.Left] {
				visited[curr.Left] = true
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil && !visited[curr.Right] {
				visited[curr.Right] = true
				queue = append(queue, curr.Right)
			}

			if parent := parents[curr]; parent != nil && !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
			}
		}
		dist++
	}

	res := make([]int, 0, len(queue))
	for _, node := range queue {
		res = append(res, node.Val)
	}

	return res
}

func LongestOnes(nums []int, k int) int {
	var left, zeroCount, maxWindow int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if window := right - left - 1; window > maxWindow {
			maxWindow = window
		}
	}

	return maxWindow
}

func LengthOfLongestSubstring(s string) int {
	var left, maxLen int

	m := make(map[byte]int)
	for right := range s {
		if idx, ok := m[s[right]]; ok && idx >= left {
			left = idx + 1
		}

		if window := right - left + 1; window > maxLen {
			maxLen = window
		}

		m[s[right]] = right
	}

	return maxLen
}

func MostVisitedPattern(usernames, websites []string, timestamps []int) []string {
	type visit struct {
		website   string
		timestamp int
	}

	userVisits := make(map[string][]visit)
	for i, username := range usernames {
		userVisits[username] = append(userVisits[username], visit{
			website:   websites[i],
			timestamp: timestamps[i],
		})
	}

	var maxPat string
	var maxCount int
	patternCount := make(map[string]int)
	for _, visits := range userVisits {
		sort.Slice(visits, func(i, j int) bool {
			return visits[i].timestamp < visits[j].timestamp
		})

		patterns := make(map[string]bool)

		for i := 0; i < len(visits)-2; i++ {
			for j := i + 1; j < len(visits)-1; j++ {
				for k := j + 1; k < len(visits); k++ {
					patterns[fmt.Sprintf("%s,%s,%s", visits[i].website, visits[j].website, visits[k].website)] = true
				}
			}
		}

		for pat := range patterns {
			patternCount[pat]++
			if count, _ := patternCount[pat]; count > maxCount || (count == maxCount && pat < maxPat) {
				maxPat = pat
				maxCount = count
			}
		}
	}

	return strings.Split(maxPat, ",")
}

func SortJumbled(mapping []int, nums []int) []int {
	type pair struct {
		originalIndex int
		mappedVal     int
	}

	pairs := make([]pair, 0, len(nums))
	for idx, num := range nums {
		var mapped int
		multiplier := 1
		if num == 0 {
			mapped = mapping[0]
		} else {
			for num > 0 {
				digit := num % 10
				mapped += mapping[digit] * multiplier
				multiplier *= 10
				num /= 10
			}
		}

		pairs = append(pairs, pair{
			originalIndex: idx,
			mappedVal:     mapped,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].mappedVal == pairs[j].mappedVal {
			return pairs[i].originalIndex < pairs[j].originalIndex
		}

		return pairs[i].mappedVal < pairs[j].mappedVal
	})

	res := make([]int, 0, len(nums))
	for _, pair := range pairs {
		res = append(res, nums[pair.originalIndex])
	}

	return res
}

func PlatesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	if n == 0 {
		return nil
	}

	prefixSum := make([]int, n+1)
	prevCandles, nextCandles := make([]int, n), make([]int, n)

	lastCandle := -1
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i]
		if s[i] == '*' {
			prefixSum[i+1]++
		} else {
			lastCandle = i
		}
		prevCandles[i] = lastCandle
	}

	lastCandle = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			lastCandle = i
		}
		nextCandles[i] = lastCandle
	}

	res := make([]int, len(queries))
	for idx, query := range queries {
		left, right := query[0], query[1]
		leftCandle, rightCandle := nextCandles[left], prevCandles[right]

		if leftCandle != -1 && rightCandle != -1 && leftCandle < rightCandle {
			res[idx] = prefixSum[rightCandle] - prefixSum[leftCandle]
		}
	}

	return res
}

func LongestPalindromeDP(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	start, maxLen := 0, 1
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] == true {
				dp[i][j] = true
				if length > maxLen {
					maxLen = length
					start = i
				}
			}
		}
	}

	return s[start : start+maxLen]
}

func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}

		return right - left - 1
	}

	var startIdx, maxLength int
	for i := 0; i < n; i++ {
		oddLen := expandAroundCenter(i, i)
		evenLen := expandAroundCenter(i, i+1)

		maxL := max(oddLen, evenLen)
		if maxL > maxLength {
			maxLength = maxL
			startIdx = i - (maxL-1)/2
		}
	}

	return s[startIdx : startIdx+maxLength]
}

func LetterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return nil
	}

	mapping := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string
	curr := make([]byte, n)
	var backtrack func(int)
	backtrack = func(index int) {
		if index == n {
			result = append(result, string(curr))
			return
		}

		letters := mapping[digits[index]]
		for i := range letters {
			curr[index] = letters[i]
			backtrack(index + 1)
		}
	}

	backtrack(0)
	return result
}

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func SearchRotatedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func NumTeams(rating []int) int {
	if len(rating) < 3 {
		return 0
	}

	n := len(rating)

	leftSmaller, leftLarger := make([]int, n), make([]int, n)
	rightSmaller, rightLarger := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				leftSmaller[i]++
			} else {
				leftLarger[i]++
			}
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if rating[j] > rating[i] {
				rightLarger[i]++
			} else {
				rightSmaller[i]++
			}
		}
	}

	var cnt int
	for i := 1; i < n-1; i++ {
		cnt += (leftSmaller[i] * rightLarger[i]) + (leftLarger[i] * rightSmaller[i])
	}

	return cnt
}

func ThreeSum(nums []int) [][]int {
	slices.Sort(nums)

	var result [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}

		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := intervals[:0]
	result = append(result, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			result = append(result, []int{curr[0], curr[1]})
		}
	}

	return result
}

func MergeIntervals2(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := intervals[:0]
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] <= end {
			end = max(end, intervals[i][1])
		} else {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	result = append(result, []int{start, end})

	return result
}

func LowestCommonAncestorDFS(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestorDFS(root.Left, p, q)
	right := LowestCommonAncestorDFS(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func TopKFrequent(nums []int, k int) []int {
	freqCount := make(map[int]int)
	for _, num := range nums {
		freqCount[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range freqCount {
		buckets[freq] = append(buckets[freq], num)
	}

	res := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, buckets[i]...)
		if len(res) > k {
			res = res[:k]
			break
		}
	}

	return res
}

func NextGreaterElements(nums []int) []int {
	n := len(nums)
	if n == 1 {
		return []int{-1}
	}

	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}

	var stack []int
	for i := range 2 * n {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}

		if i < n {
			stack = append(stack, i)
		}
	}

	return res
}

func WidthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type QueueItem struct {
		node     *TreeNode
		level    int
		position int
	}

	maxWdith := 0
	queue := []QueueItem{{root, 0, 0}}
	for len(queue) > 0 {
		levelSize := len(queue)
		firstPos := queue[0].position

		for i := 0; i < levelSize; i++ {
			item := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				maxWdith = max(maxWdith, item.position-firstPos+1)
			}

			if item.node.Left != nil {
				queue = append(queue, QueueItem{item.node.Left, item.level + 1, item.position * 2})
			}
			if item.node.Right != nil {
				queue = append(queue, QueueItem{item.node.Right, item.level + 1, item.position*2 + 1})
			}
		}
	}

	return maxWdith
}

func Jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	var jumps, currEnd, currFarthest int
	for i := 0; i < n-1; i++ {
		currFarthest = max(currFarthest, i+nums[i])
		if i == currEnd {
			jumps++
			currEnd = currFarthest
			if currEnd >= n-1 {
				break
			}
		}
	}

	return jumps
}

func MinimumAddedCoins(coins []int, target int) int {
	sort.Ints(coins)

	var maxT, cIdx, added int
	for curr := 1; curr <= target; curr++ {
		for len(coins) > cIdx && coins[cIdx] <= curr {
			maxT += coins[cIdx]
			cIdx++
		}

		if maxT < curr {
			maxT += curr
			added++
		}
	}

	return added
}

func WordBreak(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	m := make(map[string]struct{})
	for _, word := range wordDict {
		m[word] = struct{}{}
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if _, ok := m[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func IsValidParenthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s)/2)
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmpHead := new(ListNode)
	curr := tmpHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return tmpHead.Next
}

func CanJumpBackwards(nums []int) bool {
	n := len(nums)
	lastPost := n - 1

	for i := lastPost; i >= 0; i-- {
		if i+nums[i] >= lastPost {
			lastPost = i
		}
	}

	return lastPost == 0
}

func CanJumpForwards(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	var maxReach int
	for i := 0; i <= maxReach && i < n; i++ {
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= n-1 {
			return true
		}
	}

	return false
}

func MergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)

	var sb strings.Builder
	sb.Grow(len1 + len2)

	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			sb.WriteByte(word1[i])
		}

		if i < len2 {
			sb.WriteByte(word2[i])
		}
	}

	return sb.String()
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	curr := head
	for curr != nil {
		newNode := &Node{Val: curr.Val}
		newNode.Next = curr.Next
		curr.Next = newNode
		curr = newNode.Next
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	tmpHead := new(Node)
	copyCurr, curr := tmpHead, head
	for curr != nil {
		copyCurr.Next = curr.Next
		curr.Next = curr.Next.Next
		copyCurr = copyCurr.Next
		curr = curr.Next
	}

	return tmpHead.Next
}

type ParkingSystem struct{ spaces [3]uint16 }

func ParkingSystemConstructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{spaces: [3]uint16{uint16(big), uint16(medium), uint16(small)}}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	idx := carType - 1
	if ps.spaces[idx] > 0 {
		ps.spaces[idx] = ps.spaces[idx] - 1
		return true
	}

	return false
}

func NQueensAllSolutions(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	constructSolution := func() []string {
		solution := make([]string, 0, n)
		for i := range board {
			solution = append(solution, string(board[i]))
		}

		return solution
	}

	isValid := func(row, col int) bool {
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	var result [][]string
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			result = append(result, constructSolution())
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)
	return result
}

func NQueensFirstSolution(n int) []string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	constructSolution := func() []string {
		solution := make([]string, 0, n)
		for i := range board {
			solution = append(solution, string(board[i]))
		}

		return solution
	}

	isValid := func(row, col int) bool {
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	var result []string

	var backtrack func(int) bool
	backtrack = func(row int) bool {
		if row == n {
			result = constructSolution()
			return true
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				board[row][col] = 'Q'
				if backtrack(row + 1) {
					return true
				}
				board[row][col] = '.'
			}
		}

		return false
	}

	backtrack(0)
	return result
}

func Sudoku(board [][]int) [][]int {
	const size = 9

	findEmptyCell := func() (int, int) {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if board[i][j] == 0 {
					return i, j
				}
			}
		}

		return -1, -1
	}

	isValid := func(row, col, num int) bool {
		for i := 0; i < size; i++ {
			if board[row][i] == num {
				return false
			}
		}

		for i := 0; i < size; i++ {
			if board[i][col] == num {
				return false
			}
		}

		startRow := row - row%3
		startCol := col - col%3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i+startRow][j+startCol] == num {
					return false
				}
			}
		}

		return true
	}

	var backtrack func() bool
	backtrack = func() bool {
		row, col := findEmptyCell()
		if row == -1 && col == -1 {
			return true
		}

		for num := 1; num <= size; num++ {
			if isValid(row, col, num) {
				board[row][col] = num

				if backtrack() {
					return true
				}

				board[row][col] = 0
			}
		}

		return false
	}

	if backtrack() {
		return board
	}

	return nil
}

func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	if root == nil {
		return maxDiameter
	}

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := dfs(node.Left)
		rh := dfs(node.Right)

		maxDiameter = max(maxDiameter, lh+rh)
		return max(lh, rh) + 1
	}

	dfs(root)
	return maxDiameter
}

func IsValidSudoku(board [][]byte) bool {
	const size = 9

	hasDuplicates := func(arr [size]byte) bool {
		seen := make(map[byte]struct{})
		for i := range arr {
			if arr[i] != '.' {
				if _, exists := seen[arr[i]]; exists {
					return true
				} else {
					seen[arr[i]] = struct{}{}
				}
			}
		}

		return false
	}

	for i := range board {
		if hasDuplicates([9]byte(board[i])) {
			return false
		}
	}

	var arr [size]byte
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			arr[j] = board[i][j]
		}

		if hasDuplicates(arr) {
			return false
		}
	}

	for i := 0; i < size; i += 3 {
		for j := 0; j < size; j += 3 {
			idx := 0
			for m := i; m < i+3; m++ {
				for n := j; n < j+3; n++ {
					arr[idx] = board[m][n]
					idx++
				}
			}

			if hasDuplicates(arr) {
				return false
			}
		}
	}

	return true
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func FindKthLargest(nums []int, k int) int {
	h := new(IntHeap)
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func LongestCommonPrefix(strs []string) string {
	var builder strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		s := strs[0][i]
		for idx := range strs {
			if len(strs[idx]) <= i || strs[idx][i] != s {
				return builder.String()
			}
		}
		builder.WriteByte(s)
	}

	return builder.String()
}

func MoveZeroes(nums []int) {
	n := len(nums)
	var insertIdx int
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[insertIdx], nums[i] = nums[i], nums[insertIdx]
			insertIdx++
		}
	}
}

func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var total int
	for idx := range boxTypes {
		if truckSize >= boxTypes[idx][0] {
			total += boxTypes[idx][0] * boxTypes[idx][1]
			truckSize -= boxTypes[idx][0]
		} else {
			total += truckSize * boxTypes[idx][1]
			break
		}
	}

	return total
}

type MinStack struct {
	stack    []int
	minStack []int
}

func MinStackConstructor() MinStack { return MinStack{} }

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || val <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, val)
	}
}

func (s *MinStack) Pop() {
	if s.stack[len(s.stack)-1] == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int { return s.stack[len(s.stack)-1] }

func (s *MinStack) GetMin() int { return s.minStack[len(s.minStack)-1] }

func RotateArray(arr []int, k int) {
	n := len(arr)
	if n < 2 {
		return
	}

	d := k % n
	if d == 0 {
		return
	}

	revArr := func(arr []int) {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
	}

	revArr(arr)
	revArr(arr[:d])
	revArr(arr[d:])
}

func ReverseLinkedList(head *ListNode) *ListNode {
	if head == nil {
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

func CountPrimes(n int) int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}

	var total int
	for i := 2; i < n; i++ {
		if isPrime[i] {
			total++
		}
	}

	return total
}

func MinSwaps(nums []int) int {
	n := len(nums)

	var totalOnes int
	for i := range n {
		if nums[i] == 1 {
			totalOnes++
		}
	}

	if totalOnes == n || totalOnes == 0 {
		return 0
	}

	var onesCount int
	for i := range totalOnes {
		if nums[i] == 1 {
			onesCount++
		}
	}

	maxOnes := onesCount
	for i := range n {
		onesCount -= nums[i]
		onesCount += nums[(i+totalOnes)%n]
		maxOnes = max(maxOnes, onesCount)
	}

	return totalOnes - maxOnes
}

type WordFreqs []WordFreq

type WordFreq struct {
	word string
	freq int
}

func (w WordFreqs) Len() int { return len(w) }

func (w WordFreqs) Less(i, j int) bool {
	if w[i].freq == w[j].freq {
		return w[i].word > w[j].word
	}

	return w[i].freq < w[j].freq
}

func (w WordFreqs) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func (w *WordFreqs) Push(x any) { *w = append(*w, x.(WordFreq)) }

func (w *WordFreqs) Pop() any {
	old := *w
	n := len(old)
	x := old[n-1]
	*w = old[:n-1]
	return x
}

func TopKFrequentWords(words []string, k int) []string {
	freqCount := make(map[string]int)
	for _, word := range words {
		freqCount[word]++
	}

	h := new(WordFreqs)
	heap.Init(h)

	for word, freq := range freqCount {
		if h.Len() < k {
			heap.Push(h, WordFreq{word, freq})
			continue
		}
		top := (*h)[0]
		if top.freq < freq || (top.freq == freq && word < top.word) {
			heap.Pop(h)
			heap.Push(h, WordFreq{word, freq})
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(WordFreq).word
	}

	return res
}

func MyAtoi(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	result, idx, sign := 0, 0, 1

	for idx < n && s[idx] == ' ' {
		idx++
	}

	if idx < n && (s[idx] == '-' || s[idx] == '+') {
		if s[idx] == '-' {
			sign = -1
		}
		idx++
	}

	for idx < n {
		digit := int(s[idx] - '0')
		if digit < 0 || digit > 9 {
			break
		}

		// Tricky tricky ;)
		if result > math.MaxInt32 || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		idx++
	}

	result *= sign
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		course, req := pair[0], pair[1]
		graph[req] = append(graph[req], course) // prereq -> course dep
	}

	visitCnt := make([]int, numCourses)
	var stack []int

	var dfs func(int) bool
	dfs = func(course int) bool {
		if visitCnt[course] == 1 {
			return false // visiting again
		}

		if visitCnt[course] == 2 {
			return true // visited
		}

		visitCnt[course] = 1 // visiting
		for _, req := range graph[course] {
			if !dfs(req) {
				return false
			}
		}

		stack = append(stack, course)
		visitCnt[course] = 2 // visited
		return true
	}

	for course := 0; course < numCourses; course++ {
		if visitCnt[course] == 0 {
			if !dfs(course) {
				return []int{} // cycle detected, no solution
			}
		}
	}

	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}

	return stack
}

func FindOrderBFS(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)

	for _, pair := range prerequisites {
		course, req := pair[0], pair[1]
		graph[req] = append(graph[req], course)
		inDegrees[course]++
	}

	var queue []int
	for i := range numCourses {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	var order []int
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		order = append(order, course)
		for _, dep := range graph[course] {
			inDegrees[dep]--
			if inDegrees[dep] == 0 {
				queue = append(queue, dep)
			}
		}
	}

	if len(order) != numCourses {
		return []int{}
	}

	return order
}

func FirstUniqChar(s string) int {
	var charCount [26]int

	for i := range s {
		charCount[s[i]-'a']++
	}

	for i := range s {
		if charCount[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func RotateImageClockW(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}

	// Transpose.
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse rows.
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

func RotateImageCounterClockW(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}

	// Reverse rows.
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}

	// Transpose.
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func RotateImageOneEighty(matrix [][]int) {
	n := len(matrix)
	if n < 2 {
		return
	}

	// Rotate rows.
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}

	// Rotate columns.
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
}

func FindRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	eq, rot90, rot180, rot270 := true, true, true, true

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				eq = false
			}

			if mat[i][j] != target[j][n-1-i] {
				rot90 = false
			}

			if mat[i][j] != target[n-1-i][n-1-j] {
				rot180 = false
			}

			if mat[i][j] != target[n-1-j][i] {
				rot270 = false
			}

			if !(eq || rot90 || rot180 || rot270) {
				return false
			}
		}
	}

	return true
}

type CharFreq struct {
	char byte
	freq int
}

type MaxHeap []CharFreq

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(CharFreq)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func ReorganizeString(s string) string {
	if len(s) < 2 {
		return s
	}

	charFreq := make(map[byte]int)
	for i := range s {
		charFreq[s[i]]++
		if charFreq[s[i]] > (len(s)+1)/2 {
			return ""
		}
	}

	h := new(MaxHeap)
	heap.Init(h)

	for char, freq := range charFreq {
		heap.Push(h, CharFreq{char, freq})
	}

	var prev CharFreq
	var result strings.Builder
	for h.Len() > 0 {
		curr := heap.Pop(h).(CharFreq)
		result.WriteByte(curr.char)

		if prev.freq > 0 {
			heap.Push(h, prev)
		}

		curr.freq--
		prev = curr
	}

	str := result.String()
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			return ""
		}
	}

	return str
}

func SubarraySum(nums []int, k int) int {
	count, sum := 0, 0
	sumFreq := make(map[int]int)
	sumFreq[0] = 1

	for _, num := range nums {
		sum += num
		if freq, exists := sumFreq[sum-k]; exists {
			count += freq
		}

		sumFreq[sum]++
	}

	return count
}

func NumberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	var (
		ones      = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
		tens      = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		thousands = []string{"", "Thousand", "Million", "Billion"}
	)

	var helper func(int) string
	helper = func(i int) string {
		if i == 0 {
			return ""
		} else if i < 10 {
			return ones[i] + " "
		} else if i < 20 {
			return teens[i-10] + " "
		} else if i < 100 {
			return tens[i/10] + " " + helper(i%10)
		}

		return ones[i/100] + " Hundred " + helper(i%100)
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

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MinMeetingRoomsHeap(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	h := new(MinHeap)
	heap.Init(h)
	heap.Push(h, intervals[0][1])

	maxRooms := 1
	for i := 1; i < n; i++ {
		if h.Len() > 0 && (*h)[0] <= intervals[i][0] {
			heap.Pop(h)
		}

		heap.Push(h, intervals[i][1])
		if h.Len() > maxRooms {
			maxRooms = h.Len()
		}
	}

	return maxRooms
}

func MinMeetingRoomSweep(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 1
	}

	starts, ends := make([]int, n), make([]int, n)
	for i := range intervals {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	sp, ep := 0, 0
	rooms, maxRooms := 0, 1
	for sp < n {
		if starts[sp] < ends[ep] {
			sp++
			rooms++
		} else {
			ep++
			rooms--
		}

		if rooms > maxRooms {
			maxRooms = rooms
		}
	}

	return maxRooms
}

func IsBalancedBinaryTree(root *TreeNode) bool {
	var dfs func(*TreeNode) (bool, int)
	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}

		leftBal, lh := dfs(node.Left)
		if !leftBal {
			return false, lh
		}

		rightBal, rh := dfs(node.Right)
		if !rightBal {
			return false, rh
		}

		if math.Abs(float64(lh-rh)) > 1 {
			return false, 0
		}

		return true, max(lh, rh) + 1
	}

	isBal, _ := dfs(root)
	return isBal
}

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, pair := range prerequisites {
		course, req := pair[0], pair[1]
		graph[req] = append(graph[req], course)
	}

	visited := make([]bool, numCourses)
	recurStack := make([]bool, numCourses)

	var hasCycle func(int) bool
	hasCycle = func(course int) bool {
		if recurStack[course] {
			return true // visiting again
		}

		if visited[course] {
			return false // visited
		}

		visited[course] = true
		recurStack[course] = true

		for _, rec := range graph[course] {
			if hasCycle(rec) {
				return true
			}
		}

		recurStack[course] = false
		return false
	}

	for course := 0; course < numCourses; course++ {
		if hasCycle(course) {
			return false
		}
	}

	return true
}

func MaxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func CharacterReplacement(s string, k int) int {
	var charCount [26]int

	var left, maxCount, maxLen int
	for right := 0; right < len(s); right++ {
		charCount[s[right]-'A']++
		maxCount = max(maxCount, charCount[s[right]-'A'])

		if right-left+1-maxCount > k {
			charCount[s[left]-'A']--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func OrangesRotting(grid [][]int) int {
	type Coord struct{ x, y int }
	var queue []Coord

	var freshCnt int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				freshCnt++
			} else if grid[i][j] == 2 {
				queue = append(queue, Coord{i, j})
			}
		}
	}

	if freshCnt == 0 {
		return 0
	}

	minutes := -1

	rows, cols := len(grid), len(grid[0])
	directions := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, coord := range directions {
				newX, newY := curr.x+coord.x, curr.y+coord.y
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && grid[newX][newY] == 1 {
					grid[newX][newY] = 2
					freshCnt--
					queue = append(queue, Coord{newX, newY})
				}
			}
		}

		minutes++
	}

	if freshCnt > 0 {
		return -1
	}

	return minutes
}

func SpiralMatrix2(n int) [][]int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1

	num, target := 1, n*n
	for num <= target {
		for i := left; i <= right; i++ {
			mat[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			mat[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			mat[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			mat[i][left] = num
			num++
		}
		left++
	}

	return mat
}

func SpiralOrder(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])

	top, bottom, left, right := 0, rows-1, 0, cols-1

	result := make([]int, 0, rows*cols)
	target := rows * cols
	for len(result) < target {
		for i := left; i <= right && len(result) < target; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i := top; i <= bottom && len(result) < target; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		for i := right; i >= left && len(result) < target; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--

		for i := bottom; i >= top && len(result) < target; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}

	return result
}

func KthFactor(n, k int) int {
	var factors []int

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)

			if i != n/i { // Make sure we avoid adding perfect squares twice
				factors = append(factors, n/i)
			}
		}
	}

	sort.Ints(factors)
	if len(factors) >= k {
		return factors[k-1]
	}

	return -1
}

func PartitionString(s string) int {
	cnt, mask := 1, 0

	for i := range s {
		bitPos := 1 << (s[i] - 'a')
		if bitPos&mask != 0 {
			cnt++
			mask = 0
		}
		mask |= bitPos
	}

	return cnt
}

func MinSwapsNoWrap(data []int) int {
	onesCnt := 0
	for _, v := range data {
		onesCnt += v
	}

	if onesCnt == 1 {
		return 0
	}

	currSum := 0
	for i := 0; i < onesCnt; i++ {
		currSum += data[i]
	}

	minSwaps := onesCnt - currSum
	for i := onesCnt; i < len(data); i++ {
		currSum += data[i] - data[i-onesCnt]
		if swaps := onesCnt - currSum; swaps < minSwaps {
			minSwaps = swaps
		}
	}

	return minSwaps
}

type Node2 struct {
	Val    int
	Left   *Node2
	Right  *Node2
	Parent *Node2
}

func LowestCommonAncestorWithParent(p, q *Node2) *Node2 {
	a, b := p, q

	for a != b {
		if a != nil {
			a = a.Parent
		} else {
			a = q
		}

		if b != nil {
			b = b.Parent
		} else {
			b = p
		}
	}

	return a
}

func AppendCharacters(s string, t string) int {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}

	return len(t) - j
}

func MaximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)

	if k > n {
		return 0
	}

	currSum, maxSum := 0, 0
	seenCnt := make(map[int]int, k)

	for i := 0; i < k; i++ {
		currSum += nums[i]
		seenCnt[nums[i]]++
	}

	if len(seenCnt) == k {
		maxSum = currSum
	}

	for i := k; i < n; i++ {
		currSum -= nums[i-k]
		seenCnt[nums[i-k]]--

		if seenCnt[nums[i-k]] == 0 {
			delete(seenCnt, nums[i-k])
		}

		currSum += nums[i]
		seenCnt[nums[i]]++

		if len(seenCnt) == k && currSum > maxSum {
			maxSum = currSum
		}
	}

	return int64(maxSum)
}

func FirstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; {
		if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

func NumberOfWaysSlice(s string) int64 {
	n := len(s)

	zeros, ones := make([]int64, n+1), make([]int64, n+1)
	for i, c := range s {
		zeros[i+1] = zeros[i]
		ones[i+1] = ones[i]
		if c == '0' {
			zeros[i+1]++
		} else {
			ones[i+1]++
		}
	}

	var result int64
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			result += ones[i] * (ones[n] - ones[i])
		} else {
			result += zeros[i] * (zeros[n] - zeros[i])
		}
	}

	return result
}

func NumberOfWaysCounters(s string) int64 {
	var zerosAfter, zerosBefore, onesAfter, onesBefore int64
	for _, c := range s {
		if c == '0' {
			zerosAfter++
		} else {
			onesAfter++
		}
	}

	var result int64
	for _, c := range s {
		if c == '0' {
			result += onesBefore * onesAfter

			zerosAfter--
			zerosBefore++
		} else {
			result += zerosBefore * zerosAfter

			onesAfter--
			onesBefore++
		}
	}

	return result
}

func ReorderLogFiles(logs []string) []string {
	var letterLogs, digitLogs []string

	for _, log := range logs {
		splitIdx := strings.Index(log, " ")
		content := log[splitIdx+1:]

		if content[0] >= '0' && content[0] <= '9' {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}

	sort.SliceStable(letterLogs, func(i, j int) bool {
		splitIdxI := strings.Index(letterLogs[i], " ")
		contentI := letterLogs[i][splitIdxI+1:]
		iID := letterLogs[i][:splitIdxI]

		splitIdxJ := strings.Index(letterLogs[j], " ")
		contenJ := letterLogs[j][splitIdxJ+1:]
		jID := letterLogs[j][:splitIdxJ]

		if contentI != contenJ {
			return contentI < contenJ
		}

		return iID < jID
	})

	return append(letterLogs, digitLogs...)
}

// CountTheNumOfKFreeSubsets - https://leetcode.com/problems/count-the-number-of-k-free-subsets/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func CountTheNumOfKFreeSubsets(nums []int, k int) int64 {
	fib := func(v int) int {
		if v == 1 {
			return 1
		}
		if v == 2 {
			return 1
		}

		a, b := 1, 2
		for i := 3; i <= v; i++ {
			a, b = b, a+b
		}

		return b
	}

	sort.Ints(nums)
	m := make(map[int]int)
	for _, num := range nums {
		m[num] = m[num-k] + 1
		delete(m, num-k)
	}

	result := int64(1)
	for _, v := range m {
		result *= int64(fib(v + 1))
	}

	return result
}

// MissingNumber - https://leetcode.com/problems/missing-number/
func MissingNumber(nums []int) int {
	n := len(nums)

	for i := 0; i < n; {
		if nums[i] < n && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}

	for i := range nums {
		if nums[i] != i {
			return i
		}
	}

	return n
}

func MissingNumberMath(nums []int) int {
	n := len(nums)

	expectedSum := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return expectedSum - sum
}

// FindDuplicateFloyd - https://leetcode.com/problems/find-the-duplicate-number/
func FindDuplicateFloyd(nums []int) int {
	slow, fast := nums[0], nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

func FindDuplicateBinarySearch(nums []int) int {
	left, right := 1, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}

		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// MinimumKeypresses - https://leetcode.com/problems/minimum-number-of-keypresses/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func MinimumKeypresses(s string) int {
	var freqCnt [26]int
	for i := range s {
		freqCnt[s[i]-'a']++
	}

	freqs := make([]int, 0, len(freqCnt))
	for _, freq := range freqCnt {
		if freq != 0 {
			freqs = append(freqs, freq)
		}
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i] > freqs[j]
	})

	var result int
	for i, freq := range freqs {
		if i < 9 {
			result += freq * 1
		} else if i < 18 {
			result += freq * 2
		} else {
			result += freq * 3
		}
	}

	return result
}

// MinCost - https://leetcode.com/problems/jump-game-viii/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func MinCost(nums []int, costs []int) int64 {
	n := len(nums)

	var inc, dec []int
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for len(inc) > 0 && nums[i] >= nums[inc[len(inc)-1]] {
			idx := inc[len(inc)-1]
			inc = inc[:len(inc)-1]
			dp[i] = min(dp[i], dp[idx]+costs[i])
		}

		for len(dec) > 0 && nums[i] < nums[dec[len(dec)-1]] {
			idx := dec[len(dec)-1]
			dec = dec[:len(dec)-1]
			dp[i] = min(dp[i], dp[idx]+costs[i])
		}

		inc = append(inc, i)
		dec = append(dec, i)
	}

	return int64(dp[n-1])
}

// MakePalindrome - https://leetcode.com/problems/valid-palindrome-iv/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func MakePalindrome(s string) bool {
	n := len(s)

	cnt := 0
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			cnt++
		}

		if cnt > 2 {
			return false
		}
	}

	return true
}

// MinimumOperations - https://leetcode.com/problems/merge-operations-to-turn-array-into-a-palindrome/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func MinimumOperations(nums []int) int {
	n := len(nums)

	cnt := 0
	for i, j := 0, n-1; i < j; {
		if nums[i] == nums[j] {
			i++
			j--
			continue
		}

		if nums[i] < nums[j] {
			nums[i+1] += nums[i]
			i++
		} else {
			nums[j-1] += nums[j]
			j--
		}
		cnt++
	}

	return cnt
}

// MinimumSwaps - https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func MinimumSwaps(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	minIdx := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[minIdx] {
			minIdx = i
		}
	}

	maxIdx := n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	minSwaps := minIdx
	maxSwaps := n - 1 - maxIdx

	totalSwaps := minSwaps + maxSwaps

	if minIdx < maxIdx {
		return totalSwaps
	}
	return totalSwaps - 1
}

// CanCompleteCircuit - https://leetcode.com/problems/gas-station/
func CanCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	maxGas, maxCost := 0, 0
	for i := range n {
		maxGas += gas[i]
		maxCost += cost[i]
	}

	if maxCost > maxGas {
		return -1
	}

	currGas, startIdx := 0, 0
	for i := 0; i < n; i++ {
		currGas += gas[i] - cost[i]

		if currGas < 0 {
			startIdx = i + 1
			currGas = 0
		}
	}

	return startIdx
}

type TrieNode struct {
	children map[byte]*TrieNode
	word     string
}

func (t *TrieNode) Insert(word string) {
	curr := t
	for i := range word {
		char := word[i]
		if curr.children[char] == nil {
			curr.children[char] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		curr = curr.children[char]
	}
	curr.word = word
}

// FindWords - https://leetcode.com/problems/word-search-ii/
func FindWords(board [][]byte, words []string) []string {
	trie := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, word := range words {
		trie.Insert(word)
	}

	rows, cols := len(board), len(board[0])
	var result []string

	var dfs func(x, y int, node *TrieNode)
	dfs = func(x, y int, node *TrieNode) {
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] == '#' {
			return
		}

		char := board[x][y]
		child := node.children[char]
		if child == nil {
			return
		}

		if child.word != "" {
			result = append(result, child.word)
			child.word = ""
		}

		board[x][y] = '#'
		dfs(x-1, y, child)
		dfs(x+1, y, child)
		dfs(x, y-1, child)
		dfs(x, y+1, child)
		board[x][y] = char
	}

	for i := range rows {
		for j := range cols {
			dfs(i, j, trie)
		}
	}

	return result
}

// SnakesAndLadders - https://leetcode.com/problems/snakes-and-ladders/
func SnakesAndLadders(board [][]int) int {
	n := len(board)

	getCoord := func(pos int) (int, int) {
		row := (pos - 1) / n
		col := (pos - 1) % n

		if row%2 == 1 {
			col = n - 1 - col
		}

		return n - 1 - row, col
	}

	queue := list.New()
	queue.PushBack([2]int{1, 0})
	boardSize := n * n
	visited := make([]bool, boardSize+1)
	visited[1] = true

	for queue.Len() > 0 {
		front := queue.Remove(queue.Front()).([2]int)

		curr, moves := front[0], front[1]
		for nextPos := curr + 1; nextPos <= min(curr+6, boardSize); nextPos++ {
			row, col := getCoord(nextPos)

			next := nextPos
			if board[row][col] != -1 {
				next = board[row][col]
			}

			if next == boardSize {
				return moves + 1
			}

			if !visited[next] {
				visited[next] = true
				queue.PushBack([2]int{next, moves + 1})
			}
		}
	}

	return -1
}

// Exist - https://leetcode.com/problems/word-search/?envType=problem-list-v2&envId=954v5ops
func Exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs func(x, y, idx int) bool
	dfs = func(x, y, idx int) bool {
		if idx == len(word) {
			return true
		}

		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] != word[idx] {
			return false
		}

		char := board[x][y]
		board[x][y] = '#'
		found := dfs(x-1, y, idx+1) || dfs(x+1, y, idx+1) || dfs(x, y-1, idx+1) || dfs(x, y+1, idx+1)
		board[x][y] = char

		return found
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// GenerateParenthesis - https://leetcode.com/problems/generate-parentheses/description/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func GenerateParenthesis(n int) []string {
	validSize := n * 2

	var result []string
	var backtrack func(curr string, open, close int)
	backtrack = func(curr string, open, close int) {
		if len(curr) == validSize {
			result = append(result, curr)
			return
		}

		if open < n {
			backtrack(curr+"(", open+1, close)
		}

		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}

// UpdateBoard - https://leetcode.com/problems/minesweeper/description/?envType=problem-list-v2&envId=954v5ops
func UpdateBoard(board [][]byte, click []int) [][]byte {
	rowClick, colClick := click[0], click[1]

	if board[rowClick][colClick] == 'M' {
		board[rowClick][colClick] = 'X'
		return board
	}

	rows, cols := len(board), len(board[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	mineCount := func(x, y int) int {
		cnt := 0
		for _, dir := range directions {
			newRow, newCol := x+dir[0], y+dir[1]
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || board[newRow][newCol] != 'M' {
				continue
			}
			cnt++
		}
		return cnt
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] != 'E' {
			return
		}

		if cnt := mineCount(x, y); cnt > 0 {
			board[x][y] = byte('0' + cnt)
		} else {
			board[x][y] = 'B'

			for _, dir := range directions {
				newRow, newCol := x+dir[0], y+dir[1]
				dfs(newRow, newCol)
			}
		}
	}

	dfs(rowClick, colClick)

	return board
}

// LongestValidParentheses - https://leetcode.com/problems/longest-valid-parentheses/
func LongestValidParentheses(s string) int {
	n := len(s)

	if n < 2 {
		return 0
	}

	stack := []int{-1}
	maxLen := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}

		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			maxLen = max(maxLen, i-stack[len(stack)-1])
		}
	}

	return maxLen
}

// LongestValidParentheses - https://leetcode.com/problems/longest-valid-parentheses/
func LongestValidParenthesesTwoPass(s string) int {
	n := len(s)

	if n < 2 {
		return 0
	}

	left, right := 0, 0
	maxLen := 0

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2*right)
		} else if right > left {
			right, left = 0, 0
		}
	}

	left, right = 0, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			maxLen = max(maxLen, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}

// GameOfLife - https://leetcode.com/problems/game-of-life/?envType=problem-list-v2&envId=954v5ops
func GameOfLife(board [][]int) {
	rows, cols := len(board), len(board[0])

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	countAliveNeighbors := func(x, y int) int {
		aliveCnt := 0
		for _, dir := range directions {
			newRow, newCol := x+dir[0], y+dir[1]
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}

			if board[newRow][newCol] == 1 || board[newRow][newCol] == 2 {
				aliveCnt++
			}
		}
		return aliveCnt
	}

	for i := range rows {
		for j := range cols {
			aliveCnt := countAliveNeighbors(i, j)

			if board[i][j] == 1 {
				if aliveCnt < 2 || aliveCnt > 3 {
					board[i][j] = 2
				}
			} else {
				if aliveCnt == 3 {
					board[i][j] = -1
				}
			}
		}
	}

	for i := range rows {
		for j := range cols {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}

	return
}

// IsSameTree - https://leetcode.com/problems/same-tree/description/?envType=problem-list-v2&envId=954v5ops
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTree(p.Left, p.Right) && IsSameTree(q.Left, q.Right)
}

// MinPathSum - https://leetcode.com/problems/minimum-path-sum/?envType=problem-list-v2&envId=954v5ops
func MinPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	for i := 1; i < cols; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for j := 1; j < rows; j++ {
		grid[j][0] += grid[j-1][0]
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += min(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[rows-1][cols-1]
}

// RotateRight - https://leetcode.com/problems/rotate-list/description/?envType=problem-list-v2&envId=954v5ops
func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	size := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		size++
	}

	tail.Next = head

	k = k % size
	if k == 0 {
		tail.Next = nil
		return head
	}

	newTail := head
	for range size - k - 1 {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

// PathSum - https://leetcode.com/problems/path-sum-ii/?envType=problem-list-v2&envId=954v5ops
func PathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int

	var dfs func(node *TreeNode, sum int, arr []int)
	dfs = func(node *TreeNode, sum int, arr []int) {
		if node == nil {
			return
		}

		sum += node.Val
		arr = append(arr, node.Val)

		if node.Left == nil && node.Right == nil && sum == targetSum {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
			return
		}

		dfs(node.Left, sum, arr)
		dfs(node.Right, sum, arr)
	}

	dfs(root, 0, []int{})
	return result
}

// LevelOrder - https://leetcode.com/problems/binary-tree-level-order-traversal/?envType=problem-list-v2&envId=954v5ops
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		var levelResults []int
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levelResults = append(levelResults, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelResults)
	}

	return result
}

// SortedListToBST - https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/?envType=problem-list-v2&envId=954v5ops
func SortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	findMiddle := func(head *ListNode) *ListNode {
		var prev *ListNode
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}

		if prev != nil {
			prev.Next = nil
		}
		return slow
	}

	mid := findMiddle(head)
	root := &TreeNode{Val: mid.Val}

	root.Left = SortedListToBST(head)
	root.Right = SortedListToBST(mid.Next)

	return root
}

// IsSymmetric - https://leetcode.com/problems/symmetric-tree/?envType=problem-list-v2&envId=954v5ops
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	return isMirror(root.Left, root.Right)
}

// IsValidBST - https://leetcode.com/problems/validate-binary-search-tree/?envType=problem-list-v2&envId=954v5ops
func IsValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(node *TreeNode, lower, upper int) bool
	dfs = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}

		if node.Val <= lower || node.Val >= upper {
			return false
		}

		if !dfs(node.Left, lower, node.Val) {
			return false
		}

		if !dfs(node.Right, node.Val, upper) {
			return false
		}

		return true
	}

	return dfs(root, math.MinInt, math.MaxInt)
}

// Candy - https://leetcode.com/problems/candy/?envType=problem-list-v2&envId=954v5ops
func Candy(ratings []int) int {
	n := len(ratings)

	if n < 2 {
		return n
	}

	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	var total int
	for _, cnt := range candies {
		total += cnt
	}

	return total
}

// Trap - https://leetcode.com/problems/trapping-rain-water/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	maxLeft, maxRight := height[left], height[right]

	water := 0
	for left < right {
		if maxLeft < maxRight {
			left++

			if height[left] < maxLeft {
				water += maxLeft - height[left]
			} else {
				maxLeft = height[left]
			}
		} else {
			right--

			if height[right] < maxRight {
				water += maxRight - height[right]
			} else {
				maxRight = height[right]
			}
		}
	}

	return water
}

// PacificAtlantic - https://leetcode.com/problems/pacific-atlantic-water-flow/?envType=problem-list-v2&envId=954v5ops
func PacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])

	if rows == 0 || cols == 0 {
		return [][]int{}
	}

	// Invert the problem statement and work from the oceans inward.
	// This simplifies the number of directions we need to compare. (Move toward higher heights)

	pacific, atlantic := make([][]bool, rows), make([][]bool, rows)
	for i := range pacific {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(x, y int, ocean [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		ocean[x][y] = true

		for _, dir := range directions {
			newRow, newCol := x+dir[0], y+dir[1]
			if newRow < 0 ||
				newRow >= rows ||
				newCol < 0 ||
				newCol >= cols ||
				ocean[newRow][newCol] ||
				heights[newRow][newCol] < heights[x][y] {
				continue
			}
			dfs(newRow, newCol, ocean)
		}
	}

	for i := range rows {
		dfs(i, 0, pacific)
		dfs(i, cols-1, atlantic)
	}

	for j := range cols {
		dfs(0, j, pacific)
		dfs(rows-1, j, atlantic)
	}

	var result [][]int
	for i := range rows {
		for j := range cols {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}

// SortColors - https://leetcode.com/problems/sort-colors/?envType=problem-list-v2&envId=954v5ops (Dutch Flag Problem)
func SortColors(nums []int) {
	n := len(nums)
	l, m, h := 0, 0, n-1
	for m <= h {
		switch nums[m] {
		case 0:
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		case 1:
			m++
		case 2:
			nums[m], nums[h] = nums[h], nums[m]
			h--
		}
	}
}

// RestoreIpAddresses - https://leetcode.com/problems/restore-ip-addresses/description/?envType=problem-list-v2&envId=954v5ops
func RestoreIpAddresses(s string) []string {
	n := len(s)

	if n < 4 || n > 12 {
		return []string{}
	}

	isValid := func(segment string) bool {
		n := len(segment)
		if n == 0 || n > 3 {
			return false
		}

		if segment[0] == '0' && n > 1 {
			return false
		}

		num, err := strconv.Atoi(segment)
		if err != nil {
			return false
		}

		return num >= 0 && num <= 255
	}

	var result []string
	var backtrack func(start int, path []string)
	backtrack = func(start int, path []string) {
		if len(path) == 4 && start == n {
			result = append(result, fmt.Sprintf("%s.%s.%s.%s", path[0], path[1], path[2], path[3]))
			return
		}

		if len(path) == 4 {
			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > n {
				break
			}

			segment := s[start : start+i]
			if isValid(segment) {
				backtrack(start+i, append(path, segment))
			}
		}
	}

	backtrack(0, []string{})
	return result
}

type ListNodeMinHeap []*ListNode

func (h ListNodeMinHeap) Len() int           { return len(h) }
func (h ListNodeMinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeMinHeap) Push(x any) { *h = append(*h, x.(*ListNode)) }
func (h *ListNodeMinHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n-1]
	return x
}

// MergeKLists - https://leetcode.com/problems/merge-k-sorted-lists/description/?envType=problem-list-v2&envId=954v5ops
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	h := new(ListNodeMinHeap)
	heap.Init(h)

	for _, lst := range lists {
		if lst != nil {
			heap.Push(h, lst)
		}
	}

	tmp := &ListNode{}
	curr := tmp

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)

		if node.Next != nil {
			heap.Push(h, node.Next)
			node.Next = nil
		}

		curr.Next = node
		curr = node
	}

	return tmp.Next
}

type MinHeap2 []int
type MaxHeap2 []int

func (h MinHeap2) Len() int { return len(h) }
func (h MaxHeap2) Len() int { return len(h) }

func (h MinHeap2) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap2) Less(i, j int) bool { return h[i] > h[j] }

func (h MinHeap2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap2) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap2) Push(x any) { *h = append(*h, x.(int)) }
func (h *MaxHeap2) Push(x any) { *h = append(*h, x.(int)) }

func (h *MinHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MaxHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianFinder struct {
	minH    *MinHeap2
	minSize int

	maxH    *MaxHeap2
	maxSize int
}

func MedianFinderConstructor() MedianFinder {
	minH, maxH := new(MinHeap2), new(MaxHeap2)
	heap.Init(minH)
	heap.Init(maxH)

	return MedianFinder{
		minH: minH,
		maxH: maxH,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.maxH.Len() == 0 || num <= (*mf.maxH)[0] {
		heap.Push(mf.maxH, num)
		mf.maxSize++
	} else {
		heap.Push(mf.minH, num)
		mf.minSize--
	}

	if mf.maxSize > mf.minSize+1 {
		heap.Push(mf.minH, heap.Pop(mf.maxH).(int))
		mf.maxSize--
		mf.minSize++
	} else if mf.minSize > mf.maxSize {
		heap.Push(mf.maxH, heap.Pop(mf.minH).(int))
		mf.maxSize++
		mf.minSize--
	}

}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxSize > mf.minSize {
		return float64((*mf.maxH)[0])
	}

	return (float64((*mf.maxH)[0]) + float64((*mf.minH)[0])) / float64(2)
}

type WordDictionary struct {
	trie *TrieNodeW
}

type TrieNodeW struct {
	children map[rune]*TrieNodeW
	isWord   bool
}

func WordDictionaryConstructor() WordDictionary {
	return WordDictionary{trie: &TrieNodeW{children: make(map[rune]*TrieNodeW)}}
}

func (wd *WordDictionary) AddWord(word string) {
	curr := wd.trie
	for _, char := range word {
		if curr.children[char] == nil {
			curr.children[char] = &TrieNodeW{children: make(map[rune]*TrieNodeW)}
		}
		curr = curr.children[char]
	}
	curr.isWord = true

}

func (wd *WordDictionary) Search(word string) bool {
	n := len(word)

	var search func(idx int, node *TrieNodeW) bool
	search = func(idx int, node *TrieNodeW) bool {
		if idx == n {
			return node.isWord
		}

		char := rune(word[idx])
		if char != '.' {
			if wd.trie.children[char] == nil {
				return false
			}
			return search(idx+1, wd.trie.children[char])
		}

		for _, child := range wd.trie.children {
			if search(idx+1, child) {
				return true
			}
		}
		return false
	}

	return search(0, wd.trie)
}

type NextTreeNode struct {
	Val   int
	Left  *NextTreeNode
	Right *NextTreeNode
	Next  *NextTreeNode
}

// ConnectDFS - https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/?envType=problem-list-v2&envId=954v5ops
func ConnectDFS(root *NextTreeNode) *NextTreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	root.Left.Next = root.Right

	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	ConnectDFS(root.Left)
	ConnectDFS(root.Right)

	return root
}

// ConnectIterative - https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/?envType=problem-list-v2&envId=954v5ops
func ConnectIterative(root *NextTreeNode) *NextTreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	leftmost := root
	for leftmost.Left != nil {
		head := leftmost
		for head != nil {
			head.Left.Next = head.Right

			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftmost = leftmost.Left
	}

	return root
}

// MaxSlidingWindow - https://leetcode.com/problems/sliding-window-maximum/?envType=problem-list-v2&envId=954v5ops
func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	deque := []int{}
	result := make([]int, 0, n-k+1)

	for i := 0; i < n; i++ {
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

// Reverse - https://leetcode.com/problems/reverse-integer/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func Reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		result = result*10 + digit
		x /= 10
	}

	if result < math.MinInt32 {
		return 0
	} else if result > math.MaxInt32 {
		return 0
	}

	return result
}

// LongestConsecutive - https://leetcode.com/problems/longest-consecutive-sequence/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func LongestConsecutive(nums []int) int {
	n := len(nums)
	numSet := make(map[int]bool, n)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for _, num := range nums {
		if !numSet[num-1] {
			currNum := num
			currStreak := 1

			for numSet[currNum+1] {
				currStreak++
				currNum++
			}

			if currStreak > longest {
				longest = currStreak
			}
		}
	}

	return longest
}

// NumSquares - https://leetcode.com/problems/perfect-squares/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func NumSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			square := j * j
			dp[i] = min(dp[i], dp[i-square]+1)
		}
	}

	return dp[n]
}

type MxHeap []int

func (h MxHeap) Len() int           { return len(h) }
func (h MxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MxHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// LeastInterval - https://leetcode.com/problems/task-scheduler/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func LeastInterval(tasks []byte, n int) int {
	freqs := make(map[byte]int)
	for _, t := range tasks {
		freqs[t]++
	}

	h := new(MxHeap)
	heap.Init(h)

	for _, v := range freqs {
		heap.Push(h, v)
	}

	interval := 0
	for h.Len() > 0 {
		var tmp []int

		i := 0

		for i <= n {
			if h.Len() > 0 {
				if f := heap.Pop(h).(int); f > 1 {
					tmp = append(tmp, f-1)
				}
			}

			interval++
			if h.Len() == 0 && len(tmp) == 0 {
				break
			}
			i++
		}

		for _, v := range tmp {
			heap.Push(h, v)
		}
	}

	return interval
}

type WeightedUnionFind struct {
	parents map[string]string
	ratios  map[string]float64
}

func NewWeightedUnionfind() *WeightedUnionFind {
	return &WeightedUnionFind{
		parents: make(map[string]string),
		ratios:  make(map[string]float64),
	}
}

func (uf *WeightedUnionFind) Find(x string) (string, float64) {
	if uf.parents[x] == x {
		return x, 1.0
	}

	origParent, origRatio := uf.parents[x], uf.ratios[x]
	root, rootRatio := uf.Find(origParent)
	uf.parents[x] = root
	uf.ratios[x] = origRatio * rootRatio
	return root, uf.ratios[x]
}

func (uf *WeightedUnionFind) Union(x, y string, value float64) bool {
	rootX, ratioX := uf.Find(x)
	rootY, ratioY := uf.Find(y)

	if rootX == rootY {
		return math.Abs(ratioX/ratioY-value) < 1e-5
	}

	uf.parents[rootX] = rootY
	uf.ratios[rootX] = value * (ratioY / ratioX)
	return true
}

// CheckContradictions - https://leetcode.com/problems/check-for-contradictions-in-equations/?envType=study-plan-v2&envId=amazon-spring-23-high-frequency
func CheckContradictions(equations [][]string, values []float64) bool {
	uf := NewWeightedUnionfind()

	for i, eq := range equations {
		a, b := eq[0], eq[1]
		value := values[i]

		if _, ok := uf.parents[a]; !ok {
			uf.parents[a] = a
			uf.ratios[a] = 1.0
		}
		if _, ok := uf.parents[b]; !ok {
			uf.parents[b] = b
			uf.ratios[b] = 1.0
		}

		if !uf.Union(a, b, value) {
			return true
		}
	}

	return false
}

type CoordDist struct {
	point    []int
	distance float64
}

type MaxCoordHeap []CoordDist

func (h MaxCoordHeap) Len() int           { return len(h) }
func (h MaxCoordHeap) Less(i, j int) bool { return h[i].distance > h[j].distance }
func (h MaxCoordHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxCoordHeap) Push(x any) { *h = append(*h, x.(CoordDist)) }
func (h *MaxCoordHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// KClosest - https://leetcode.com/problems/k-closest-points-to-origin/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func KClosest(points [][]int, k int) [][]int {
	h := new(MaxCoordHeap)
	heap.Init(h)

	distance := func(x, y int) float64 {
		return math.Sqrt(float64((x * x) + (y * y)))
	}

	for _, point := range points {
		d := distance(point[0], point[1])
		heap.Push(h, CoordDist{point: point, distance: d})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([][]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(CoordDist).point)
	}

	return result
}

// SolveSudoku - https://leetcode.com/problems/sudoku-solver/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func SolveSudoku(board [][]byte) {
	const side = 9

	nextAvailableCell := func() (int, int) {
		for i := range side {
			for j := range side {
				if board[i][j] == '.' {
					return i, j
				}
			}
		}
		return -1, -1
	}

	isValid := func(row, col int, val byte) bool {
		for i := range side {
			if board[row][i] == val || board[i][col] == val {
				return false
			}
		}

		startRow := row - row%3
		startCol := col - col%3
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == val {
					return false
				}
			}
		}
		return true
	}

	var backtrack func() bool
	backtrack = func() bool {
		r, c := nextAvailableCell()
		if r == -1 || c == -1 {
			return true
		}

		for i := 1; i <= side; i++ {
			if isValid(r, c, byte(i)+'0') {
				board[r][c] = byte(i) + '0'
				if backtrack() {
					return true
				}
				board[r][c] = '.'
			}
		}

		return false
	}

	_ = backtrack()
}

// PeakIndexInMountainArray - https://leetcode.com/problems/peak-index-in-a-mountain-array/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func PeakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// LongestSubstring - https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func LongestSubstring(s string, k int) int {
	if k == 1 {
		return len(s)
	}

	var substring func(s string, k int) int
	substring = func(s string, k int) int {
		if len(s) == 0 {
			return 0
		}

		var charCount [26]int
		for i := range s {
			charCount[s[i]-'a']++
		}

		for i := range s {
			if charCount[s[i]-'a'] < k {
				left := substring(s[:i], k)
				right := substring(s[i+1:], k)
				return max(left, right)
			}
		}

		return len(s)
	}

	return substring(s, k)
}

// WordBreakII - https://leetcode.com/problems/word-break-ii/
func WordBreakII(s string, wordDict []string) []string {
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	memo := make(map[string][]string)

	var backtrack func(subStr string) []string
	backtrack = func(subStr string) []string {
		if res, exists := memo[subStr]; exists {
			return res
		}

		var result []string
		if _, exists := wordSet[subStr]; exists {
			result = append(result, subStr)
		}

		for i := 1; i < len(subStr); i++ {
			prefix := subStr[:i]
			if _, exists := wordSet[prefix]; exists {
				suffix := subStr[i:]
				sentences := backtrack(suffix)
				for _, sent := range sentences {
					result = append(result, prefix+" "+sent)
				}
			}
		}

		memo[subStr] = result
		return result
	}

	return backtrack(s)
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

// MinimumCost - https://leetcode.com/problems/minimum-cost-to-convert-string-i/
func MinimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	inf := math.MaxInt
	const charCount = 26

	minCost := make([][]int, charCount)
	for i := range minCost {
		minCost[i] = make([]int, charCount)
		for j := range minCost[i] {
			if i == j {
				minCost[i][j] = 0
			} else {
				minCost[i][j] = inf
			}
		}
	}

	n := len(original)
	for i := range n {
		x := original[i] - 'a'
		y := changed[i] - 'a'
		minCost[x][y] = min(minCost[x][y], cost[i])
	}

	for k := range charCount {
		for i := range charCount {
			for j := range charCount {
				if minCost[i][k] != inf && minCost[k][j] != inf {
					minCost[i][j] = min(minCost[i][j], minCost[i][k]+minCost[k][j])
				}
			}
		}
	}

	totalCost := 0
	for i := range source {
		x := source[i] - 'a'
		y := target[i] - 'a'
		if minCost[x][y] == inf {
			return -1
		}
		totalCost += minCost[x][y]
	}

	return int64(totalCost)
}

// LongestValidSubstring - https://leetcode.com/problems/length-of-the-longest-valid-substring/
func LongestValidSubstring(word string, forbidden []string) int {
	m := make(map[string]struct{}, len(forbidden))
	for _, str := range forbidden {
		m[str] = struct{}{}
	}

	n := len(word)
	left, maxLen := 0, 0
	for right := 0; right < n; right++ {
		for j := right; j >= max(left, right-9); j-- {
			str := word[j : right+1]

			if _, exists := m[str]; exists {
				left = j + 1
				break
			}
		}
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// FindMinMoves - https://leetcode.com/problems/super-washing-machines/
func FindMinMoves(machines []int) int {
	totalDresses := 0
	for _, cnt := range machines {
		totalDresses += cnt
	}

	n := len(machines)
	if totalDresses%n != 0 {
		return -1
	}

	target := totalDresses / n

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	minMoves, currBalance := 0, 0
	for _, cnt := range machines {
		diff := cnt - target
		currBalance += diff

		minMoves = max(minMoves, max(abs(currBalance), diff))
	}

	return minMoves
}

// UniqueLetterString - https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/
func UniqueLetterString(s string) int {
	n := len(s)
	const charCount = 26

	last := make([]int, charCount)
	secondLast := make([]int, charCount)
	for i := range last {
		last[i] = -1
		secondLast[i] = -1
	}

	sum := 0
	for i := 0; i < n; i++ {
		idx := s[i] - 'A'

		sum += (i - last[idx]) * (last[idx] - secondLast[idx])
		secondLast[idx] = last[idx]
		last[idx] = i
	}

	for i := range charCount {
		sum += (n - last[i]) * (last[i] - secondLast[i])
	}

	return sum
}

// FindAllConcatenatedWordsInADict - https://leetcode.com/problems/concatenated-words/
func FindAllConcatenatedWordsInADict(words []string) []string {
	wordSet := make(map[string]struct{}, len(words))
	for _, word := range words {
		wordSet[word] = struct{}{}
	}

	canConcat := func(word string) bool {
		n := len(word)
		dp := make([]bool, n+1)
		dp[0] = true

		for i := 1; i <= n; i++ {
			for j := 0; j < i; j++ {
				if _, ok := wordSet[word[j:i]]; ok && dp[j] {
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

type TNode struct {
	children [26]*TNode
	isWord   bool
}

func (t *TNode) Insert(word string) {
	curr := t
	for i := range word {
		if curr.children[word[i]-'a'] == nil {
			curr.children[word[i]-'a'] = new(TNode)
		}
		curr = curr.children[word[i]-'a']
	}
	curr.isWord = true
}

// FindAllConcatenatedWordsInADictTrie - https://leetcode.com/problems/concatenated-words/
func FindAllConcatenatedWordsInADictTrie(words []string) []string {
	trie := new(TNode)
	for _, word := range words {
		if word != "" {
			trie.Insert(word)
		}
	}

	var canConcat func(word string, idx, count int) bool
	canConcat = func(word string, idx, count int) bool {
		curr := trie
		n := len(word)

		for i := idx; i < n; i++ {
			child := curr.children[word[i]-'a']
			if child == nil {
				return false
			}

			curr = child

			if child.isWord {
				if i == n-1 {
					return count >= 1
				}

				if canConcat(word, i+1, count+1) {
					return true
				}
			}
		}
		return false
	}

	var result []string
	for _, word := range words {
		if canConcat(word, 0, 0) {
			result = append(result, word)
		}
	}

	return result
}

// LargestRectangeArea - https://leetcode.com/problems/largest-rectangle-in-histogram/
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 1 {
		return heights[0]
	}

	var stack []int
	maxArea := 0

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, height*width)
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]

		var width int
		if len(stack) == 0 {
			width = n
		} else {
			width = n - stack[len(stack)-1] - 1
		}

		maxArea = max(maxArea, height*width)
	}

	return maxArea
}

// TotalFruit - https://leetcode.com/problems/fruit-into-baskets/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func TotalFruit(fruits []int) int {
	n := len(fruits)
	if n < 3 {
		return n
	}

	const basketCapacity = 2
	basket := make(map[int]int, basketCapacity)
	left, maxFruit := 0, 0
	for right := 0; right < n; right++ {
		basket[fruits[right]]++

		for len(basket) > basketCapacity {
			basket[fruits[left]]--
			if basket[fruits[left]] == 0 {
				delete(basket, fruits[left])
			}
			left++
		}

		if window := right - left + 1; window > maxFruit {
			maxFruit = window
		}
	}

	return maxFruit
}

func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}

		res = append(res, left[i:]...)
		res = append(res, right[j:]...)
		return res
	}

	mid := n / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// ReversePairs - https://leetcode.com/problems/reverse-pairs/
func ReversePairs(nums []int) int {
	n := len(nums)
	tmp := make([]int, n)
	return mergeSort(nums, tmp, 0, n-1)
}

func mergeSort(nums, tmp []int, left, right int) int {
	if left >= right {
		return 0
	}

	merge := func(nums, tmp []int, left, mid, right int) int {
		i, j, k := left, mid+1, left
		count := 0
		for i <= mid && j <= right {
			if nums[i] > 2*nums[j] {
				count += mid - i + 1
				j++
			} else {
				i++
			}
		}

		i, j, k = left, mid+1, left
		for i <= mid && j <= right {
			if nums[i] < nums[j] {
				tmp[k] = nums[i]
				i++
			} else {
				tmp[k] = nums[j]
				j++
			}
			k++
		}

		for i <= mid {
			tmp[k] = nums[i]
			i++
			k++
		}

		for j <= right {
			tmp[k] = nums[j]
			j++
			k++
		}

		for i := left; i <= right; i++ {
			nums[i] = tmp[i]
		}

		return count
	}

	mid := (left + right) / 2
	count := mergeSort(nums, tmp, left, mid)
	count += mergeSort(nums, tmp, mid+1, right)

	return count + merge(nums, tmp, left, mid, right)
}

// GetMaxLen - https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/
func GetMaxLen(nums []int) int {
	maxLen := 0
	currMax, currMin := 0, 0

	for _, num := range nums {
		if num > 0 {
			currMax++
			if currMin > 0 {
				currMin++
			}
		} else if num < 0 {
			tmp := currMax
			if currMin > 0 {
				currMax = currMin + 1
			} else {
				currMax = 0
			}
			currMin = tmp + 1
		} else {
			currMax, currMin = 0, 0
		}

		if currMax > maxLen {
			maxLen = currMax
		}
	}

	return maxLen
}

// MaximumBooks - https://leetcode.com/problems/maximum-number-of-books-you-can-take/
func MaximumBooks(books []int) int64 {
	n := len(books)
	if n == 1 {
		return int64(books[0])
	}

	var stack []int
	dp := make([]int, n+1)
	maxBooks := 0

	asum := func(x int) int {
		return x * (x + 1) / 2
	}

	for i, book := range books {
		for len(stack) > 0 && books[stack[len(stack)-1]] >= book-i+stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		curr := asum(book)
		j := -1
		if len(stack) > 0 {
			j = stack[len(stack)-1]
		}

		if prevBooks := book - i + j; prevBooks > 0 {
			curr += dp[j+1] + asum(prevBooks)
		}

		dp[i+1] = curr
		if curr > maxBooks {
			maxBooks = curr
		}
		stack = append(stack, i)
	}

	return int64(maxBooks)
}

// ReachNumber - https://leetcode.com/problems/reach-a-number/
func ReachNumber(target int) int {
	if target < 0 {
		target = -target
	}

	sum, moves := 0, 0
	for sum < target || (sum-target)%2 != 0 {
		moves++
		sum += moves
	}

	return moves
}

// ClosedIsland - https://leetcode.com/problems/number-of-closed-islands/
func ClosedIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	if rows < 3 || cols < 3 {
		return 0
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] != 0 {
			return
		}

		grid[x][y] = 1
		for _, dir := range directions {
			newRow, newCol := x+dir[0], y+dir[1]
			dfs(newRow, newCol)
		}
	}

	// Flood.
	for i := range rows {
		for j := range cols {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && grid[i][j] == 0 {
				dfs(i, j)
			}
		}
	}

	numIslands := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 0 {
				dfs(i, j)
				numIslands++
			}
		}
	}

	return numIslands
}

// ShortestBridge - https://leetcode.com/problems/shortest-bridge/
func ShortestBridge(grid [][]int) int {
	n := len(grid)

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	isValid := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < n
	}

	var islands [][2]int
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if !isValid(x, y) || visited[x][y] || grid[x][y] == 0 {
			return
		}

		visited[x][y] = true
		islands = append(islands, [2]int{x, y})

		for _, dir := range directions {
			newRow, newCol := x+dir[0], y+dir[1]
			dfs(newRow, newCol)
		}
	}

	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				found = true
			}
		}
	}

	numChanges := 0
	queue := islands

	for len(queue) > 0 {
		levelSize := len(queue)
		for range levelSize {
			cell := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				newRow, newCol := cell[0]+dir[0], cell[1]+dir[1]
				if isValid(newRow, newCol) {
					if !visited[newRow][newCol] && grid[newRow][newCol] == 1 {
						return numChanges
					}
					if !visited[newRow][newCol] && grid[newRow][newCol] == 0 {
						visited[newRow][newCol] = true
						queue = append(queue, [2]int{newRow, newCol})
					}
				}
			}
		}
		numChanges++
	}

	return -1
}

// MostExpensiveItem - https://leetcode.com/problems/most-expensive-item-that-can-not-be-bought/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func MostExpensiveItem(primeOne, primeTwo int) int {
	return primeOne*primeTwo - primeOne - primeTwo
}

// MostExpensiveItemDP - https://leetcode.com/problems/most-expensive-item-that-can-not-be-bought/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days
func MostExpensiveItemDP(primeOne, primeTwo int) int {
	maxSum := primeOne * primeTwo

	dp := make([]bool, maxSum+1)
	dp[0] = true

	for i := 1; i <= maxSum; i++ {
		if i >= primeOne && dp[i-primeOne] {
			dp[i] = true
		}
		if i >= primeTwo && dp[i-primeTwo] {
			dp[i] = true
		}
	}

	for i := maxSum; i >= 0; i-- {
		if !dp[i] {
			return i
		}
	}

	return 0
}

// SubArrayRanges - https://leetcode.com/problems/sum-of-subarray-ranges/
func SubArrayRanges(nums []int) int64 {
	n := len(nums)
	stack := make([]int, 0, n)
	sum := 0

	for right := 0; right <= n; right++ {
		for len(stack) > 0 && (right == n || nums[stack[len(stack)-1]] > nums[right]) {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}

			sum -= nums[mid] * (mid - left) * (right - mid)
		}

		stack = append(stack, right)
	}

	stack = stack[:0]

	for right := 0; right <= n; right++ {
		for len(stack) > 0 && (right == n || nums[stack[len(stack)-1]] < nums[right]) {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}

			sum += nums[mid] * (mid - left) * (right - mid)
		}

		stack = append(stack, right)
	}

	return int64(sum)
}

// MinSwapsBinaryString - https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
func MinSwapsBinaryString(s string) int {
	zeroCnt, oneCnt := 0, 0
	for _, c := range s {
		if c == '1' {
			oneCnt++
		} else {
			zeroCnt++
		}
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(zeroCnt-oneCnt) > 1 {
		return -1
	}

	evenMiss, oddMiss := 0, 0
	for i, c := range s {
		if i%2 == 0 {
			if c != '0' {
				evenMiss++
			}
			if c != '1' {
				oddMiss++
			}
		} else {
			if c != '1' {
				evenMiss++
			}
			if c != 0 {
				oddMiss++
			}
		}
	}

	if zeroCnt == oneCnt {
		return min(evenMiss, oddMiss) / 2
	}

	if zeroCnt > oneCnt {
		return evenMiss / 2
	}
	return oddMiss / 2
}

// - MinimumOperationsArrayZero - https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
func MinimumOperationsArrayZero(nums []int) int {
	var set [101]bool
	set[0] = true

	counter := 0
	for _, num := range nums {
		if !set[num] {
			set[num] = true
			counter++
		}
	}

	return counter
}

// SequentialDigits - https://leetcode.com/problems/sequential-digits/
func SequentialDigits(low, high int) []int {
	var result []int

	baseDigits := "123456789"

	for length := 2; length <= 9; length++ {
		for start := 0; start+length <= 9; start++ {
			num := baseDigits[start : start+length]

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

// TriangularSum - https://leetcode.com/problems/find-triangular-sum-of-an-array/
func TriangularSum(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	newNums := make([]int, n-1)
	for i := 1; i < n; i++ {
		newNums[i-1] = (nums[i] + nums[i-1]) % 10
	}

	return TriangularSum(newNums)
}

// GoodDaysToRobBank - https://leetcode.com/problems/find-good-days-to-rob-the-bank/
func GoodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	if time == 0 {
		result := make([]int, n)
		for i := range n {
			result[i] = i
		}
		return result
	}

	dec := make([]int, n)
	inc := make([]int, n)

	for i := 1; i < n; i++ {
		if security[i-1] >= security[i] {
			dec[i] = dec[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if security[i+1] >= security[i] {
			inc[i] = inc[i+1] + 1
		}
	}

	var result []int
	for i := time; i < n-time; i++ {
		if inc[i] >= time && dec[i] >= time {
			result = append(result, i)
		}
	}

	return result
}

// MinMovesToMakePalindrome - https://leetcode.com/problems/minimum-number-of-moves-to-make-palindrome/
func MinMovesToMakePalindrome(s string) int {
	n := len(s)

	arr := []rune(s)
	lp, rp := 0, n-1
	moves := 0

	for lp < rp {
		if arr[lp] == arr[rp] {
			lp++
			rp--
			continue
		}

		left := rp
		for arr[left] != arr[lp] {
			left--
		}

		if left == lp {
			// Middle character.
			arr[left], arr[left+1] = arr[left+1], arr[left]
			moves++
		} else {
			for left < rp {
				arr[left], arr[left+1] = arr[left+1], arr[left]
				left++
				moves++
			}

			lp++
			rp--
		}
	}

	return moves
}

// Racecar - https://leetcode.com/problems/race-car/
func Racecar(target int) int {
	type State struct{ speed, pos int }

	queue := list.New()
	queue.PushBack(State{speed: 1, pos: 0})

	visited := make(map[State]bool)
	visited[State{speed: 1, pos: 0}] = true

	upperLimit := 2 * target
	steps := 0

	for queue.Len() > 0 {
		for range queue.Len() {
			curr := queue.Remove(queue.Front()).(State)

			if curr.pos == target {
				return steps
			}

			newSpeed, newPos := curr.speed*2, curr.pos+curr.speed
			newState := State{speed: newSpeed, pos: newPos}
			if newPos > 0 && newPos < upperLimit && !visited[newState] {
				visited[newState] = true
				queue.PushBack(newState)
			}

			revSpeed := -1
			if curr.speed < 0 {
				revSpeed = 1
			}
			newState = State{speed: revSpeed, pos: curr.pos}
			if !visited[newState] {
				visited[newState] = true
				queue.PushBack(newState)
			}
		}
		steps++
	}

	return -1
}

// AppealSum - https://leetcode.com/problems/total-appeal-of-a-string/
func AppealSum(s string) int64 {
	var seen [26]int
	for i := range seen {
		seen[i] = -1
	}

	total, curr := 0, 0
	for i := range s {
		if idx := seen[s[i]-'a']; idx != -1 {
			// Remove the previous contribution
			curr -= idx + 1
		}

		curr += i + 1
		total += curr

		seen[s[i]-'a'] = i
	}

	return int64(total)
}

// TicTacToe - https://leetcode.com/problems/design-tic-tac-toe/
type TicTacToe struct {
	board      [][]byte
	directions [8][2]int
	players    [3]byte
	size       int
}

func TicTacToeConstructor(n int) TicTacToe {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
	}

	players := [3]byte{'.', 'x', 'o'}
	directions := [8][2]int{
		{-1, 0}, {1, 0}, // vertical (up, down)
		{0, -1}, {0, 1}, // horizontal (left, right)
		{-1, -1}, {1, 1}, // diagonal (top-left to bottom-right)
		{-1, 1}, {1, -1}, // anti-diagonal (top-right to bottom-left)
	}

	return TicTacToe{board: board, size: n, players: players, directions: directions}
}

func (t *TicTacToe) Move(row int, col int, player int) int {
	t.board[row][col] = t.players[player]

	checkDirection := func(dir [2]int) int {
		depth := 0
		r, c := row+dir[0], col+dir[1]
		for r >= 0 && r < t.size && c >= 0 && c < t.size && t.board[r][c] == t.players[player] {
			depth++
			r += dir[0]
			c += dir[1]
		}
		return depth
	}

	for i := 0; i < len(t.directions); i += 2 {
		fwdDir, bckDir := t.directions[i], t.directions[i+1]

		fwdDepth, bckDepth := checkDirection(fwdDir), checkDirection(bckDir)
		if fwdDepth+bckDepth+1 >= t.size {
			return player
		}
	}

	return 0
}

// MinimumHealth - https://leetcode.com/problems/minimum-health-to-beat-game/
func MinimumHealth(damage []int, armor int) int64 {
	n := len(damage)
	maxV, total := damage[0], damage[0]+1

	for i := 1; i < n; i++ {
		total += damage[i]
		if d := damage[i]; d > maxV {
			maxV = d
		}
	}

	return int64(total - min(armor, maxV))
}

// EvalRPN - https://leetcode.com/problems/evaluate-reverse-polish-notation/
func EvalRPN(tokens []string) int {
	type expressionType int
	const (
		add expressionType = iota
		sub
		mult
		div
	)

	expressions := map[string]expressionType{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}

	var stack []int
	for _, tok := range tokens {
		expr, ok := expressions[tok]
		if !ok {
			val, _ := strconv.Atoi(tok)
			stack = append(stack, val)
			continue
		}

		switch expr {
		case add:
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op2+op1)
		case sub:
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op2-op1)
		case mult:
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op2*op1)
		case div:
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op2/op1)
		}
	}

	return stack[0]
}

// ReverseKGroup - https://leetcode.com/problems/reverse-nodes-in-k-group/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	getKthNode := func(start *ListNode, k int) *ListNode {
		for start != nil && k > 0 {
			start = start.Next
			k--
		}
		return start
	}

	dummy := &ListNode{Next: head}
	prevGroup := dummy

	for {
		kthNode := getKthNode(prevGroup, k)
		if kthNode == nil {
			break
		}

		nextGroup := kthNode.Next

		prev, curr := kthNode.Next, prevGroup.Next
		for curr != nextGroup {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		tmp := prevGroup.Next
		prevGroup.Next = kthNode
		prevGroup = tmp
	}

	return dummy.Next
}

// NthUglyNumber - https://leetcode.com/problems/ugly-number-ii/
func NthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1

	p2, p3, p5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5

	for i := 1; i < n; i++ {
		nextUgly := min(next2, next3, next5)
		ugly[i] = nextUgly

		if nextUgly == next2 {
			p2++
			next2 = ugly[p2] * 2
		}

		if nextUgly == next3 {
			p3++
			next3 = ugly[p3] * 3
		}

		if nextUgly == next5 {
			p5++
			next5 = ugly[p5] * 5
		}
	}

	return ugly[n-1]
}

// Insert - https://leetcode.com/problems/insert-interval/
func Insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)

	var result [][]int
	idx := 0

	for idx < n && intervals[idx][1] < newInterval[0] {
		result = append(result, intervals[idx])
		idx++
	}

	for idx < n && intervals[idx][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[idx][0])
		newInterval[1] = max(newInterval[1], intervals[idx][1])
		idx++
	}

	result = append(result, []int{newInterval[0], newInterval[1]})

	for idx < n {
		result = append(result, intervals[idx])
		idx++
	}

	return result
}

// UpdateMatrix - https://leetcode.com/problems/01-matrix/
func UpdateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])

	var queue [][2]int

	for i := range rows {
		for j := range mat[i] {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				mat[i][j] = math.MaxInt32
			}
		}
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			nr, nc := curr[0]+dir[0], curr[1]+dir[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}

			if mat[nr][nc] > mat[curr[0]][curr[1]]+1 {
				mat[nr][nc] = mat[curr[0]][curr[1]] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	return mat
}
