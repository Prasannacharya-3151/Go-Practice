package main
import "fmt"

func main() {
	//structs in go 
	type Person struct {
		Name string
		Age int
	}
	p1 := Person{Name: "Alice", Age:30 }
	fmt.Printf("Person 1:Name: %s, Age: %d\n", p1.Name, p1.Age)
	fmt.Printf("Person 1: Name: %s, Age: %d", p1.Name, p1.Age)

	//structures with nested structs
	type Address struct {
		City string
		Country string
	}
	type Employee struct {
		Name string
		Age int
		Address Address
	}
	emp := Employee{
		Name: "Bob",
		Age: 25,
		Address: Address{
			City: "New York",
			Country: "USA",
		},
	}
	fmt.Printf("Employee:Name: %s, Age:%d, city:%s, country:%s\n", emp.Name, emp.Age, emp.Address.City, emp.Address.Country)

	//anonymous struct
	annon := struct {
		Name string
		Age int
	}{
		Name: "Charlie",
		Age: 28,
	}
	fmt.Println("Anonymous struct: Name:", annon.Name, "Age:", annon.Age)

	
	circle := Circle{Radius: 5}
	fmt.Println("Circle area:", circle.Area())
}

    //structs with methods
	type Circle struct {
		Radius float64
	}
	func (c Circle) Area() float64 {
		return 3.14 * c.Radius * c.Radius
	}
    
		