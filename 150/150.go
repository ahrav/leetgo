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
