package tiktok

import "math"

// MinSubArrayLen - https://leetcode.com/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=tiktok-spring-23-high-frequency
func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	var left, currSum int
	minWindow := math.MaxInt

	for right := 0; right < n; right++ {
		currSum += nums[right]

		for currSum >= target {
			minWindow = min(minWindow, right-left+1)
			currSum -= nums[left]
			left++
		}
	}

	if minWindow == math.MaxInt {
		return 0
	}

	return minWindow
}
