package basics
import "fmt"

func main(){
	//function with multiple return values 
	result1, result2 := multipleReturns()
	fmt.Println("result1:", result1)
	fmt.Println("result2:", result2)
}

func multipleReturns() (string, int) {
	return "heloo", 42
}