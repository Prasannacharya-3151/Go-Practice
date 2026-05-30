package basics
import "fmt"

func main() {
	//panic and recover
	//panic is used to raise an error and stop the execution of the program
	//recover is used to handle the panic and continue the execution of the program

	//panic(interface{})
	//example of valid input
	process(10)

	process(-6)

}

func process(input int){
	if input <0 {
		panic("input cannot be nagative")

	}
	fmt.Println("proccessing:", input)
}