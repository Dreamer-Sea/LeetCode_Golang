package main

func main() {

}

func letterCombinations(digits string) []string {
	if len(digits) == 0 || len(digits) > 4 ||
		digits == "" || digits == "1" || digits == "0" ||
		digits == "*" || digits == "#" {
		return []string{}
	}
	res := make([]string, 0)
	digitsMap := [10]string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	var backtracking func(temp string, idx int)
	backtracking = func(temp string, idx int) {
		if idx == len(digits) {
			res = append(res, temp)
			return
		}
		num := digits[idx] - '0'
		chars := digitsMap[num]
		for i := 0; i < len(chars); i++ {
			temp += string(chars[i])
			backtracking(temp, idx+1)
			temp = temp[:len(temp)-1]
		}
	}
	backtracking("", 0)
	return res
}
