package tiktok

import (
	"fmt"
	"math"
)

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

// MaxProduct - https://leetcode.com/problems/maximum-product-subarray/?envType=study-plan-v2&envId=tiktok-spring-23-high-frequency
func MaxProduct(nums []int) int {
	n := len(nums)
	minProd, maxProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {
		if nums[i] < 0 {
			minProd, maxProd = maxProd, minProd
		}

		minProd = min(nums[i], minProd*nums[i])
		maxProd = max(nums[i], maxProd*nums[i])

		result = max(result, maxProd)
	}

	return result
}

// FindNthDigit - https://leetcode.com/problems/nth-digit/?envType=study-plan-v2&envId=tiktok-spring-23-high-frequency
func FindNthDigit(n int) int {
	length := 1
	count := 9
	start := 1

	for n > length*count {
		n -= length * count
		length++
		count *= 10
		start *= 10
	}

	num := start + (n-1)/length
	digitIdx := (n - 1) % length

	numStr := fmt.Sprintf("%d", num)
	return int(numStr[digitIdx] - '0')
}
