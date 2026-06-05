package intermediate
import "fmt"

//errors in go are represented by the built-in error type, which is an interface that as a single method Error() string.
//you can create custom error types by implementing the error interface. Here's an example of how to create and use custom errors in Go:

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func doSomthing() error {
	return &MyError{Message: "Something went wrong!"}
}

func main() {
	err := doSomthing()
	if err != nil {
		fmt.Println("Error:", err)
	}
}