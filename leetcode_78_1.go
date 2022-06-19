package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	backtracking14(nums, []int{}, 0, &res)
	return res
}

func backtracking14(nums, temp []int, start int, res *[][]int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtracking14(nums, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}
