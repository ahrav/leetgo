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
