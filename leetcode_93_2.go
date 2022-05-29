package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "101023"
	res := restoreIpAddresses2(s)
	fmt.Println(res)
}

func restoreIpAddresses2(s string) []string {
	res := make([]string, 0)
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	dfs5(s, []string{}, 0, &res, dp)
	return res
}

func dfs5(s string, temp []string, start int, res *[]string, dp [][]int) {
	if len(s) == start && len(temp) == 4 {
		tmpString := temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]
		*res = append(*res, tmpString)
		return
	}
	for i := start; i < len(s); i++ {
		if dp[start][i] == 2 {
			continue
		}
		temp = append(temp, s[start:i+1])
		if i-start+1 <= 3 && len(temp) <= 4 && (dp[start][i] == 1 || isValid2(s, start, i, dp)) {
			dfs5(s, temp, i+1, res, dp)
		}
		temp = temp[:len(temp)-1]
	}
}

func isValid2(s string, start, end int, dp [][]int) bool {
	num, _ := strconv.Atoi(s[start : end+1])
	if (end-start+1 > 1 && s[start] == '0') || num > 255 {
		dp[start][end] = 2
		return false
	}
	dp[start][end] = 1
	return true
}
