package onefifty

// MaxSubArray - https://leetcode.com/problems/maximum-subarray/?envType=study-plan-v2&envId=top-interview-150
func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	currMax, globalMax := nums[0], nums[0]
	for i := 1; i < n; i++ {
		currMax = max(nums[i], currMax+nums[i])
		if currMax > globalMax {
			globalMax = currMax
		}
	}

	return globalMax
}

// TwoSum - https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan-v2&envId=top-interview-150
func TwoSum(numbers []int, target int) []int {
	n := len(numbers)

	lp, rp := 0, n-1
	for lp < rp {
		sum := numbers[lp] + numbers[rp]
		if sum == target {
			return []int{lp + 1, rp + 1}
		}

		if sum > target {
			rp--
		} else {
			lp++
		}
	}

	return nil // Shouldn't occur
}

// Merge - https://leetcode.com/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
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
}

// RemoveDuplicatesII - https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
func RemoveDuplicatesII(nums []int) int {
	n := len(nums)
	if n == 1 || n == 2 {
		return n
	}

	lastIdx, prev, cnt := 1, nums[0], 1
	for i := 1; i < n; i++ {
		if nums[i] != prev {
			nums[lastIdx] = nums[i]
			lastIdx++
			cnt = 1
			prev = nums[i]
		} else if cnt < 2 {
			nums[lastIdx] = nums[i]
			cnt++
			lastIdx++
		}
	}

	return lastIdx
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ZigzagLevelOrder - https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result [][]int
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		level := len(queue)
		levelQ := make([]int, level)

		for i := range level {
			node := queue[0]
			queue = queue[1:]

			if leftToRight {
				levelQ[i] = node.Val
			} else {
				levelQ[level-i-1] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node)
			}
			if node.Right != nil {
				queue = append(queue, node)
			}
		}

		leftToRight = !leftToRight
		result = append(result, levelQ)
	}

	return result
}

// RemoveElement - https://leetcode.com/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func RemoveElement(nums []int, val int) int {
	idx := 0
	for _, num := range nums {
		if num != val {
			nums[idx] = num
			idx++
		}
	}
	return idx
}

// RemoveDuplicatesI - https://leetcode.com/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func RemoveDuplicatesI(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	idx := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}

// MajorityElement - https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
func MajorityElement(nums []int) int {
	candidate, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// Rotate - https://leetcode.com/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
func Rotate(nums []int, k int) {
	n := len(nums)
	if n < 2 {
		return
	}

	d := k % n
	if d == 0 {
		return
	}

	rev := func(arr []int) {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	rev(nums)
	rev(nums[:d])
	rev(nums[d:])
}

// MaxProfit - https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
func MaxProfit(prices []int) int {
	currMin, maxProfit := prices[0], 0

	for _, p := range prices[1:] {
		if p < currMin {
			currMin = p
		} else {
			maxProfit = max(maxProfit, p-currMin)
		}
	}
	return maxProfit
}

// MaxProfitII - https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/?envType=study-plan-v2&envId=top-interview-150
func MaxProfitII(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

// CanJump - https://leetcode.com/problems/jump-game/?envType=study-plan-v2&envId=top-interview-150
func CanJump(nums []int) bool {
	n := len(nums)

	maxReach := 0
	for i := 0; i <= maxReach && i < n; i++ {
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= n-1 {
			return true
		}
	}

	return false
}

// Jump - https://leetcode.com/problems/jump-game-ii/?envType=study-plan-v2&envId=top-interview-150
func Jump(nums []int) int {
	n := len(nums)

	var jumps, currEnd, currFarthest int
	for i := range n - 1 {
		currFarthest = max(currFarthest, i+nums[i])

		if i == currEnd {
			jumps++
			currEnd = currFarthest
			if currEnd >= n-1 {
				return jumps
			}
		}
	}
	return jumps
}
