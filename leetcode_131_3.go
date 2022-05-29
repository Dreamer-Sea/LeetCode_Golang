package main

import (
	"fmt"
)

func main() {
	s := "aab"
	res := partition3(s)
	fmt.Println(res)
}

func partition3(s string) [][]string {
	res := make([][]string, 0)
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
			} else if j-i == 1 && s[j] == s[i] {
				dp[i][j] = true
			} else if j-i > 1 && s[j] == s[i] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}
	dfs3(s, []string{}, 0, &res, dp)
	return res
}

func dfs3(s string, temp []string, start int, res *[][]string, dp [][]bool) {
	if start == len(s) {
		tmp := make([]string, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if !dp[start][i] {
			continue
		}
		if dp[start][i] {
			temp = append(temp, s[start:i+1])
			dfs3(s, temp, i+1, res, dp)
			temp = temp[:len(temp)-1]
		}
	}
}
