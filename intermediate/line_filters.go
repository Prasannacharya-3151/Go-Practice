package intermediate
import (
	"fmt"
	"os"
	"strings"
	"bufio" //this package is used for the read every line upto user entering the enter button forom the keyboard
)

func main() {
	file, err := os.Open("test.txt") //open a file for a reading
	if err != nil {
		fmt.Println(err) //if the error goes when executing reading file shown a error
		return
	}
	defer file.Close() //

	scanner := bufio.NewScanner(file) //read a file line by line

	for scanner.Scan() {  //keep a reading lines until file ends
		line := scanner.Text() //get a current line text 

		if strings.Contains(line, "go") { //does lines containes a text go
			fmt.Println(line) //print it
		}
	}
}