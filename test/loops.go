package main

import "fmt"

func main() {
	fmt.Println(sumN(5))

	fmt.Println("-----------------------------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-----------------------------------")
	for j := 0; j <= 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

}

func sumN(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}
