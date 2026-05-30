package main
import "fmt"

func main(){
	//function with parameter with return type
	//func <name> (parameter list) return type{
	//code block 
	//return value
	//}


	// result := add(5, 10)
	// fmt.Println("result:", result)

	fmt.Println(add(5,10))

	greet := func() {
		fmt.Println("Helloo anonymous function")
	}
	greet()

	operation := add
	result := operation(20, 30)
	fmt.Println("result:", result)
}

func add(a int, b int) int {
	return a+b
}