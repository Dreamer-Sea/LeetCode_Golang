package main

import (
	"fmt"
)

func main() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	var backtracking func([]int)
	backtracking = func(temp []int) {
		if len(temp) == len(nums) {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
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
