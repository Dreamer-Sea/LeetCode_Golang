package main

func main() {

}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtracking func(temp []int, start int)
	backtracking = func(temp []int, start int) {
		if len(temp) == k {
			tmp := make([]int, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			temp = append(temp, i)
			backtracking(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking([]int{}, 1)
	return res
}
