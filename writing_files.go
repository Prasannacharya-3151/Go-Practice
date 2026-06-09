package main
import (
	"fmt"
	"os"
)

func main() {
	//Writing files means storing data from your Go program into a file on your computer.
	data := "Hello Prasanna"

	err := os.WriteFile(
		"test.txt",
		[]byte(data),
		0644,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return 
	}

	fmt.Println("File Written successfully")
}