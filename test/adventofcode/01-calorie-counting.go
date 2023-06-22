package main

import (
	"fmt"
	"os"
)

func main() {

	dat, err := os.ReadFile("adventofcode/calorie-counting")
	check(err)
	//calories := parseInput(string(dat))
	fmt.Printf("type of a is %T\n", dat)
	fmt.Println(len(dat))

	//for i, _ := range calories {
	//	fmt.Println(i)
	//	//fmt.Printf("%v - %v\n", i, el)
	//}

	//fmt.Print(strings.Split(string(dat), "\n\n")[0])

	//for _, el := range strings.Split(string(dat), "\n\n") {
	//	fmt.Println(el)
	//}

	//caloriesTest := "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000\n"
	//testArray := strings.Split(caloriesTest, "\n\n")
	//
	//for _, el := range testArray {
	//	fmt.Println(el)
	//
	//}
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//func parseInput(input string) (ans [][]int) {
//	for _, group := range strings.Split(input, "\n\n") {
//		row := []int{}
//		for _, line := range strings.Split(group, "\n") {
//			row = append(
//				row,
//				cast.ToInt(line),
//			)
//		}
//		ans = append(ans, row)
//	}
//	return ans
//}
