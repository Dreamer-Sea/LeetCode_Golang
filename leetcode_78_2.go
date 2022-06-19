package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := subsets1(nums)
	fmt.Println(res)
}

func subsets1(nums []int) [][]int {
	res := make([][]int, 0)
	var backtracking func(int, []int)
	backtracking = func(start int, temp []int) {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			temp = append(temp, nums[i])
			backtracking(i+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking(0, []int{})
	return res
}
