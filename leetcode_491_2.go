package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 6, 7, 7}
	res := findSubsequences1(nums)
	fmt.Println(res)
}

func findSubsequences1(nums []int) [][]int {
	res := make([][]int, 0)
	backtracking12(nums, []int{}, 0, &res)
	return res
}

func backtracking12(nums, temp []int, start int, res *[][]int) {
	if len(temp) > 1 {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
	}
	used := make([]bool, 201)
	for i := start; i < len(nums); i++ {
		if used[nums[i]+100] || (len(temp) > 0 && nums[i] < temp[len(temp)-1]) {
			continue
		}
		used[nums[i]+100] = true
		temp = append(temp, nums[i])
		backtracking12(nums, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}
