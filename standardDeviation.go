package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	var floatNumArray []float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("************************")
	fmt.Println("Enter \"q\" or \"Q\" to Exit")
	fmt.Println("Enter floating point nums :")
	
	for scanner.Scan() {
		userInput := scanner.Text()
		
		if strings.ToLower(userInput) == "q" {
			break
		}

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter valid floating point number : ")
			continue
		}

		floatNumArray = append(floatNumArray, num)
	}

	average := findAverage(floatNumArray)
	standardDeviation := findStandardDeviation(floatNumArray)
	fmt.Printf("Average : %f\n", average)
	fmt.Printf("Standard deviation : %f\n", standardDeviation)
}

func findAverage(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func findStandardDeviation(nums []float64) float64 {
	sumOfSquares := 0.0

	average := findAverage(nums)

	for _, num := range nums {
		sumOfSquares += (num - average) * (num - average)
	}

	return math.Sqrt(sumOfSquares / float64(len(nums)))
}