package main

import "fmt"

func main() {
	fmt.Println(solve(5.50))
	fmt.Println(solve(4.35))
}

func solve(grade float32) string {

	if grade >= 5.50 {
		return "Excellent"
	} else {
		return "Not excellent"
	}
}