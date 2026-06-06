package intermediate
import "fmt"

//string formating in Go can be done sting concatenation, using the fmt.sprintf fucntion, or using the fmt.Printf function. Here are some examples of string formating in Go:

func main() {
	//string concatenation
	greeting := "hello"
	name := "world"
	fmt.Println(greeting + " " + name)

	//using fmt.Sprintf
	greeting2 := fmt.Sprintf("%s %s", greeting, name)
	fmt.Println(greeting2)

	//using fmt.Printf
	fmt.Printf("%s %s\n", greeting, name)

	//using fmt.Sprintf with different data types
	age := 30
	info := fmt.Sprintf("My name is %s and I am %d year old.", name, age)
	fmt.Println(info)
}