package main

import (
	"fmt"
)

func main() {
	res := partition21("aab")
	fmt.Println(res)
}

func partition21(s string) [][]string {
	res := make([][]string, 0)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	var backtracking func(start int, temp []string, dp [][]int)
	backtracking = func(start int, temp []string, dp [][]int) {
		if start == len(s) {
			tmp := make([]string, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(s); i++ {
			if dp[start][i] == 2 ||
				(dp[start][i] == 0 && !isPalindrome21(s, start, i, dp)) {
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

func isPalindrome21(s string, start, end int, dp [][]int) bool {
	left, right := start, end
	for left < right {
		if s[left] != s[right] {
			dp[start][end] = 2
			return false
		}
		left++
		right--
	}
	dp[start][end] = 1
	return true
}
