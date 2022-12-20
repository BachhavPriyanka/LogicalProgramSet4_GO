package main 
import "fmt"

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	arr1 := append(arr[4:14])
	fmt.Println(arr1)

	isOdd(arr1)
	isEven(arr1)
}

func isOdd(array []int){ 
	var oddSlice = []int{}

	for i := 0; i < len(array); i++ {
		if(array[i] % 2 != 0){
			oddSlice = append(oddSlice, array[i])
		}
	}
	fmt.Println("New Odd Slice : ",oddSlice)

}

func isEven(array []int){
	var evenSlice int 
	
	for i := 0; i < len(array); i++ {
		if(array[i] % 2 == 0){
			oddSlice = append(evenSlice, array[i])
		}
	}
	fmt.Println("New Even Slice :", evenSlice)
}