package intermediate
import (
	"fmt"
	"math/rand"
)

func main() {
	//random integer
	fmt.Println("Random int:")
	fmt.Println(rand.Int())

	//radnom interger 0-9
	fmt.Println("\nRandom Intn(10):")
	fmt.Println(rand.Intn(10))

	//Random Interger 1-100
	fmt.Println("nRandom Numner 1-100:")
	fmt.Println(rand.Intn(100) +1)

	//radnom float64
	fmt.Println("\nRandom float64:")
	fmt.Println(rand.Float64())

	//radnom flaodt32
	fmt.Println("\nRandom float32:")
	fmt.Println(float32(rand.Float64()))

	//rand otp
	fmt.Println("\nRandom OTP:")
	fmt.Println(rand.Intn(900000) + 100000)

	//random boolean
	fmt.Println("\nRandom Boolean:")
	fmt.Println("\nRandom Boolean:")
	fmt.Println(rand.Intn(2) == 1 )

	//radnom element form slice
	colors := []string{
		"red",
		"Blue",
		"green", 
		"yellow",
	}

	randomColor := colors[rand.Intn(len(colors))]

	fmt.Println("\nRadnom Color:")
	fmt.Println(randomColor)
}