package top50

import "strings"

// NumberToWords - https://leetcode.com/problems/integer-to-english-words/
func NumberToWords(num int) string {
	// Handle zero edge case.
	if num == 0 {
		return "Zero"
	}

	// Setup arrays in order to quickly determine word for a place value. (ones, hundreds, etc..)
	var (
		ones      = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
		tens      = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
		thousands = []string{"", "Thousand", "Million", "Billion"}
	)

	var helper func(num int) string
	helper = func(num int) string {
		switch {
		case num == 0:
			return ""
		case num < 10:
			return ones[num] + " "
		case num < 20:
			return teens[num-10] + " "
		case num < 100:
			return tens[num/10] + " " + helper(num%10)
		default:
			return ones[num/100] + " Hundred " + helper(num%100)
		}
	}

	var result []string
	for i := 0; num > 0; i++ {
		rem := num % 1000
		if rem != 0 {
			result = append([]string{helper(rem) + thousands[i]}, result...)
		}
		num /= 1000
	}

	return strings.TrimSpace(strings.Join(result, " "))
}

// Trap - https://leetcode.com/problems/trapping-rain-water/
func Trap(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	leftMax, rightMax := arr[0], arr[n-1]
	left, right := 0, n-1

	water := 0
	for left < right {
		if leftMax < rightMax {
			left++

			if arr[left] < leftMax {
				water += leftMax - arr[left]
			} else {
				leftMax = arr[left]
			}
		} else {
			right--

			if arr[right] < rightMax {
				water += rightMax - arr[right]
			} else {
				rightMax = arr[right]
			}
		}
	}

	return water
}

// NumIslands - https://leetcode.com/problems/number-of-islands/
func NumIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] == '0' {
			return
		}

		grid[x][y] = '0'
		for _, dir := range directions {
			nr, nc := x+dir[0], y+dir[1]
			dfs(nr, nc)
		}
	}

	numIslands := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == '1' {
				dfs(i, j)
				numIslands++
			}
		}
	}

	return numIslands
}

// GroupAnagrams - https://leetcode.com/problems/group-anagrams/
func GroupAnagrams(strs []string) [][]string {
	freq := make(map[[26]int][]string)

	for _, word := range strs {
		var cnt [26]int
		for _, char := range word {
			cnt[char-'a']++
		}

		freq[cnt] = append(freq[cnt], word)
	}

	result := make([][]string, 0, len(freq))
	for _, words := range freq {
		result = append(result, words)
	}

	return result
}
