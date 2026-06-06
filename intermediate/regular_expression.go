package intermediate
import (
	"fmt"
	"regexp"
)

func main(){
	//user input
	email:= "prasannapattar966@gmail.com"
	phone:= "1234567890"

	//email regex
	emailRegex := regexp.MustCompile(
		`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
	)

	//phone regex
	phoneRegex := regexp.MustCompile(
		`^[0-9]{10}$`,
	)

	//check email

	if emailRegex.MatchString(email) {
		fmt.Println("valid email")
	}else {
		fmt.Println("invalid email")
	}

	if phoneRegex.MatchString(phone) {
		fmt.Println("valid phone number")
	} else {
		fmt.Println("invalid phone number")
	}
}