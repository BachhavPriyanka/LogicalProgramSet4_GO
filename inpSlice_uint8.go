package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := []int{}
	var input string
	var size int
	var numbers []int

	fmt.Println("Enter a size : ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		fmt.Println("Enter a number : ")
		fmt.Scanln(&input)
		if input == "X" {
			break
		}

		slicevar, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		slice = append(slice, slicevar)

		sort.Ints(slice)
		fmt.Println(slice)
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] != 0 {
			numbers = append(numbers, slice[i])
		}

	}
	fmt.Println("Sorted : ",numbers)

}