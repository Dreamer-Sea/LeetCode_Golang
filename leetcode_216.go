package main

import (
	"fmt"
)

func main() {
	combinationSum3(9, 45)
	fmt.Println(res5)
}

var res5 [][]int

func combinationSum3(k int, n int) [][]int {
	res5 = make([][]int, 0)
	backtracking5([]int{}, k, n, 0, 1)
	return res5
}

func backtracking5(temp []int, k, n, sum, start int) {
	if sum > n {
		return
	}
	if sum == n {
		if len(temp) != k {
			return
		}
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res5 = append(res5, tmp)
		return
	}
	for i := start; i <= 9; i++ {
		temp = append(temp, i)
		backtracking5(temp, k, n, sum+i, i+1)
		temp = temp[:len(temp)-1]
	}
}
