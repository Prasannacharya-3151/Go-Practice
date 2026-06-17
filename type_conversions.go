package main
import "fmt"

func main(){
	//Type conversion is the process of converting one data type into another. Go requires explicit type conversion and does not perform automatic conversions between incompatible types.

	//int to float
	var a int = 10

	var b float64 = float64(a)


	fmt.Println("Int:", a)
	fmt.Println("Float:", b)

	//float64 to int

	var pi float64 = 3.14

	var num int = int(pi)

	fmt.Println("\nFloat:", pi)
	fmt.Println("Int:", num)

	//int to string

	var x int = 100

	fmt.Println("\nNumber:", x)

	fmt.Println("String:", fmt.Sprint(x))

	//string to int

	//strconv.Atoi() is used here
}