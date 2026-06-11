package intermediate
import (
	"fmt"
	"os"
)

func main(){
	//reading files means loading data form the file into your go program
	data, err := os.ReadFile("test.txt") //reading the file with entire line in this text file 

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data)) //converting a byte to string
}