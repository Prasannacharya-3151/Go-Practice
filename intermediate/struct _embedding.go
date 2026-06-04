package intermediate
import "fmt"

//struct embedding is a powerful feature in Go that allows you to create complex data structure by embedding one struct within another, this promotes code reuse and allows you to create more organization and modularity in your code.
//in this example we have an address struct that is embedded within the employee struct, this allows us to access the fields of the address struct directly form the employee struct with out needing to reference the address struct explicity.
type Address struct {
	City string
	Country string
}

type Employee struct {
	Name string
	Age int
	Address //emvedding the address struct in here 
}

func main() {
	emp := Employee{
		Name: "Bob",
		Age:25,
		Address: Address{
			City: "New York",
			Country: "USA",
		},
	}
	fmt.Printf("Employee:Name: %s, Age: %d, City: %s, Country: %s\n", emp.Name, emp.Age, emp.City, emp.Country)
	fmt.Printf("City:%s\n",emp.City)

	//anonymous embedding struct
	annon := struct {
		Name string
		Age int
	}{
		Name: "charlie",
		Age: 28,
	}
	fmt.Println("Anonymous struct:Name:", annon.Name, "Age:", annon.Age)

}