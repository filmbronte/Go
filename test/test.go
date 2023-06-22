package main

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}

// func main() {
// fmt.Println(Hello("Vic"))
// }

func printNums(a, b int, c int) string {
	nums := fmt.Sprintf("%v, %v, %v", a, b, c)
	return nums
}

func main() {
	var (
		a    int
		b    int     = 10
		c    float32 = 15.5
		name string  = "Vic"
	)

	fmt.Println(printNums(3, 2, 1))
	fmt.Println(Hello("Vic"))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(name)
}

// func main() {
// 	const name = "Vic"

// 	// name = "John"

// 	fmt.Println(name)
// }
