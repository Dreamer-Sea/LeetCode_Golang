package main

func main() {

}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	var backtracking func(temp []int, start, sum int)
	backtracking = func(temp []int, start, sum int) {
		if sum == n {
			if len(temp) != k {
				return
			}
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
			return
		}
		for i := start; i <= 9 && i+sum <= n; i++ {
			temp = append(temp, i)
			backtracking(temp, i+1, sum+i)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking([]int{}, 1, 0)
	return res
}
