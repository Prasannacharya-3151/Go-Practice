package basics
import "fmt"

func main() {
	//recursion is a itself calling function
	result := factorial(5)
	fmt.Println("factorial of 5 is:", result)

}

func factorial(n int) int {
	if n==0 {
		return 1
	}
	return n * factorial(n-1)
}