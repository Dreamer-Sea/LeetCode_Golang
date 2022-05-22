package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	subsets(nums)
	fmt.Println(res)
}

var res [][]int

func subsets(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	backtracking(nums, []int{}, 0)
	return res
}

func backtracking(nums, temp []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backtracking(nums, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}
