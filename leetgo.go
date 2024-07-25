package leetgo

import (
	"math"
	"sort"
	"strings"
)

func MaxArea(heights []int) int {
	var total int

	left, right := 0, len(heights)-1
	for left < right {
		hl, hr := heights[left], heights[right]
		h := min(hl, hr)
		area := h * (right - left)

		if area > total {
			area = total
		}

		if hl < hr {
			left++
		} else {
			right--
		}
	}

	return total
}

func IntToRoman(number int) string {
	romanNumerals := []struct {
		symbol string
		value  int
	}{
		{symbol: "M", value: 1000},
		{symbol: "CM", value: 900},
		{symbol: "D", value: 500},
		{symbol: "CD", value: 400},
		{symbol: "C", value: 100},
		{symbol: "XC", value: 90},
		{symbol: "L", value: 50},
		{symbol: "XL", value: 40},
		{symbol: "X", value: 10},
		{symbol: "IX", value: 9},
		{symbol: "V", value: 5},
		{symbol: "IV", value: 4},
		{symbol: "I", value: 1},
	}

	var res strings.Builder
	for _, rn := range romanNumerals {
		for number >= rn.value {
			res.WriteString(rn.symbol)
			number -= rn.value
		}
	}

	return res.String()
}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for idx, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, idx}
		}
		m[num] = idx
	}

	return nil
}

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

	return
}

func MinEatingSpeed(piles []int, h int) int {
	canFinish := func(rate int) bool {
		var hours int
		for _, pile := range piles {
			hours += (pile + rate - 1) / rate
		}
		return hours <= h
	}

	minV, maxV := 1, 0
	for _, pile := range piles {
		if pile > maxV {
			maxV = pile
		}
	}

	for minV < maxV {
		mid := minV + (maxV-minV)/2
		if canFinish(mid) {
			maxV = mid
		} else {
			minV = mid + 1
		}
	}

	return minV
}

func CountGroups(related []string) int {
	visited := make([]bool, len(related))

	var dfs func(int)
	dfs = func(user int) {
		visited[user] = true
		for idx := range related {
			if related[user][idx] == '1' && !visited[idx] {
				dfs(idx)
			}
		}
	}

	var groups int

	for idx := range related {
		if !visited[idx] {
			dfs(idx)
			groups++
		}
	}

	return groups
}

func RomanToInteger(s string) int {
	numerals := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	var total, prev int

	for i := len(s) - 1; i >= 0; i-- {
		curr, _ := numerals[s[i]]
		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}

	return total
}

func FindMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	len1, len2 := len(nums1), len(nums2)
	left, right := 0, len1

	for left <= right {
		partX := (left + right + 1) / 2
		partY := ((len1 + len2 + 1) / 2) - partX

		var maxPartLeftX, minPartRightX int
		if partX == 0 {
			maxPartLeftX = math.MinInt64
		} else {
			maxPartLeftX = nums1[partX-1]
		}

		if partX == len1 {
			minPartRightX = math.MaxInt64
		} else {
			minPartRightX = nums1[partX]
		}

		var maxPartLeftY, minPartRightY int
		if partY == 0 {
			maxPartLeftY = math.MinInt64
		} else {
			maxPartLeftY = nums2[partY-1]
		}

		if partY == len2 {
			minPartRightY = math.MaxInt64
		} else {
			minPartRightY = nums2[partY]
		}

		if maxPartLeftX <= minPartRightX && maxPartLeftX <= minPartRightY {
			if (len1+len2)%2 == 0 {
				return (float64(max(maxPartLeftX, maxPartLeftY)) + float64(min(minPartRightX, minPartRightY))) / float64(2)
			}
			return float64(max(maxPartLeftY, maxPartLeftX))
		}

		if maxPartLeftX > maxPartLeftY {
			right = partX - 1
		} else {
			left = partX + 1
		}
	}

	return 0.0
}

func MinimumDifference(nums []int) int {
	const threshold = 4
	if len(nums) <= threshold {
		return 0
	}

	top, bottom := make([]int, threshold), make([]int, threshold)
	copy(top, nums[:threshold])
	copy(bottom, nums[:threshold])

	sort.Ints(top)
	sort.Ints(bottom)

	for _, num := range nums[threshold:] {
		if num < bottom[3] {
			bottom[3] = num
			for i := 3; i > 0; i-- {
				if bottom[i] < bottom[i-1] {
					bottom[i], bottom[i-1] = bottom[i-1], bottom[i]
				}
			}
		}

		if num > top[0] {
			top[0] = num
			for i := 0; i < 3; i++ {
				if top[i] > top[i+1] {
					top[i], top[i+1] = top[i+1], top[i]
				}
			}
		}
	}

	return min(
		top[0]-bottom[0],
		top[1]-bottom[1],
		top[2]-bottom[2],
		top[3]-bottom[3],
	)
}
