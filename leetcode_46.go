package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	dfs9(nums, []int{}, used, &res)
	return res
}

func dfs9(nums, temp []int, used []bool, res *[][]int) {
	if len(temp) == len(nums) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		temp = append(temp, nums[i])
		dfs9(nums, temp, used, res)
		temp = temp[:len(temp)-1]
		used[i] = false
	}
}
