package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//Get Current Working Droctory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current Directory:")
	fmt.Println(currentDir)

	//Relative Path
	relativePath := "test.txt"

	fmt.Println("\nRealative path:")
	fmt.Println(relativePath)

	// Absolute path
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("\nRalative path:")
	fmt.Println(absolutePath)

	//build path using a filepath.Join()
	joinpath := filepath.Join(
		"files",
		"data.txt",
	)

	fmt.Println("\njoined Path:")
	fmt.Println(joinpath)
}