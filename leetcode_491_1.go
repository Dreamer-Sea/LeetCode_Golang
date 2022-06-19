package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 6, 7, 7}
	res := findSubsequences(nums)
	fmt.Println(res)
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	var backtracking func(int, []int)
	backtracking = func(start int, temp []int) {
		if len(temp) >= 2 {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
		}
		used := make([]bool, 201)
		for i := start; i < len(nums); i++ {
			if used[nums[i]+100] || (len(temp) > 0 && nums[i] < temp[len(temp)-1]) {
				continue
			}
			used[nums[i]+100] = true
			temp = append(temp, nums[i])
			backtracking(i+1, temp)
			temp = temp[:len(temp)-1]

		}
	}
	backtracking(0, []int{})
	return res
}
