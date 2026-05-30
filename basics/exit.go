package basics
import (
	"fmt"
	"os"
)

func main() {
	//exit the program with specific status code 
	//os.exit(status code) is used to exit the program with specific status code
	//status code 0 means successful execution of the program
	//status code 1 means unsuccessful execution of the program
	fmt.Println("start")
	os.Exit(0)
	fmt.Println("end")

}