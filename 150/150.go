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

// RemoveDuplicates - https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
func RemoveDuplicates(nums []int) int {
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
