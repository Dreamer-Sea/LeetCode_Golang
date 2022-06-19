package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "101023"
	res := restoreIpAddresses12(s)
	fmt.Println(res)
}

func restoreIpAddresses12(s string) []string {
	res := make([]string, 0)
	var backtracking func([]string, int)
	backtracking = func(temp []string, start int) {
		if start == len(s) && len(temp) == 4 {
			tmpStr := temp[0] + "." + temp[1] + "." + temp[2] + "." + temp[3]
			res = append(res, tmpStr)
			return
		}
		for i := start; i < len(s); i++ {
			if i-start+1 >= 4 || len(temp) >= 4 || !isValid7(s, start, i) {
				continue
			}
			temp = append(temp, s[start:i+1])
			backtracking(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking([]string{}, 0)
	return res
}

func isValid7(s string, start, end int) bool {
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s[start : end+1])
	if num > 255 {
		return false
	}
	return true
}
