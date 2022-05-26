package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtracking7(candidates, []int{}, target, 0, 0, &res)
	return res
}

func backtracking7(candidates, temp []int, target, sum, start int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	if sum+candidates[start] > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		backtracking7(candidates, temp, target, sum+candidates[i], i, res)
		temp = temp[:len(temp)-1]
	}
}
