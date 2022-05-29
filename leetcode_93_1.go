package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "25525511135"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	dfs4(s, []string{}, 0, &res)
	return res
}

func dfs4(s string, temp []string, start int, res *[]string) {
	if start == len(s) && len(temp) == 4 {
		tmpString := temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]
		*res = append(*res, tmpString)
		return
	}
	for i := start; i < len(s); i++ {
		temp = append(temp, s[start:i+1])
		if i-start+1 <= 3 && len(temp) <= 4 && isValid(s, start, i) {
			dfs4(s, temp, i+1, res)
		} else {
			return
		}
		temp = temp[:len(temp)-1]
	}
}

func isValid(s string, start, end int) bool {
	num, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	if num > 255 {
		return false
	}
	return true
}
