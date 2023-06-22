package main

import "fmt"

func main() {
	fmt.Println(solve(2))
	fmt.Println(solve(5))
	fmt.Println(solve(20))
	fmt.Println(solve(10))

}

func solve(n int) int {
	return n * 2
}
