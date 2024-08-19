package recursionI

func GetRow(index int) []int {
	if index == 0 {
		return []int{1}
	}

	prevRow := GetRow(index - 1)

	currRow := make([]int, index+1)
	currRow[0], currRow[index] = 1, 1

	for i := 1; i < index; i++ {
		currRow[i] = prevRow[i-1] + prevRow[i]
	}

	return currRow
}

func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}

	half := MyPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func KthGrammar(n, k int) int {
	if n == 1 {
		return 0
	}

	length := 1 << (n - 1)
	mid := length / 2

	if k <= mid {
		return KthGrammar(n-1, k)
	} else {
		return 1 - KthGrammar(n-1, k-mid)
	}
}
