package main
import (
	"fmt"
 	"strings"
	"os"
	"bufio"
)

func main(){
	fmt.Println("Enter numbers with comma")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userNumber := scanner.Text()
	
	displayNum :=strings.Split(userNumber,",")
	fmt.Println(displayNum)
}