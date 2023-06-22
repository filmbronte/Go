package main

import "fmt"

func main() {
	// 'John', 15, 5.54678
	fmt.Println(studentInformation("John", 15, 5.54678))
	// 'Steve', 16, 2.1426
	fmt.Println(studentInformation("Steve", 16, 2.1426))
	// 'Marry', 12, 6.00 
	fmt.Println(studentInformation("Marry", 12, 6.00))
}

func studentInformation(name string, age int, avgGrade float32) string {
	result := fmt.Sprintf("Name: %s, Age: %v, Grade: %.2f", name, age, avgGrade)
	return result
}
