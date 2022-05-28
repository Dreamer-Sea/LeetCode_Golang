package main

import (
	"fmt"
)

func main() {
	s := "aab"
	res := partition2(s)
	fmt.Println(res)
}

func partition2(s string) [][]string {
	res := make([][]string, 0)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	dfs2(s, []string{}, 0, &res, dp)
	return res
}

func dfs2(s string, temp []string, start int, res *[][]string, dp [][]int) {
	if start == len(s) {
		tmp := make([]string, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if dp[start][i] == 2 {
			continue
		}
		if dp[start][i] == 1 || isPalindrome2(s, start, i, dp) {
			temp = append(temp, s[start:i+1])
			dfs2(s, temp, i+1, res, dp)
			temp = temp[:len(temp)-1]
		}
	}
}

func isPalindrome2(s string, start, end int, dp [][]int) bool {
	for start < end {
		if s[start] != s[end] {
			dp[start][end] = 2
			return false
		}
		start++
		end--
	}
	dp[start][end] = 1
	return true
}
