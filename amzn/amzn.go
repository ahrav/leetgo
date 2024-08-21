package amzn

import (
	"container/heap"
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

	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			result = append(result, []int{curr[0], curr[1]})
		}
	}

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
