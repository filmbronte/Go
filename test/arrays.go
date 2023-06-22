package main

import (
	"fmt"
	"math"
)

func main() {
	// could assign length
	var myArray = [3]int{1, 2, 3}

	// without assigning length
	otherArray := [...]int{4, 5, 6, 7, 8}

	fmt.Println(myArray)
	fmt.Println(otherArray)

	otherArray[0] = 100
	fmt.Println(otherArray)

	fmt.Println(otherArray[3])
	fmt.Println(otherArray[len(otherArray)-1])  //getting last index
	fmt.Println(otherArray[:len(otherArray)-1]) //removing last index (does not affect original array)
	fmt.Println(otherArray)
	fmt.Println(len(otherArray))

	emptyArray := [...]int{}

	fmt.Println(emptyArray)
	fmt.Println(cap(emptyArray)) // 0

	zeroArray := [5]int{}

	fmt.Println(zeroArray)
	fmt.Println(cap(zeroArray)) // 5

	arrayToSlice := [6]int{1, 2, 3, 4, 5, 6}
	slicedArray := arrayToSlice[2:4]
	// slicedArray := arrayToSlice
	// slicedArray[0] = 100

	fmt.Println(arrayToSlice)
	fmt.Println(slicedArray)
	fmt.Println(cap(slicedArray)) // 4
	fmt.Println(len(slicedArray)) // 2

	mySlice1 := make([]int, 5, 10)

	fmt.Println(mySlice1)
	fmt.Println(cap(mySlice1))
	fmt.Println(len(mySlice1))
	newSlice := append(mySlice1, 0)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println(mySlice1)

	//str := "strings are easy"

	//fmt.Println(len(str))
	//fmt.Println(str)

	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}

	newArr := append(arr1, arr2...)

	fmt.Println(newArr) // slice!

	fmt.Println(math.Pi) //math methods
}
