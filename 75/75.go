package seventyfive

import (
	"container/heap"
	"encoding/binary"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

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

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	return str1[:gcd(len(str1), len(str2))]
}

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxV := candies[0]

	n := len(candies)
	for i := 1; i < n; i++ {
		if candies[i] > maxV {
			maxV = candies[i]
		}
	}

	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = candies[i]+extraCandies >= maxV
	}

	return res
}

func CanPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if n == 0 || size == 0 {
		return true
	}

	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 {
			leftEmpty := (i == 0) || flowerbed[i-1] == 0
			rightEmpty := (i == size-1) || flowerbed[i+1] == 0

			if leftEmpty && rightEmpty {
				flowerbed[i] = 1
				i++
				n--
			}

			if n == 0 {
				return true
			}
		}
	}

	return false
}

func ReverseVowels(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	vowels := [128]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	res := []byte(s)
	lp, rp := 0, n-1
	for lp < rp {
		for lp < rp && !vowels[res[lp]] {
			lp++
		}
		for lp < rp && !vowels[res[rp]] {
			rp--
		}

		if lp < rp {
			res[lp], res[rp] = res[rp], res[lp]
			lp++
			rp--
		}
	}

	return string(res)
}

func ReverseWords(s string) string {
	chars := []byte(s)

	reverse := func(start, end int) {
		for start < end {
			chars[start], chars[end] = chars[end], chars[start]
			start++
			end--
		}
	}

	n := len(chars)
	reverse(0, n-1)

	start, writeIdx := 0, 0
	for start < n {
		if chars[start] == ' ' {
			start++
			continue
		}

		if writeIdx != 0 {
			chars[writeIdx] = ' '
			writeIdx++
		}

		end := start
		for end < n && chars[end] != ' ' {
			chars[writeIdx] = chars[end]
			end++
			writeIdx++
		}

		reverse(writeIdx-(end-start), writeIdx-1)
		start = end
	}

	return string(chars[:writeIdx])
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

func IsSubsequence(s string, t string) bool {
	slen, tlen := len(s), len(t)
	if slen > tlen {
		return false
	}

	sp, tp := 0, 0
	for sp < slen && tp < tlen {
		if s[sp] == t[tp] {
			sp++
			tp++
		} else {
			tp++
		}
	}

	return sp == slen
}

func FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if k > n {
		return 0
	}

	var sum int
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < n; i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func MaxArea(height []int) int {
	n := len(height)

	lp, rp, maxArea := 0, n-1, 0
	for lp < rp {
		width := rp - lp
		minH := 0

		if height[lp] > height[rp] {
			minH = height[rp]
			rp--
		} else {
			minH = height[lp]
			lp++
		}

		area := width * minH
		if width*minH > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func MaxOperationsTwoPointer(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == k {
			return 1
		}
		return 0
	}

	sort.Ints(nums)

	var operations int
	lp, rp := 0, n-1
	for lp < rp {
		sum := nums[lp] + nums[rp]
		if sum == k {
			operations++
			lp++
			rp--
			continue
		}

		if sum > k {
			rp--
		} else {
			lp++
		}
	}

	return operations
}

func MaxOperationsComplement(nums []int, k int) int {
	freq := make(map[int]int, len(nums)/2)
	var operations int

	for _, num := range nums {
		complement := k - num
		if count, exists := freq[complement]; exists && count > 0 {
			operations++
			freq[complement]--
		} else {
			freq[num]++
		}
	}

	return operations
}

func IncreasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, second := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}

func IncreasingTripletDP(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] >= 3 {
			return true
		}
	}

	return false
}

func Compress(chars []byte) int {
	n := len(chars)
	if n < 2 {
		return n
	}

	var rPtr, wPtr int
	for rPtr < n {
		curr := chars[rPtr]
		cnt := 0

		for rPtr < n && chars[rPtr] == curr {
			rPtr++
			cnt++
		}

		chars[wPtr] = curr
		wPtr++

		if cnt > 1 {
			str := strconv.Itoa(cnt)
			for i := 0; i < len(str); i++ {
				chars[wPtr] = str[i]
				wPtr++
			}
		}
	}

	return wPtr
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

		if window := right - left + 1; window > maxWindow {
			maxWindow = window
		}
	}

	return maxWindow
}

func FindDifference(nums1 []int, nums2 []int) [][]int {
	set := make(map[int16]uint8)
	for _, num := range nums1 {
		set[int16(num)] |= 1
	}
	for _, num := range nums2 {
		set[int16(num)] |= 2
	}

	res := make([][]int, 2)
	for k, v := range set {
		switch v {
		case 1:
			res[0] = append(res[0], int(k))
		case 2:
			res[1] = append(res[1], int(k))
		}
	}

	return res
}

func MaxVowels(s string, k int) int {
	if len(s) < 1 {
		return 0
	}

	vowels := [26]bool{0: true, 'e' - 'a': true, 'i' - 'a': true, 'o' - 'a': true, 'u' - 'a': true}
	var count int
	for i := range k {
		if vowels[s[i]-'a'] {
			count++
		}
	}

	maxVowels := count
	for i := k; i < len(s); i++ {
		if vowels[s[i-k]-'a'] {
			count--
		}
		if vowels[s[i]-'a'] {
			count++
		}
		if count > maxVowels {
			maxVowels = count
		}
	}

	return maxVowels
}

func LongestSubarray(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	var zeroCnt, maxWindow, left int
	for right, num := range nums {
		if num == 0 {
			zeroCnt++
		}

		for zeroCnt > 1 {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}

		if window := right - left; window > maxWindow {
			maxWindow = window
		}
	}

	return min(maxWindow, n-1)
}

func LargestAltitude(gain []int) int {
	var maxGain, currAlt int
	for _, g := range gain {
		currAlt += g
		if currAlt > maxGain {
			maxGain = currAlt
		}
	}

	return maxGain
}

func PivotIndex(nums []int) int {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}

func UniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, val := range arr {
		freq[val]++
	}

	set := make(map[int]struct{}, len(freq))
	for _, val := range freq {
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}

	return true
}

func CloseStrings(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	const alphabet = 26
	freq1, freq2 := make([]int, alphabet), make([]int, alphabet)
	chars1, chars2 := make([]bool, alphabet), make([]bool, alphabet)

	for i := range word1 {
		freq1[word1[i]-'a']++
		chars1[word1[i]-'a'] = true
		freq2[word2[i]-'a']++
		chars2[word2[i]-'a'] = true
	}

	for i := range alphabet {
		if chars1[i] != chars2[i] {
			return false
		}
	}

	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := range alphabet {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}

func EqualPairs(grid [][]int) int {
	n := len(grid)
	rowMap := make(map[string]int)

	// Pre-allocate a single buffer for encoding.
	buf := make([]byte, n*4) // 4 bytes per int32.

	// Encode and count rows.
	for _, row := range grid {
		for j, val := range row {
			binary.LittleEndian.PutUint32(buf[j*4:], uint32(val))
		}
		rowMap[string(buf)]++
	}

	count := 0

	// Process columns and check against encoded rows.
	for col := 0; col < n; col++ {
		for row := 0; row < n; row++ {
			binary.LittleEndian.PutUint32(buf[row*4:], uint32(grid[row][col]))
		}
		if rowCount, exists := rowMap[string(buf)]; exists {
			count += rowCount
		}
	}

	return count
}

func RemoveStars(s string) string {
	var stack []byte
	for i := range s {
		if s[i] == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}

func AsteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	for _, astr := range asteroids {
		for len(stack) > 0 && astr < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top < -astr {
				stack = stack[:len(stack)-1]
				continue
			}
			if top == -astr {
				stack = stack[:len(stack)-1]
			}
			astr = 0
		}

		if astr != 0 {
			stack = append(stack, astr)
		}
	}

	return stack
}

func DecodeString(s string) string {
	var (
		countStack []int
		strStack   []string
	)

	currStr := ""
	k := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			k = k*10 + int(char-'0')
		} else if char == '[' {
			countStack = append(countStack, k)
			strStack = append(strStack, currStr)
			k = 0
			currStr = ""
		} else if char == ']' {
			count := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			decodedStr := strings.Repeat(currStr, count)
			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currStr = prevStr + decodedStr
		} else {
			currStr += string(char)
		}
	}

	return currStr
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
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

type RecentCounter struct {
	arr []int
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{arr: make([]int, 0)}
}

func (c *RecentCounter) Ping(t int) int {
	c.arr = append(c.arr, t)

	for len(c.arr) > 0 && c.arr[0] < t-3000 {
		c.arr = c.arr[1:]
	}

	return len(c.arr)
}

func PredictPartyVictory(senate string) string {
	var dire, radiant []int

	for i, char := range senate {
		if char == 'D' {
			dire = append(dire, i)
		} else {
			radiant = append(radiant, i)
		}
	}

	n := len(senate)
	for len(dire) > 0 && len(radiant) > 0 {
		d, r := dire[0], radiant[0]
		if d < r {
			dire = append(dire, d+n)
		} else {
			radiant = append(radiant, r+n)
		}

		dire, radiant = dire[1:], radiant[1:]
	}

	if len(dire) > 0 {
		return "Dire"
	}
	return "Radiant"
}

func DeleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	tmp := &ListNode{Next: head}
	slow, fast := tmp, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next
	return tmp.Next
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}

func PairSum(head *ListNode) int {
	if head == nil {
		return 0
	}

	if head.Next == nil {
		return head.Val
	}

	reverse := func(head *ListNode) *ListNode {
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

	curr := head
	slow, fast := curr, curr.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var maxSum int
	mid, start := reverse(slow.Next), head
	for mid != nil {
		if v := mid.Val + start.Val; v > maxSum {
			maxSum = v
		}
		mid = mid.Next
		start = start.Next
	}

	return maxSum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := dfs(node.Left)
		rh := dfs(node.Right)

		return max(lh, rh) + 1
	}

	return dfs(root)
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(*TreeNode, *[]int)
	dfs = func(node *TreeNode, arr *[]int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			*arr = append(*arr, node.Val)
		}

		dfs(node.Left, arr)
		dfs(node.Right, arr)

	}

	var leaves []int
	dfs(root1, &leaves)
	leaves2 := make([]int, 0, len(leaves))
	dfs(root2, &leaves2)

	if len(leaves) != len(leaves2) {
		return false
	}

	for i, val := range leaves {
		if val != leaves2[i] {
			return false
		}
	}

	return true
}

func LeafSimilarConcurrent(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(*TreeNode, chan int)
	dfs = func(node *TreeNode, c chan int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			c <- node.Val
			return
		}

		dfs(node.Left, c)
		dfs(node.Right, c)

	}

	chan1, chan2 := make(chan int, 1), make(chan int, 1)
	go func() {
		dfs(root1, chan1)
		close(chan1)
	}()
	go func() {
		dfs(root2, chan2)
		close(chan2)
	}()

	for {
		v1, ok1 := <-chan1
		v2, ok2 := <-chan2

		if !ok1 && !ok2 {
			return true
		}

		if !ok1 || !ok2 {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
}

func GoodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, maxSoFar int) int {
		if node == nil {
			return 0
		}

		count := 0
		if node.Val >= maxSoFar {
			count++
			maxSoFar = node.Val
		}

		return count + dfs(node.Left, maxSoFar) + dfs(node.Right, maxSoFar)
	}

	return dfs(root, root.Val)
}

func PathSum(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, currSum int) int {
		if node == nil {
			return 0
		}

		currSum += node.Val
		count := prefixSum[currSum-targetSum]

		prefixSum[currSum]++

		count += dfs(node.Left, currSum)
		count += dfs(node.Right, currSum)

		prefixSum[currSum]--

		return count
	}

	return dfs(root, 0)
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func LongestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxLen int

	var dfs func(*TreeNode, bool, int)
	dfs = func(node *TreeNode, isLeft bool, length int) {
		if node == nil {
			return
		}

		if length > maxLen {
			maxLen = length
		}

		if isLeft {
			dfs(node.Right, false, length+1)
			dfs(node.Left, true, 1)
		} else {
			dfs(node.Left, true, length+1)
			dfs(node.Right, false, 1)
		}
	}

	dfs(root.Left, true, 1)
	dfs(root.Right, false, 1)

	return maxLen
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := len(queue)

		res = append(res, queue[len(queue)-1].Val)
		for i := 0; i < level; i++ {
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

	return res
}

func MaxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minLevel, currLevel, maxSum := 1, 1, root.Val
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := len(queue)

		levelSum := 0
		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			levelSum += curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			minLevel = currLevel
		}
		currLevel++
	}

	return minLevel
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	findMin := func(node *TreeNode) *TreeNode {
		curr := node
		for curr.Left != nil {
			curr = curr.Left
		}
		return curr
	}

	if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil { // Leaf node
			return nil
		}

		// Single child.
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 2 children.
		successor := findMin(root.Right)
		root.Val = successor.Val
		root.Right = DeleteNode(root.Right, successor.Val)
	}

	return root
}

func CanVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	if n < 1 {
		return true
	}

	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(room int) {
		if visited[room] {
			return
		}

		visited[room] = true
		for _, key := range rooms[room] {
			dfs(key)
		}
	}

	dfs(0)
	for _, k := range visited {
		if !k {
			return false
		}
	}

	return true
}

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)

	visited := make([]bool, n)
	var dfs func(int)
	dfs = func(city int) {
		visited[city] = true

		for neighbor := 0; neighbor < n; neighbor++ {
			if isConnected[city][neighbor] == 1 && !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	var provinceCount int
	for city := 0; city < n; city++ {
		if !visited[city] {
			provinceCount++
			dfs(city)
		}
	}

	return provinceCount
}

func MinReorder(n int, connections [][]int) int {
	adj := make([][]int, n)
	directedEdges := make(map[[2]int]struct{}, n)
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		directedEdges[[2]int{u, v}] = struct{}{}
	}

	var dfs func(int, int) int
	dfs = func(node, parent int) int {
		reversals := 0
		for _, neighbor := range adj[node] {
			if neighbor == parent {
				continue
			}

			if _, ok := directedEdges[[2]int{node, neighbor}]; ok {
				reversals++
			}

			reversals += dfs(neighbor, node)
		}
		return reversals
	}

	return dfs(0, -1)
}

func OrangesRotting(grid [][]int) int {
	type Coord struct{ x, y int }

	var queue []Coord
	freshCnt := 0

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
	directions := []Coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			currX, currY := curr.x, curr.y
			for _, dir := range directions {
				newX, newY := currX+dir.x, currY+dir.y
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && grid[newX][newY] == 1 {
					freshCnt--
					grid[newX][newY] = 2
					queue = append(queue, Coord{newX, newY})
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

func NearestExit(maze [][]byte, entrance []int) int {
	type Coord struct{ x, y int }

	directions := []Coord{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	rows, cols := len(maze), len(maze[0])

	queue := []Coord{{entrance[0], entrance[1]}}
	maze[entrance[0]][entrance[1]] = '+'

	steps := 0

	for len(queue) > 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			currX, currY := curr.x, curr.y

			for _, dir := range directions {
				newX, newY := currX+dir.x, currY+dir.y
				if newX >= 0 && newX < rows && newY >= 0 && newY < cols && maze[newX][newY] == '.' {
					if newX == 0 || newX == rows-1 || newY == 0 || newY == cols-1 {
						return steps + 1
					}

					queue = append(queue, Coord{newX, newY})
					maze[newX][newY] = '+'
				}
			}
		}
		steps++
	}

	return -1
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

func FindKthLargest(nums []int, k int) int {
	h := new(MinHeap)
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}

func GuessNumber(n int) int {
	guess := func(num int) int {
		pick := 6
		if num == pick {
			return 0
		} else if num > pick {
			return -1
		}
		return 1
	}

	low, high := 0, n
	for low <= high {
		mid := low + (high-low)/2
		val := guess(mid)

		if val == 0 {
			return mid
		} else if val == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

type SmallestInfiniteSet struct {
	heap *MinHeap
	arr  [1001]int
}

func SmallestInfiniteSetConstructor() SmallestInfiniteSet {
	h := new(MinHeap)
	heap.Init(h)

	var arr [1001]int

	for i := 1; i <= 1000; i++ {
		heap.Push(h, i)
		arr[i] = 1
	}

	return SmallestInfiniteSet{
		heap: h,
		arr:  arr,
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	v := heap.Pop(s.heap)
	s.arr[v.(int)] = 0
	return v.(int)
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if s.arr[num] == 1 {
		return
	}
	heap.Push(s.heap, num)
	s.arr[num] = 1
}

func MaxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := range nums1 {
		pairs[i] = [2]int{nums2[i], nums1[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	h := new(MinHeap)
	heap.Init(h)

	var sum, maxScore int
	for _, pair := range pairs {
		heap.Push(h, pair[1])
		sum += pair[1]

		if h.Len() > k {
			sum -= heap.Pop(h).(int)
		}

		if h.Len() == k {
			score := sum * pair[0]
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return int64(maxScore)
}

type Worker struct {
	cost  int
	index int
}

type MinWorkerHeap []Worker

func (h MinWorkerHeap) Len() int { return len(h) }

func (h MinWorkerHeap) Less(i, j int) bool {
	if h[i].cost == h[j].cost {
		return h[i].index < h[j].index
	}
	return h[i].cost < h[j].cost
}

func (h MinWorkerHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinWorkerHeap) Push(x any) { *h = append(*h, x.(Worker)) }

func (h *MinWorkerHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func TotalCost(costs []int, k int, candidates int) int64 {
	if candidates == 0 {
		return 0
	}

	n := len(costs)

	lh, rh := new(MinWorkerHeap), new(MinWorkerHeap)
	heap.Init(lh)
	heap.Init(rh)

	lp, rp := 0, n-1
	for i := 0; i < candidates && lp <= rp; i++ {
		heap.Push(lh, Worker{costs[lp], lp})
		lp++
	}
	for i := 0; i < candidates && lp <= rp; i++ {
		heap.Push(rh, Worker{costs[rp], rp})
		rp--
	}

	var totalCost int
	for i := 0; i < k; i++ {
		if rh.Len() == 0 || (lh.Len() > 0 && (*lh)[0].cost <= (*rh)[0].cost) {
			worker := heap.Pop(lh).(Worker)
			totalCost += worker.cost
			if lp <= rp {
				heap.Push(lh, Worker{costs[lp], lp})
				lp++
			}
		} else {
			worker := heap.Pop(rh).(Worker)
			totalCost += worker.cost
			if lp <= rp {
				heap.Push(rh, Worker{costs[rp], rp})
				rp--
			}
		}
	}

	return int64(totalCost)
}

func MinEatingSpeed(piles []int, h int) int {
	canFinish := func(rate int) bool {
		var hours int
		for _, pile := range piles {
			hours += (rate + pile - 1) / rate
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

func SuccessfulPairs(spells, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	sort.Ints(potions)

	successCount := func(spell int) int {
		left, right := 0, m-1
		for left <= right {
			mid := left + (right-left)/2
			if int64(potions[mid]*spell) >= success {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		return m - left
	}

	res := make([]int, 0, n)
	for _, spell := range spells {
		res = append(res, successCount(spell))
	}

	return res
}

func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
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

	var res []string
	curr := make([]byte, n)

	var backtrack func(int)
	backtrack = func(index int) {
		if index == n {
			res = append(res, string(curr))
			return
		}

		letters := mapping[digits[index]]
		for i := range letters {
			curr[index] = letters[i]
			backtrack(index + 1)
		}
	}

	backtrack(0)
	return res
}

func CombinationSum3(k int, n int) [][]int {
	var res [][]int
	var combination []int

	var backtrack func(int, int, int)
	backtrack = func(start, remaining, count int) {
		if remaining == 0 && count == k {
			tmp := make([]int, len(combination))
			copy(tmp, combination)
			res = append(res, tmp)
			return
		}

		if remaining < 0 || count > k {
			return
		}

		for i := start; i <= 9; i++ {
			combination = append(combination, i)
			backtrack(i+1, remaining-i, count+1)
			combination = combination[:len(combination)-1]
		}
	}

	backtrack(1, n, 0)
	return res
}

func TribonacciDP(n int) int {
	if n == 0 {
		return 0
	}

	dp := make([]int, max(n+1, 3))
	dp[0], dp[1], dp[2] = 0, 1, 1
	if n < 2 {
		return dp[n]
	}

	for i := 3; i <= n; i++ {
		val := 0
		for j := i - 3; j < i; j++ {
			val += dp[j]
		}
		dp[i] = val
	}

	return dp[n]
}

func TribonacciMemo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	a, b, c := 0, 1, 1

	for i := 3; i <= n; i++ {
		next := a + b + c
		a, b, c = b, c, next
	}

	return c
}

func MinCostClimbingStairsDP(cost []int) int {
	n := len(cost)

	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}

	return min(dp[n-1], dp[n-2])
}

func MinCostClimbingStairsMemo(cost []int) int {
	n := len(cost)

	prev2, prev1 := cost[0], cost[1]

	for i := 2; i < n; i++ {
		curr := cost[i] + min(prev1, prev2)
		prev2, prev1 = prev1, curr
	}

	return min(prev1, prev2)
}

func Rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	prev2, prev := 0, nums[0]
	for i := 1; i < n; i++ {
		curr := max(prev, nums[i]+prev2)
		prev2, prev = prev, curr
	}

	return prev
}

func NumTilings(n int) int {
	const mod int = 1_000_000_007

	dp := make([]int, max(n, 3))
	dp[0], dp[1], dp[2] = 1, 2, 5

	i := 3
	for i < n {
		dp[i] = ((dp[i-1] * 2) + dp[i-3]) % mod
		i++
	}
	return dp[n-1]
}

func CountBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

func SingleNumber(nums []int) int {
	var res int
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

func UniquePaths2DP(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for j := 1; j < m; j++ {
		dp[j][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func UniquePaths1DP(m int, n int) int {
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}

func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func MaxProfit(prices []int, fee int) int {
	n := len(prices)

	hold, notHold := -prices[0], 0
	for i := 1; i < n; i++ {
		hold = max(hold, notHold-prices[i])
		notHold = max(notHold, hold+prices[i]-fee)
	}

	return notHold
}

func DeleteAndEarn(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	maxV := math.MinInt
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += num
		if num > maxV {
			maxV = num
		}
	}

	prev2, prev1 := 0, m[1]
	for i := 2; i <= maxV; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+m[i])
	}

	return prev1
}

func MaximumScore(nums, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := m - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			multiplier := multipliers[i]
			right := n - 1 - (i - left)
			dp[i][left] = max(multiplier*nums[left]+dp[i+1][left+1], multiplier*nums[right]+dp[i+1][left])
		}
	}

	return dp[0][0]
}

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}

	return dp[m][n]
}

func MinFlips(a int, b int, c int) int {
	flips := 0

	for i := 0; i < 64; i++ {
		bitA, bitB, bitC := (a>>i)&1, (b>>i)&1, (c>>i)&1

		if bitC == 0 {
			if bitA == 1 {
				flips++
			}
			if bitB == 1 {
				flips++
			}
		} else {
			if bitA == 0 && bitB == 0 {
				flips++
			}
		}
	}

	return flips
}

func MaximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}

	maxSide := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

type TrieNode struct {
	children [26]*TrieNode
	// children    map[rune]*TrieNode // Use for more flexible use cases
	isEndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func TrieConstructor() Trie { return Trie{root: new(TrieNode)} }

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		c := char - 'a'
		child := node.children[c]
		if child == nil {
			node.children[c] = new(TrieNode)
		}
		node = node.children[c]
	}
	node.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		c := char - 'a'
		child := node.children[c]
		if child == nil {
			return false
		}
		node = child
	}

	return node.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		c := char - 'a'
		child := node.children[c]
		if child == nil {
			return false
		}
		node = child
	}

	return true
}

func (t *Trie) Matches(prefix string) []string {
	var res []string

	node := t.root
	for _, char := range prefix {
		c := char - 'a'
		child := node.children[c]
		if child == nil {
			return res
		}
		node = child
	}

	var dfs func(*TrieNode, []byte)
	dfs = func(node *TrieNode, bytes []byte) {
		if node == nil || len(res) >= 3 {
			return
		}

		if node.isEndOfWord {
			res = append(res, string(bytes))
		}

		for idx, child := range node.children {
			if child != nil {
				// b := append(bytes, byte('a'+idx))
				dfs(child, append(bytes, byte('a'+idx)))
			}
		}
	}
	dfs(node, []byte(prefix))

	return res
}

func SuggestedProducts(products []string, searchWord string) [][]string {
	if len(products) == 0 {
		return nil
	}

	slices.Sort(products)

	trie := TrieConstructor()
	for _, product := range products {
		trie.Insert(product)
	}

	res := make([][]string, len(searchWord))

	for idx := range searchWord {
		matches := trie.Matches(searchWord[0 : idx+1])
		if len(matches) >= 3 {
			res[idx] = matches[:3]
		} else {
			res[idx] = matches
		}
	}

	return res
}

func EraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 || n == 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	cnt := 1
	prevEnd := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] >= prevEnd {
			cnt++
			prevEnd = intervals[i][1]
		}
	}

	return n - cnt
}
