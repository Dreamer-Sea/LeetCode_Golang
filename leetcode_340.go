package main

import (
	"fmt"
)

func LengthOfLongestSubstringWithKDistinct(s string, k int) int {
	if s == "" || k == 0 {
		return 0
	}
	left, right, maxLen := 0, 0, 0
	charMap := make(map[string]interface{}, k)
	for right < len(s) {
		charMap[string(s[right])] = struct{}{}
		if len(charMap) > k {
			left++
			delete(charMap, string(s[left]))
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		right++
	}
	return maxLen
}

func main() {
	s, k := "eceba", 2
	res := LengthOfLongestSubstringWithKDistinct(s, k)
	fmt.Println(res)
}
