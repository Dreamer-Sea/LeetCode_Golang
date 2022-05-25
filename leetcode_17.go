package main

import (
	"fmt"
)

func main() {
	result := letterCombinations("23")
	fmt.Println(result)
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 || len(digits) > 4 ||
		digits == "" || digits == "1" || digits == "0" ||
		digits == "*" || digits == "#" {
		return []string{}
	}
	res6 := make([]string, 0)
	digitsMap := [10]string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	backtracking6(digits, "", 0, digitsMap, &res6)
	return res6
}

func backtracking6(digits, temp string, idx int, digitsMap [10]string, res *[]string) {
	if len(digits) == idx {
		*res = append(*res, temp)
		return
	}
	num := digits[idx] - '0'
	chars := digitsMap[num]
	for i := 0; i < len(chars); i++ {
		temp = temp + string(chars[i])
		backtracking6(digits, temp, idx+1, digitsMap, res)
		temp = temp[:len(temp)-1]
	}
}
