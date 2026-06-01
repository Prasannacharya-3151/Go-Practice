package intermediate
import "fmt"

func main() {
	//closure
	//A clouser is a function that captures and retains access to variables from its surrounding scope, even after that cope has finished executing.
	//example of clouser
	counter := 0 
	increament := func() int {
		counter++
		return counter
	}
	fmt.Println(increament())
	fmt.Println(increament())
	fmt.Println(increament())
}