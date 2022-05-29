package main

import (
	"fmt"
)

func main() {
	s := "aab"
	res := partition(s)
	fmt.Println(res)
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	dfs1(s, []string{}, 0, &res)
	return res
}

func dfs1(s string, temp []string, start int, res *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome1(s, start, i) {
			temp = append(temp, s[start:i+1])
		} else {
			continue
		}
		dfs1(s, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}

func isPalindrome1(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
