package main

import "fmt"

func main() {
	var array = [5][2]int{}
	array[0][0] = 10
	array[1][0] = 12
	array[2][0] = 15
	array[3][0] = 19
	array[4][0] = 24

	for i := 0; i < len(array); i++ {
		array[i][1] = array[i][0] * 2	
	}
	fmt.Println(array)
}