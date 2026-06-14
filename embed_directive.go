package main
import (
	"fmt"
	_ "embed"
)

//go;embed message.txt
var message string

func main(){
	// The //go:embed directive is used to include files and folders directly inside the Go executable.
	fmt.Println(message)
}