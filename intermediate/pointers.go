package basics
import "fmt"

func main() {
	//pointers in go
	var x int = 10
	var p *int = &x //p is a pointer to an interger variable x
	fmt.Println("value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("value of p:", p)
	fmt.Println("value at address p:", *p) //derefering the pointer to get the value of x

	*p = 20 //modifying the value of x using pointer p
	fmt.Println("value of x after modification:", x)
}