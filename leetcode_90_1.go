package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	sort.Ints(nums)
	res := subsetsWithDup(nums)
	fmt.Println(res)
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	dfs7(nums, []int{}, 0, &res)
	return res
}

func dfs7(nums, temp []int, start int, res *[][]int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		dfs7(nums, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}
