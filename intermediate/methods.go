package intermediate
import "fmt"

//methods with a value recevier
type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++
}

func (c Counter) Decrement() {
	c.Value--
}



func main() {
	counter := Counter{Value:0}
	counter.Increment()
	fmt.Println("Counter value after increment:", counter.Value)
	counter.Decrement()
	fmt.Println("Counter value after decrement:", counter.Value)
}