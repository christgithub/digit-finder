package main

import "fmt"

const (
	inc = 10
	num int = 599009509
	find int = 5
)
var arr []int

func main() {
	displayResult(findNum(1, 0, 0))
}

func findNum(i, res, res2 int) []int {
	res = num/i; i*=inc
	res2 = res%inc
	if res2 == find {
		arr = append(arr, res2)
	}
	if i >= num {
		return arr
	}
	return findNum(i, res, res2)
}

func displayResult (found []int) {
	for _, n := range found {
		fmt.Print(n, " ")
	}
}