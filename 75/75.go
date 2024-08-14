package seventyfive

import (
	"encoding/binary"
	"math"
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
