package intermediate
import (
	"bufio"
	"fmt"
	"os"
)
//The bufio package is used to read complete input from the keyboard or files, including spaces, until the Enter key is pressed.
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)
}