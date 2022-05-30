package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 4, 3, 2, 1}
	res := findSubsequences(nums)
	fmt.Println(res)
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	dfs8(nums, []int{}, 0, &res)
	return res
}

func dfs8(nums, temp []int, start int, res *[][]int) {
	if len(temp) >= 2 {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
	}
	used := make([]int, 201)
	for i := start; i < len(nums); i++ {
		if (len(temp) > 0 && nums[i] < temp[len(temp)-1]) || used[nums[i]+100] == 1 {
			continue
		}
		used[nums[i]+100] = 1
		temp = append(temp, nums[i])
		dfs8(nums, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}
