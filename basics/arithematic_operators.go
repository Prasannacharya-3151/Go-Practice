package basics
import "fmt"

func main() {
	var a, b int = 10, 3
	var result int

	result = a-b
	fmt.Println("Addition:", result)
	
	result = a*b
	fmt.Println("Multiplication:", result)

	result = a/b
	fmt.Println("Division:", result)

	result = a%b
	fmt.Println("Modulus:", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)
	

	var maxInt int = 12345678923456
	fmt.Println(maxInt)

	var uMaxInt uint = 12345678906789567
	fmt.Println(uMaxInt)
}