package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	//Get Current Working Droctory
	currentDir, err := os.Getwd() //if we want to find a current directory: C:\Users\prasa\go-practice
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current Directory:")
	fmt.Println(currentDir)

	//Relative Path
	relativePath := "test.txt"  //existed path directly we want find. Relative path:test.txt

	fmt.Println("\nRealative path:")
	fmt.Println(relativePath)

	// Absolute path
	absolutePath, err := filepath.Abs(relativePath) //get a full path. Absolute path:C:\Users\prasa\go-practice\test.txt
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("\nRalative path:")
	fmt.Println(absolutePath)

	//build path using a filepath.Join()
	joinpath := filepath.Join(  //create a path
		"files",
		"data.txt",
	)

	fmt.Println("\njoined Path:")
	fmt.Println(joinpath)

	//check file is exists or not 
    _ , err = os.Stat(relativePath) //check exist os.Stat

	if err == nil {
		fmt.Println("\nFile exist:", relativePath)
	} else {
		fmt.Println("\nfile not found", relativePath)
	}
}

//age := 22 create a varible 
//age = 23 updated a variable. look at the diff in the := and = 