package basics
import "fmt"

func main() {
	//range is used to iterate over elements in a collection such as arrays, slices, maps, etc.
	//iterate over a slice of numbers


	// numbers := []int{1,2,3,4,5}
	// for index, value := range numbers {
	// 	fmt.Println("index:", index, "value:", value)
	// }

	messages := "hello world"
   for i, v := range messages {
	fmt.Println(i,v)
	fmt.Printf("index:", %d, "Rune:", %c\n", i,v)
   }

}