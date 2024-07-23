package leetgo

import (
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
