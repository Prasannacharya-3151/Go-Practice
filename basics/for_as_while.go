package basics
import "fmt"

func main(){
	//for as while with break
	sum:= 0
	for {
		sum += 10
		fmt.Println("sum:", sum)
		if sum >= 50 {
			break
		}
	}
}