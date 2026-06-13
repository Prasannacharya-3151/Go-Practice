package main
import (
	"fmt"
	"os"
)

func main() {
	//create a temporory file
	//Temporary files and directories are created for short-term storage during program execution. Go provides os.CreateTemp() and os.MkdirTemp() to create temporary files and folders, which can be removed using os.Remove() and os.RemoveAll().
	tempFile, err := os.CreateTemp("", "example-*.text")  //create a temporary file example-123456.txt ..go automatically generates a unique name and then ...."" = use system temp folder exaxmple like a C:\Users\prasa\AppData\Local\Temp
	if err != nil {
		fmt.Println(err)
		return 
	}

	defer os.Remove(tempFile.Name())

	fmt.Println("Temp File:")
	fmt.Println(tempFile.Name())

	//write data
	tempFile.WriteString("Hello Prasanna") //write a text into a temp file.

	//close file
	tempFile.Close()

	//Create temporary directory
	tempDir, err := os.MkdirTemp("", "mydir-*") //creates a mydir-123456
	if err != nil {
		fmt.Println(err)
		return 
	}

	defer os.RemoveAll(tempDir) //when programs end delete file 1)folder, all files, all subfolders

	fmt.Println("\nTemp Directory")
	fmt.Println(tempDir)
}