package main

import (
	"fmt"
)

func main() {
	s := "aab"
	partition(s)
	fmt.Println(res)
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	backtracking9(s, []string{}, 0, &res)
	return res
}

func backtracking9(s string, temp []string, start int, res *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		if isPalindrome(s, start, i) {
			temp = append(temp, s[start:i+1])
		} else {
			continue
		}
		backtracking9(s, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}

func isPalindrome(s string, start, end int) bool {
	if len(s) == 0 {
		return false
	}
	i, j := start, end
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
