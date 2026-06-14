package intermediate
import (
	"fmt"
	"os"
)

func main() {
	//A subcommand is a command that operates under a main command. In Go, subcommands are commonly implemented by reading os.Args and executing different logic based on the specified command.
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [hello|bye]")
	}

	switch os.Args[1] { //it will checkes the which subcommand was entered?

    case "hello":
		fmt.Println("Hello Prasanna")
	case "bye":
		fmt.Println("Good bye")
	default:
		fmt.Println("Unknown Command")
   }
}