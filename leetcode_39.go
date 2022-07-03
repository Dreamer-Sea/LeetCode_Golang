package main

import (
	"sort"
)

func main() {

}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	var backtracking func(temp []int, start, sum int)
	backtracking = func(temp []int, start, sum int) {
		if sum == target {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
		}
		for i := start; i < len(candidates) && sum+candidates[i] <= target; i++ {
			temp = append(temp, candidates[i])
			backtracking(temp, i, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	backtracking([]int{}, 0, 0)
	return res
}
