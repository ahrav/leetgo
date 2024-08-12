package seventyfive

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func MergeAlternately(word1 string, word2 string) string {
	len1, len2 := len(word1), len(word2)

	var sb strings.Builder
	sb.Grow(len1 + len2)

	maxLen := len1
	if len2 > len1 {
		maxLen = len2
	}

	for i := 0; i < maxLen; i++ {
		if i < len1 {
			sb.WriteByte(word1[i])
		}

		if i < len2 {
			sb.WriteByte(word2[i])
		}
	}

	return sb.String()
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	return str1[:gcd(len(str1), len(str2))]
}

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxV := candies[0]

	n := len(candies)
	for i := 1; i < n; i++ {
		if candies[i] > maxV {
			maxV = candies[i]
		}
	}

	res := make([]bool, n)
	for i := 0; i < n; i++ {
		res[i] = candies[i]+extraCandies >= maxV
	}

	return res
}

func CanPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if n == 0 || size == 0 {
		return true
	}

	for i := 0; i < size; i++ {
		if flowerbed[i] == 0 {
			leftEmpty := (i == 0) || flowerbed[i-1] == 0
			rightEmpty := (i == size-1) || flowerbed[i+1] == 0

			if leftEmpty && rightEmpty {
				flowerbed[i] = 1
				i++
				n--
			}

			if n == 0 {
				return true
			}
		}
	}

	return false
}

func ReverseVowels(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	vowels := [128]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	res := []byte(s)
	lp, rp := 0, n-1
	for lp < rp {
		for lp < rp && !vowels[res[lp]] {
			lp++
		}
		for lp < rp && !vowels[res[rp]] {
			rp--
		}

		if lp < rp {
			res[lp], res[rp] = res[rp], res[lp]
			lp++
			rp--
		}
	}

	return string(res)
}

func ReverseWords(s string) string {
	chars := []byte(s)

	reverse := func(start, end int) {
		for start < end {
			chars[start], chars[end] = chars[end], chars[start]
			start++
			end--
		}
	}

	n := len(chars)
	reverse(0, n-1)

	start, writeIdx := 0, 0
	for start < n {
		if chars[start] == ' ' {
			start++
			continue
		}

		if writeIdx != 0 {
			chars[writeIdx] = ' '
			writeIdx++
		}

		end := start
		for end < n && chars[end] != ' ' {
			chars[writeIdx] = chars[end]
			end++
			writeIdx++
		}

		reverse(writeIdx-(end-start), writeIdx-1)
		start = end
	}

	return string(chars[:writeIdx])
}

func MoveZeroes(nums []int) {
	n := len(nums)
	var insertIdx int
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[insertIdx], nums[i] = nums[i], nums[insertIdx]
			insertIdx++
		}
	}
}

func IsSubsequence(s string, t string) bool {
	slen, tlen := len(s), len(t)
	if slen > tlen {
		return false
	}

	sp, tp := 0, 0
	for sp < slen && tp < tlen {
		if s[sp] == t[tp] {
			sp++
			tp++
		} else {
			tp++
		}
	}

	return sp == slen
}

func FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if k > n {
		return 0
	}

	var sum int
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < n; i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func MaxArea(height []int) int {
	n := len(height)

	lp, rp, maxArea := 0, n-1, 0
	for lp < rp {
		width := rp - lp
		minH := 0

		if height[lp] > height[rp] {
			minH = height[rp]
			rp--
		} else {
			minH = height[lp]
			lp++
		}

		area := width * minH
		if width*minH > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func MaxOperationsTwoPointer(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == k {
			return 1
		}
		return 0
	}

	sort.Ints(nums)

	var operations int
	lp, rp := 0, n-1
	for lp < rp {
		sum := nums[lp] + nums[rp]
		if sum == k {
			operations++
			lp++
			rp--
			continue
		}

		if sum > k {
			rp--
		} else {
			lp++
		}
	}

	return operations
}

func MaxOperationsComplement(nums []int, k int) int {
	freq := make(map[int]int, len(nums)/2)
	var operations int

	for _, num := range nums {
		complement := k - num
		if count, exists := freq[complement]; exists && count > 0 {
			operations++
			freq[complement]--
		} else {
			freq[num]++
		}
	}

	return operations
}

func IncreasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, second := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}

func IncreasingTripletDP(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		if dp[i] >= 3 {
			return true
		}
	}

	return false
}

func Compress(chars []byte) int {
	n := len(chars)
	if n < 2 {
		return n
	}

	var rPtr, wPtr int
	for rPtr < n {
		curr := chars[rPtr]
		cnt := 0

		for rPtr < n && chars[rPtr] == curr {
			rPtr++
			cnt++
		}

		chars[wPtr] = curr
		wPtr++

		if cnt > 1 {
			str := strconv.Itoa(cnt)
			for i := 0; i < len(str); i++ {
				chars[wPtr] = str[i]
				wPtr++
			}
		}
	}

	return wPtr
}

func LongestOnes(nums []int, k int) int {
	var left, zeroCount, maxWindow int
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		if window := right - left + 1; window > maxWindow {
			maxWindow = window
		}
	}

	return maxWindow
}

func FindDifference(nums1 []int, nums2 []int) [][]int {
	set := make(map[int16]uint8)
	for _, num := range nums1 {
		set[int16(num)] |= 1
	}
	for _, num := range nums2 {
		set[int16(num)] |= 2
	}

	res := make([][]int, 2)
	for k, v := range set {
		switch v {
		case 1:
			res[0] = append(res[0], int(k))
		case 2:
			res[1] = append(res[1], int(k))
		}
	}

	return res
}

func MaxVowels(s string, k int) int {
	if len(s) < 1 {
		return 0
	}

	vowels := [26]bool{0: true, 'e' - 'a': true, 'i' - 'a': true, 'o' - 'a': true, 'u' - 'a': true}
	var count int
	for i := range k {
		if vowels[s[i]-'a'] {
			count++
		}
	}

	maxVowels := count
	for i := k; i < len(s); i++ {
		if vowels[s[i-k]-'a'] {
			count--
		}
		if vowels[s[i]-'a'] {
			count++
		}
		if count > maxVowels {
			maxVowels = count
		}
	}

	return maxVowels
}
