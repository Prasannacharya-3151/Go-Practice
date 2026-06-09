package intermediate
import (
	"fmt"
	"encoding/base64"
)

func main() {
	//Base64 is an encoding technique used to convert binary or text data into a text representation so it can be safely transmitted through systems that handle text data.

	//original string
	text := "hello world"

	//encode
	encoded := base64.StdEncoding.EncodeToString(
		[]byte(text), //it converts a string to byte.base64 function work with bytes.
	)

	fmt.Println("Encoded:")
	fmt.Println(encoded)

	//Decode
	decoded, _ := base64.StdEncoding.DecodeString(
		encoded,
	)

	fmt.Println("\nDecoded:")
	fmt.Println(string(decoded))
}