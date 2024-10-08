package recursion

import (
	"bytes"
	"sort"
	"strings"
)

func GetRow(index int) []int {
	if index == 0 {
		return []int{1}
	}

	prevRow := GetRow(index - 1)

	currRow := make([]int, index+1)
	currRow[0], currRow[index] = 1, 1

	for i := 1; i < index; i++ {
		currRow[i] = prevRow[i-1] + prevRow[i]
	}

	return currRow
}

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	half := MyPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func KthGrammar(n, k int) int {
	if n == 1 {
		return 0
	}

	length := 1 << (n - 1)
	mid := length / 2

	if k <= mid {
		return KthGrammar(n-1, k)
	} else {
		return 1 - KthGrammar(n-1, k-mid)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	dp := make([][][]*TreeNode, n+1)
	for i := range dp {
		dp[i] = make([][]*TreeNode, n+1)
	}

	var generateTrees func(start, end int) []*TreeNode
	generateTrees = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}

		if dp[start][end] != nil {
			return dp[start][end]
		}

		var allTrees []*TreeNode

		for i := start; i <= end; i++ {
			leftTrees := generateTrees(start, i-1)
			rightTrees := generateTrees(i+1, end)

			for _, left := range leftTrees {
				for _, right := range rightTrees {
					tree := &TreeNode{Val: i}
					tree.Left = left
					tree.Right = right

					allTrees = append(allTrees, tree)
				}
			}
		}

		dp[start][end] = allTrees

		return allTrees
	}

	return generateTrees(1, n)
}

func Permute(nums []int) [][]int {
	n := len(nums)

	var result [][]int
	var backtrack func(start, end int)
	backtrack = func(start, end int) {
		if start == end {
			tmp := make([]int, n)
			copy(tmp, nums)
			result = append(result, tmp)
		}

		for i := start; i < end; i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrack(start+1, end)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backtrack(0, n)
	return result
}

func PermuteUnique(nums []int) [][]int {
	n := len(nums)

	var results [][]int
	visited := make([]bool, n)

	var backtrack func(tmp []int)
	backtrack = func(tmp []int) {
		if len(tmp) == n {
			t := make([]int, n)
			copy(t, tmp)
			results = append(results, t)
			return
		}

		for i := range n {
			if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
				continue
			}
			visited[i] = true
			tmp = append(tmp, nums[i])
			backtrack(tmp)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}

	sort.Ints(nums)
	backtrack(make([]int, 0, n))

	return results
}

// CombinationSum - https://leetcode.com/problems/combination-sum/
func CombinationSum(candidates []int, target int) [][]int {
	n := len(candidates)

	var (
		result [][]int
		curr   []int
	)

	var backtrack func(start, end, sum int)
	backtrack = func(start, end, sum int) {
		if sum == target {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}

		if sum > target { // short circuit and return early
			return
		}

		for i := start; i < n; i++ {
			sum += candidates[i]
			curr = append(curr, candidates[i])
			backtrack(i, end, sum)
			sum -= curr[len(curr)-1]
			curr = curr[:len(curr)-1]
		}
	}

	sort.Ints(candidates)
	backtrack(0, n, 0)

	return result
}

// CombinationSum2 - https://leetcode.com/problems/combination-sum-ii/
func CombinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)

	var (
		result [][]int
		curr   []int
	)

	var backtrack func(start, end, sum int)
	backtrack = func(start, end, sum int) {
		if sum == target {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}

		if sum > target {
			return
		}

		for i := start; i < end; i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			curr = append(curr, candidates[i])
			backtrack(i+1, end, sum)
			sum -= curr[len(curr)-1]
			curr = curr[:len(curr)-1]
		}
	}

	sort.Ints(candidates)
	backtrack(0, n, 0)

	return result
}

// Subsets - https://leetcode.com/problems/subsets/
func Subsets(nums []int) [][]int {
	n := len(nums)

	var result [][]int

	var backtrack func(start, end int, tmp []int)
	backtrack = func(start, end int, tmp []int) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		result = append(result, t)

		for i := start; i < end; i++ {
			tmp = append(tmp, nums[i])
			backtrack(i+1, end, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backtrack(0, n, []int{})

	return result
}

// SubsetsWithDup - https://leetcode.com/problems/subsets-ii/
func SubsetsWithDup(nums []int) [][]int {
	n := len(nums)

	var result [][]int

	var backtrack func(start, end int, tmp []int)
	backtrack = func(start, end int, tmp []int) {
		t := make([]int, len(tmp))
		copy(t, tmp)
		result = append(result, t)

		for i := start; i < end; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			tmp = append(tmp, nums[i])
			backtrack(i+1, end, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	sort.Ints(nums)
	backtrack(0, n, []int{})

	return result
}

// Partition - https://leetcode.com/problems/palindrome-partitioning/
func Partition(s string) [][]string {
	isPalindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	n := len(s)

	var result [][]string

	var backtrack func(start, end int, tmp []string)
	backtrack = func(start, end int, tmp []string) {
		if start == end {
			t := make([]string, len(tmp))
			copy(t, tmp)
			result = append(result, t)
			return
		}

		for i := start; i < end; i++ {
			curr := s[start : i+1]
			if isPalindrome(curr) {
				tmp = append(tmp, curr)
				backtrack(i+1, end, tmp)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	backtrack(0, n, []string{})

	return result
}

// GeneratePalindromes - https://leetcode.com/problems/palindrome-permutation-ii/
func GeneratePalindromes(s string) []string {
	if len(s) == 0 {
		return []string{}
	}

	freqCount := make(map[byte]int)
	for i := range s {
		freqCount[s[i]]++
	}

	var mid byte
	oddCount, size := 0, 0
	for k, v := range freqCount {
		if v%2 != 0 {
			oddCount++
			mid = k
		}
		size += v

		if oddCount > 1 {
			return []string{}
		}
	}

	half := make([]byte, 0, size/2)
	for k, v := range freqCount {
		half = append(half, bytes.Repeat([]byte{k}, v/2)...)
	}

	reverse := func(str []byte) []byte {
		for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
			str[i], str[j] = str[j], str[i]
		}

		return str
	}

	n := len(s) / 2

	var result []string
	visited := make([]bool, n)

	var backtrack func(tmp []byte)
	backtrack = func(tmp []byte) {
		if len(tmp) == n {
			t := make([]byte, n)
			copy(t, tmp)

			var sb strings.Builder
			sb.WriteString(string(t))
			if oddCount != 0 {
				sb.WriteByte(mid)
			}
			sb.WriteString(string(reverse(t)))
			result = append(result, sb.String())
			return
		}

		for i := range n {
			if visited[i] || (i > 0 && half[i] == half[i-1] && !visited[i-1]) {
				continue
			}
			visited[i] = true
			tmp = append(tmp, half[i])
			backtrack(tmp)
			tmp = tmp[:len(tmp)-1]
			visited[i] = false
		}
	}

	backtrack([]byte{})
	return result
}

// MaxAreaOfIsland - https://leetcode.com/problems/max-area-of-island/
func MaxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == 0 {
			return 0
		}

		grid[x][y] = 0
		return 1 + dfs(x-1, y) + dfs(x+1, y) + dfs(x, y-1) + dfs(x, y+1)
	}

	var maxArea int
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 {
				if area := dfs(i, j); area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// NumTrees - https://leetcode.com/problems/unique-binary-search-trees/
func NumTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for node := 2; node <= n; node++ {
		for root := 1; root <= node; root++ {
			leftTrees := dp[root-1]
			rightTress := dp[node-root]
			dp[node] += leftTrees * rightTress
		}
	}

	return dp[n]
}
