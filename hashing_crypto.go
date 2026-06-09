package main
import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	data := "Prsanna"

	sha256Hash := sha256.Sum256([]byte(data))

	sha512Hash := sha512.Sum512([]byte(data))

	fmt.Println("Original:")
	fmt.Println(data)

	fmt.Printf("\nSHA256:\n%x\n", sha256Hash)

	fmt.Printf("\nSHA512:\n%x\n", sha512Hash)

}