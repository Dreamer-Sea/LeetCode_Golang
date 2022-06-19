package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup1(nums)
	fmt.Println(res)
}

func subsetsWithDup1(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	backtracking13(nums, []int{}, 0, &res)
	return res
}

func backtracking13(nums, temp []int, start int, res *[][]int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		temp = append(temp, nums[i])
		backtracking13(nums, temp, i+1, res)
		temp = temp[:len(temp)-1]
	}
}
