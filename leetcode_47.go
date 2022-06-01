package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	dfs10(nums, []int{}, used, &res)
	return res
}

func dfs10(nums, temp []int, used []bool, res *[][]int) {
	if len(nums) == len(temp) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		temp = append(temp, nums[i])
		dfs10(nums, temp, used, res)
		temp = temp[:len(temp)-1]
		used[i] = false
	}
}
