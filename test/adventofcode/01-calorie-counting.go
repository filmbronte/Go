package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	dat, err := os.ReadFile("test")

	if err != nil {
		panic(err)
	}

	calories := string(dat)
	// fmt.Printf("type of a is %T\n", calories)
	caloriesPerElf := strings.Split(calories, "\n\n")

	// result := []

	for _, row := range caloriesPerElf {
			fmt.Println("--------")
			fmt.Println(row)
			// for _, el := range row {
			// 	fmt.Println("--------")
			// 	fmt.Println(el)
			// }
	}

}