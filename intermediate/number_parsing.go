package intermediate
import (
	"fmt"
	"strconv"

)

func main() {
	//string to int
	intStr := "123"
	intValue, _ := strconv.Atoi(intStr)

	fmt.Println("String:", intStr)
	fmt.Println("int", intValue)

	//string to int64
	int64Str := "999999"
	int64Value, _ := strconv.ParseInt(
		int64Str,
		10,
		54,
	)

	fmt.Println("\nInt64:", int64Value)

	//string to float64
	floatStr := "3.14159"
	floatValue, _ := strconv.ParseFloat(
		floatStr,
		64,
	)
	fmt.Println("\nFloat64:", floatValue)

	//string to bool
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(
		boolStr,
	)

	fmt.Println("\nBool:", boolValue)

	//number to string
	number := 500
	numberStr := strconv.Itoa(number)

	fmt.Println("\nNumber:", number)
	fmt.Println("String:", numberStr)
}