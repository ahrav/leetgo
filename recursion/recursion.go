package recursion

import (
	"sort"
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
