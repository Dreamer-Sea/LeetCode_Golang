package main

import (
	"fmt"
)

func main() {
	combine(4, 2)
	fmt.Println(res4)
}

var res4 [][]int

func combine(n int, k int) [][]int {
	res4 = make([][]int, 0)
	backtracking4([]int{}, n, k, 1)
	return res4
}

func backtracking4(temp []int, n, k, start int) {
	if len(temp) == k {
		tmp := make([]int, len(temp))
		copy(tmp, temp)
		res4 = append(res4, tmp)
		return
	}
	for i := start; i <= n; i++ {
		temp = append(temp, i)
		backtracking4(temp, n, k, i+1)
		temp = temp[:len(temp)-1]
	}
}
