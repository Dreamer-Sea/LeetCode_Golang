package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "101023"
	res := restoreIpAddresses3(s)
	fmt.Println(res)
}

func restoreIpAddresses3(s string) []string {
	res := make([]string, 0)
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if isValid3(s, i, j) {
				dp[i][j] = true
			}
		}
	}
	dfs6(s, []string{}, 0, &res, dp)
	return res
}

func dfs6(s string, temp []string, start int, res *[]string, dp [][]bool) {
	if len(s) == start && len(temp) == 4 {
		tmpString := temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]
		*res = append(*res, tmpString)
		return
	}
	for i := start; i < len(s); i++ {
		if !dp[start][i] {
			continue
		}
		temp = append(temp, s[start:i+1])
		if i-start+1 <= 3 && len(temp) <= 4 && dp[start][i] {
			dfs6(s, temp, i+1, res, dp)
		}
		temp = temp[:len(temp)-1]
	}
}

func isValid3(s string, start, end int) bool {
	num, _ := strconv.Atoi(s[start : end+1])
	if (end-start+1 > 1 && s[start] == '0') || num > 255 {
		return false
	}
	return true
}
