package dp

import (
	"sort"
	"strconv"
)

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

// LargestDivisibleSubset - https://leetcode.com/problems/largest-divisible-subset/
func LargestDivisibleSubset(nums []int) []int {
	n := len(nums)

	if n < 2 {
		return nums
	}

	sort.Ints(nums)

	dp, parents := make([]int, n), make([]int, n)
	for i := range dp {
		dp[i] = 1
		parents[i] = -1
	}

	maxIdx := 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				parents[i] = j
			}

			if dp[i] > dp[maxIdx] {
				maxIdx = i
			}
		}
	}

	var results []int
	for maxIdx != -1 {
		results = append(results, nums[maxIdx])
		maxIdx = parents[maxIdx]
	}

	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	return results
}

// NumDecodings - https://leetcode.com/problems/decode-ways/
func NumDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		digit, _ := strconv.Atoi(string(s[i-1]))
		if digit != 0 {
			dp[i] += dp[i-1]
		}

		digit, _ = strconv.Atoi(string(s[i-2 : i]))
		if digit >= 10 && digit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// JobScheduling - https://leetcode.com/problems/maximum-profit-in-job-scheduling/
func JobScheduling(startTime []int, endTime []int, profit []int) int {
	type Job struct {
		start  int
		end    int
		profit int
	}

	n := len(profit)
	jobs := make([]Job, n)
	for i := range profit {
		jobs[i] = Job{start: startTime[i], end: endTime[i], profit: profit[i]}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].end < jobs[j].end
	})

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		job := jobs[i-1]

		numJobs := sort.Search(n, func(i int) bool {
			return jobs[i].end > job.start
		})

		dp[i] = max(dp[i-1], job.profit+dp[numJobs])
	}

	return dp[n]
}
