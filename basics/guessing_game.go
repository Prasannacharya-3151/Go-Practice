package basics
import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1

	fmt.Println("welocm to the gussing game!")
	fmt.Println("I have selected a number between 1 and 100. Can you guess it?")
	fmt.Println("can u guess what it is ")

	var guess int
	for {
		fmt.Print("enter you guess:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("congradulation! you guessed the number correctly!")
		break
		}else if guess < target {
			fmt.Println("too low! try again.")

		}else {
			fmt.Println("too high! try again.")
		}
	}

}