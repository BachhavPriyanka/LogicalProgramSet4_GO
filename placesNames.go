package main
import "fmt"


func main(){
	
	var arrPlaces [7]string

	for i := 0; i < len(arr); i++ {
		fmt.Println("Enter your favouriate place ")
		fmt.Scanln(&arrPlaces[i])
	}

	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		var count int
		count = 0
		for j := 0; j < len(arrPlaces[i]); j++{
			count++
		}
		fmt.Println("Count of word in fav place ", arr[i] , "is : ", count)
	}
	

}