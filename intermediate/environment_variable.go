package intermediate
import (
	"fmt"
	"os"
)

func main() {
	//Environment variables are external configuration values provided by the operating system. Go uses functions such as os.Getenv(), os.Setenv(), os.LookupEnv(), and os.Unsetenv() to manage environment variables.
	//set Variable
	os.Setenv("Name", "Prasanna")  //create varibale

	//Get variable
	fmt.Println("NAME:", os.Getenv("NAME")) //read variable

	//check varibale
	value, exist := os.LookupEnv("NAME") //check the variable exists

	fmt.Println("Value:", value)
	fmt.Println("Exist:", exist)

	//Remove Variable
	os.Unsetenv("NAME") //delete variable 

	fmt.Println("after delete:")

	fmt.Println(os.Getenv("NAME"))
}