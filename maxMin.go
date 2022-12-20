package main
import "fmt"

func findMinAndLeastMin(arr []int, size int){
	var secondMinNum int
	firstMinNum := arr[0]

	for i := 0; i < size; i++{
		if firstMinNum > arr[i]{
			secondMinNum = firstMinNum 
			firstMinNum = arr[i] 
			
		   } else if secondMinNum > arr[i] {
			   secondMinNum = arr[i]
		   }
	   }
	fmt.Println("**********************")
	fmt.Println("Minimum element : ", secondMinNum)
	fmt.Println("Least Minimum element :", firstMinNum)
}
func findMaxAndHighestMax(arr []int, size int){
	maxNum := arr[0]
	var highestMaxNum int

	for i := 0; i < size; i++{
		if highestMaxNum < arr[i]{
			maxNum = highestMaxNum 
			highestMaxNum = arr[i] 
		} else if maxNum < arr[i] {
			maxNum = arr[i]
		} 	
	}
	fmt.Println("**********************")
	fmt.Println("Maximum element : ", maxNum)
	fmt.Println("Highest Maximum element :", highestMaxNum)
}
func main(){
	var size int
	fmt.Println("Enter size of array : ")
	fmt.Scanln(&size)
	arr := make([]int, size)
		
	for i := 0; i < size; i++{
		fmt.Println("Enter the element :")
		fmt.Scanf("%d", &arr[i])
	}

	findMaxAndHighestMax(arr, size)
	findMinAndLeastMin(arr, size)

}