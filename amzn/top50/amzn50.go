package top50

import (
	"container/heap"
	"sort"
	"strings"
)

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

// MinimumSwaps - https://leetcode.com/problems/minimum-adjacent-swaps-to-make-a-valid-array/
func MinimumSwaps(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	minIdx, maxIdx := 0, 0
	for i, num := range nums {
		if num < nums[minIdx] {
			minIdx = i
		}
		if num >= nums[maxIdx] {
			maxIdx = i
		}
	}

	totalSwaps := n - 1 - maxIdx + minIdx
	if minIdx < maxIdx {
		return totalSwaps
	}

	return totalSwaps - 1
}

// MinMeetingRooms - https://leetcode.com/problems/meeting-rooms-ii/
func MinMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	if n == 1 {
		return 1
	}

	starts, ends := make([]int, n), make([]int, n)
	for i := range n {
		starts[i] = intervals[i][0]
		ends[i] = intervals[i][1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	startPtr, endPtr := 0, 0
	rooms, maxRooms := 0, 1

	for startPtr < n {
		if starts[startPtr] < ends[endPtr] {
			rooms++
			startPtr++
		} else {
			rooms--
			endPtr++
		}

		if rooms > maxRooms {
			maxRooms = rooms
		}
	}

	return maxRooms
}

type CharFreq struct {
	freq int
	char byte
}

type MaxCharFreqHeap []CharFreq

func (h MaxCharFreqHeap) Len() int           { return len(h) }
func (h MaxCharFreqHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxCharFreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxCharFreqHeap) Push(x any) { *h = append(*h, x.(CharFreq)) }
func (h *MaxCharFreqHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// ReorganizeString - https://leetcode.com/problems/reorganize-string/
func ReorganizeString(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}

	half := (n + 1) / 2
	var charFreq [26]int
	for _, char := range s {
		charFreq[char-'a']++
		if charFreq[char-'a'] > half {
			return ""
		}
	}

	h := new(MaxCharFreqHeap)
	heap.Init(h)

	for idx, freq := range charFreq {
		if freq > 0 {
			heap.Push(h, CharFreq{char: byte(idx) + 'a', freq: freq})
		}
	}

	var builder strings.Builder
	var prev CharFreq

	for h.Len() > 0 {
		cf := heap.Pop(h).(CharFreq)
		builder.WriteByte(cf.char)

		if prev.freq > 0 {
			heap.Push(h, prev)
		}

		cf.freq--
		prev = cf
	}

	return builder.String()
}

// OrangesRotting - https://leetcode.com/problems/rotting-oranges/
func OrangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var queue [][2]int
	freshCnt := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 {
				freshCnt++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	if freshCnt == 0 {
		return 0
	}

	if len(queue) == 0 {
		return -1
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	minutes := -1
	for len(queue) > 0 {
		level := len(queue)

		for i := 0; i < level; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				nr, nc := curr[0]+dir[0], curr[1]+dir[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] == 0 {
					continue
				}
				if grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					freshCnt--
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
		minutes++
	}

	if freshCnt == 0 {
		return minutes
	}
	return -1
}

// MaxSlidingWindow - https://leetcode.com/problems/sliding-window-maximum/
func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return []int{}
	}

	if n == 1 || k == 1 {
		return nums
	}

	var deque []int
	result := make([]int, 0, n-k+1)

	for i, num := range nums {
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
