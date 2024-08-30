package prefixsum

// VowelStrings - https://leetcode.com/problems/count-vowel-strings-in-ranges/
func VowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	if n == 0 {
		return nil
	}

	isVowel := func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
	}

	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		if isVowel(words[i][0]) && isVowel(words[i][len(words[i])-1]) {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}

	result := make([]int, len(queries))
	for idx, query := range queries {
		result[idx] = prefix[query[1]+1] - prefix[query[0]]
	}

	return result
}
