package intermediate
import (
	"fmt"
	_ "embed" //enabled a embded feature
	"embed"
)

//go;embed message.txt
var message string //embeded files into a executable

var files embed.FS //this one is the embeded file system A virtual folder inside your program

func main(){
	// The //go:embed directive is used to include files and folders directly inside the Go executable.
	fmt.Println(message)

	data, _ := files.ReadFile("message.tsx") //read embede file
	fmt.Println(string(data))
}