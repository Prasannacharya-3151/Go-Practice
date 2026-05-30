package basics
import "fmt"

func main() {
	//defer statement
	//defer function call is executed after the surroudning fucntion returns
	//defer statement is often used to ensure that resources are relesed or to perform cleanup tasks

	// fmt.Println("start of main fucntion")
	// defer fmt.Println("deferred statement 1")
	// defer fmt.Println("deferred statemtent 2")
	// fmt.Println("end of main function")

	process(10)
}
	func process(i int) {
		defer fmt.Println("deferred statement 1")
		defer fmt.Println("deferred statement 2")
		fmt.Println("processing:", i)
}