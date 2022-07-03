package main

func main() {

}

func partition11(s string) [][]string {
	res := make([][]string, 0)
	var backtracking func(start int, temp []string)
	backtracking = func(start int, temp []string) {
		if start == len(s) {
			tmp := make([]string, len(temp))
			copy(tmp, temp)
			res = append(res, tmp)
		}
		for i := start; i < len(s); i++ {
			if isPalindrome11(s, start, i) {
				temp = append(temp, s[start:i+1])
			} else {
				continue
			}
			backtracking(i+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking(0, []string{})
	return res
}

func isPalindrome11(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
