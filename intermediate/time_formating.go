package intermediate
import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Original:", now)

	fmt.Println("date only:",
now.Format("02-02-2005"))

fmt.Println("time only:",
now.Format("15:04:05"))

fmt.Println("Date & Time:",
now.Format("02-01-2005 15:04:05"))

fmt.Println("Year:",
now.Format("2006"))

fmt.Println("Month:",
now.Format("01"))

fmt.Println("Day:",
now.Format("02"))
}