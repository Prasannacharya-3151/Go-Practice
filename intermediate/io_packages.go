package intermediate

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	//The io package provides interfaces and helper functions for reading and writing data. The main interfaces are io.Reader and io.Writer.
	//bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Name:")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)

	//io
	strReader := strings.NewReader(
		"Hello Go",
	)

	data, _ := io.ReadAll(strReader) //read entire content
	fmt.Println(string(data))
}