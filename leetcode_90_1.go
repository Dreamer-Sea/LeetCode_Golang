package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	var backtracking func(int, []int)
	backtracking = func(start int, temp []int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			backtracking(i+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking(0, []int{})
	return res
}
