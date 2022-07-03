package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
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
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			backtracking(temp, i+1, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	backtracking([]int{}, 0, 0)
	return res
}
