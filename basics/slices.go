package basics
import "fmt"

func main() {
	//slices in go
	//slices are dyamnic arrays that can grow and shrink in size
	//slcies are built on top of arrays and provide a more flexible way to work with collection of data
	//slices are referance types, which means that when you assign a slices to another variable, both variables point to the same underlying array
	//creating slices
	var number []int= []int{1,2,3,4,5}
	fmt.Println("numbers:", number)

	//slices can be creaating a make function
	slice1 := make([]int,5)
	fmt.Println("slice1:", slice1)

	//slices can be created form the array
	array := [5]int{10,2,30,40,50}
	slices2 := array[1:4]
	fmt.Println("slices2:", slices2)

	//slices are created woth the legth of the slices and the capacity of the underlying array
}