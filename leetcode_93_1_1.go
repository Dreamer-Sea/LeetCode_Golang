package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "101023"
	res := restoreIpAddresses11(s)
	fmt.Println(res)
}

func restoreIpAddresses11(s string) []string {
	res := make([]string, 0)
	backtracking15(s, []string{}, 0, &res)
	return res
}

func backtracking15(s string, temp []string, start int, res *[]string) {
	if start == len(s) && len(temp) == 4 {
		tmpStr := temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]
		*res = append(*res, tmpStr)
		return
	}
	for i := start; i < len(s); i++ {
		if i-start+1 >= 4 || len(temp) >= 4 || !isValid6(s, start, i) {
			continue
		}
		temp = append(temp, s[start:i+1])
		backtracking15(s, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}

func isValid6(s string, start, end int) bool {
	num, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if num > 255 {
		return false
	}
	return true
}
