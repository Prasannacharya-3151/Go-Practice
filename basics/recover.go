package basics 
import "fmt"

func main() {
	//recovre form the panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from the panic:", r)
		}
	}()

	//panic and recover
	//panic is used to raise an error and stop the execution of the program
	//recover is used to handle the panic and continue the execution of the program
}