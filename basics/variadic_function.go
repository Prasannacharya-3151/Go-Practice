package basics
import "fmt"

func main() {
	//variadic function are functions that can take a variable number of arguments.
	//...elipsis 
	//func functionName(param1 type1, parm2 type2, param3 ...type3) returnType {
	//function body
	//}
	sum := variadicFunction(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("sum:", sum)
}

func variadicFunction(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num

	}
	return total
}

//in this example, varidic function takes a variable number of int argumets and calculates the sum of those numbers.
//when we call the variadic function, we can pass any number of int arguments, and the function will calculate the sum of those numbers and return the result.
// variadic functions are useful when we want to perform operations on a variable number of arguments without having to define multiple functions for different numbers of parameters.
//variadic functions can also be used with other types, such as strings, floats, etc. depending on the use case.
