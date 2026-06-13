package main
import (
	"fmt"
	"os"
)

func main(){
	//A directory is a folder that stores files and subdirectories. Go provides functions such as os.Mkdir(), os.ReadDir(), and os.Remove() to create, read, and delete directories.
	//create a dictionary
	err := os.Mkdir("myfolder", 0755) //it will create a myfolder

	if err != nil {
		fmt.Println("Create Error:", err)
	} else {
		fmt.Println("Directory Created")
	}

	//read a current directory content
	fmt.Println("\nFiles and Folders:")

	entries, err := os.ReadDir(".") //it means read a current folder means current directory

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {         //check a folder or files
			fmt.Println("[DIR]", entry.Name())
		} else {
			fmt.Println("[FILE]", entry.Name())
		}
	}

	//remove directory
	err = os.Remove("myfolder")

	if err != nil {
		fmt.Println("Remove Error:", err)
	} else {
		fmt.Println("\nDirectory removed")
	}
}