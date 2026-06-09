package intermediate
import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {
	data1 := "Prsanna"

	password := "password123"
	salt := "abc123"


	sha256Hash := sha256.Sum256([]byte(data1))

	sha512Hash := sha512.Sum512([]byte(data1))

	data := salt + password

	hash := sha256.Sum256([]byte(data))

	fmt.Println("Original:")
	fmt.Println(data)

	fmt.Printf("\nSHA256:\n%x\n", sha256Hash)

	fmt.Printf("\nSHA512:\n%x\n", sha512Hash)

	fmt.Printf("\n%x\n",hash)
}