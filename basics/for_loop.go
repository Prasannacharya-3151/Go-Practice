package basics
import "fmt"

func main() {
	//simple iteration over range of numbers
	// for i:=1; i<=5; i++ {
	// 	fmt.Println(i)
	// } 

	// //iterate over collection of items
	// numbers := []int{1,2,3,4,5,6}
	// for index, value := range numbers{
	// 	fmt.Printf("index: %d, value:%d\n", index, value)
	// }


	// for i:=1; i<=10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("odd number:", i)
	// 		continue
	// 	}
	// 	if i ==5 {
	// 		break
	// 	}
	// }

	//outer loop
	rows := 5
	for i:=1; i<=rows; i++ {
		//inner loop for spaces before stars
		for j:=1; j<=rows-i; j++ {
			fmt.Println("  ")
		}

		//inner loopp for a stars
		for k:=1; k<=2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}