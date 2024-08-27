package dp

import "sort"

// LengthOfLIS - https://leetcode.com/problems/longest-increasing-subsequence/
func LengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	currMax := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > currMax {
					currMax = dp[i]
				}
			}
		}
	}

	return currMax
}

// LengthOfLISBinarySearch - https://leetcode.com/problems/longest-increasing-subsequence/
func LengthOfLISBinarySearch(nums []int) int {
	var tails []int
	for _, num := range nums {
		idx := sort.Search(len(tails), func(i int) bool {
			return tails[i] >= num
		})

		if len(tails) == idx {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}
