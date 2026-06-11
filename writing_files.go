// package main
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	//Writing files means storing data from your Go program into a file on your computer.
// 	data := "Hello Prasanna"

// 	err := os.WriteFile(
// 		"test.txt",
// 		[]byte(data),
// 		0644,
// 	)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return 
// 	}

// 	fmt.Println("File Written successfully")
// }

package main
import (
	"fmt"
	"time"
	"os"
)

func main(){
	//first write
	os.WriteFile(
		"test.txt",
		[]byte("hello"),
		0644,
	)

	//wait a 1 min
	time.Sleep(1 * time.Minute)

	//open file in append mode
	file, _ := os.OpenFile(
		"test.txt",
		os.O_APPEND|os.O_WRONLY,
		0644,
	)

	defer file.Close()
	file.WriteString("\nPrasanna")

	fmt.Println("file written successfully")
}