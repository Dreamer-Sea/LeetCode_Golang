package main

import (
	"fmt"
	"sort"
)

func main() {
	res := permuteUnique([]int{1, 1, 2})
	fmt.Println(res)
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	used := make([]bool, len(nums), len(nums))
	var backtracking func(temp []int)
	backtracking = func(temp []int) {
		if len(temp) == len(nums) {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			backtracking(temp)
			temp = temp[:len(temp)-1]
			used[i] = false
		}
	}
	backtracking([]int{})
	return res
}
