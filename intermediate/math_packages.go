package intermediate
import (
	"fmt"
	"math"
)

func main(){
	//The math package provides standard mathematical functions such as square root, powers, trigonometric functions, logarithms, rounding operations, and mathematical constants like Pi.
	num := 25.0

	//squre root
	fmt.Println("Sqrt:", math.Sqrt(num))
	
	//power
	fmt.Println("Power:", math.Pow(2,3))

	// Round
	fmt.Println("Round:", math.Round(3.7))

	// Floor
	fmt.Println("Floor:", math.Floor(3.9))

	// Ceil
	fmt.Println("Ceil:", math.Ceil(3.1))

	// Absolute Value
	fmt.Println("Abs:", math.Abs(-10))

	// Maximum
	fmt.Println("Max:", math.Max(20, 50))

	// Minimum
	fmt.Println("Min:", math.Min(20, 50))

	// Pi
	fmt.Println("Pi:", math.Pi)

}

//other usefull functions
// math.Sin()
// math.Cos()
// math.Tan()
// math.Log()
// math.Log10()
// math.Exp()