package main
import "fmt"

//custom errors in Go are created by defining a new type that implements the error interface.
//here is an example of how to create and use custom errors in Go:

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func doSomething() error {
	return &MyError{Message: "Something went wrong!"}
}

type PrinterError struct {
	Message string
}

func (p PrinterError) Error() string {
	return p.Message
}

func printDocument() error {
	return PrinterError{
		Message: "printer is out of paper",
	}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}

	error := printDocument()
	if error != nil {
		fmt.Println("Error:", error)
	}
}