package intermediate
import "fmt"

func main () {
	//fmt packages 
	fmt.Println("Hello, world!")
	fmt.Printf("The value of pi is approximately %.2f\n", 3.14159)

    str := "go is greate then python"
	fmt.Printf("the string is: %s\n", str)

	sprintfStr := fmt.Sprintf("the string is %s\n", str)
	fmt.Print(sprintfStr)
}