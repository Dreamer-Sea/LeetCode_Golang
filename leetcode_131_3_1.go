package main

func main() {

}

func partition31(s string) [][]string {
	res := make([][]string, 0)
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if !isPalindrome31(s, i, j) {
				continue
			}
			dp[i][j] = true
		}
	}
	var backtracking func(start int, temp []string, dp [][]bool)
	backtracking = func(start int, temp []string, dp [][]bool) {
		if start == len(s) {
			tmp := make([]string, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
		}
		for i := start; i < len(s); i++ {
			if !dp[start][i] {
				continue
			}
			temp = append(temp, s[start:i+1])
			backtracking(i+1, temp, dp)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking(0, []string{}, dp)
	return res
}

func isPalindrome31(s string, start, end int) bool {
	left, right := start, end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
