package main

import "fmt"

func main() {
	fmt.Println("---------- If statements ----------")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("------- Normal For loop -----------")
	for j := 0; j <= 10; j++ {
		if j%2 != 0 {
			continue
		}
		fmt.Println(j)
	}

	fmt.Println("---------- While loop -------------")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	
	fmt.Println("------- Loop in a function --------")
	fmt.Println(sumN(5))
	
}

func sumN(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		sum += i
	}

	return sum
}
