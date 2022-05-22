package main

import (
	"fmt"
)

func main() {
	subsets2([]int{1, 2, 3})
	fmt.Println(res2)
}

var res2 [][]int

func subsets2(nums []int) [][]int {
	res2 = make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		dfs(nums, []int{}, i, 0)
	}
	return res2
}

func dfs(nums, temp []int, length, start int) {
	if len(temp) == length {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res2 = append(res2, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		dfs(nums, temp, length, i+1)
		temp = temp[:len(temp)-1]
	}
}
