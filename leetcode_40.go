package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	backtracking8(candidates, []int{}, target, 0, 0, &res)
	return res
}

func backtracking8(candidates, temp []int, target, sum, start int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < len(candidates) && sum+candidates[i] <= target; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		temp = append(temp, candidates[i])
		backtracking8(candidates, temp, target, sum+candidates[i], i+1, res)
		temp = temp[:len(temp)-1]
	}
}
