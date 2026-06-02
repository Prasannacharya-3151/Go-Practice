package intermediate
import "fmt"

func main(){
	//formatting verbs
	name := "Alice"
	age := 30
	fmt.Printf("name:%s, age:%d\n", name, age)

	pi := 3.14159
	fmt.Printf("pi: %.2f\n", pi)

	//using %v for default formatting
	fmt.Printf("name: %v, age: %v\n", name, age)

	//using %T to print the type of a variable
	fmt.Printf("name type: %T, age type: %T\n", name, age)

	//using %% to print a literal percent sign
	fmt.Printf("pi is approximately %.2f%%\n", pi*100)
}