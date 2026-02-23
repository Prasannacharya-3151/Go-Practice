// package main
// import "fmt"

// func add(x int, y int) int{
// 	return x+y
// }
// func main(){
// 	result := add(3, 4)
// 	fmt.Println("the result is", result)
// }

// package main
// import "fmt"
// func add(a, b int ) {
// 	fmt.Println("inside a function :", a+b)
// }

// func main(){
// 	add(10, 5)
// }

// package main
// import "fmt"

// type student struct {
// 	Name string 
// 	Age int
// 	Marks int
// }

// func main(){
// 	s := student{
// 		Name: "John",
// 		Age: 20,
// 		Marks: 85,
// 	}
// 	fmt.Println(s.Name)
// 	fmt.Println(s.Age)
// 	fmt.Println(s.Marks)
// }


// package main 
// import "fmt"

// type Address struct {
// 	City string
// 	State string
// }
// type Person struct {
// 	Name string
// 	Address Address
// }

// func main(){
// 	p := Person{
// 		Name:"prasanna",
// 		Address : Address{
// 			City: "Banglore",
// 			State: "Karnataka",
// 		},
		
// 	}
// 	fmt.Println("Name:", p.Name)
// 		fmt.Println("City:", p.Address.City)
// 		fmt.Println("State:", p.Address.State)
// }


// package main 
// import "fmt"

// func main()  {
// 	person := struct{
// 		Name string
// 		Age int
// 	}{
// 		Name: "Prasanna",
// 		Age:30,
// 	}
// 	fmt.Println("Name", person.Name)
// 	fmt.Println("Age", person.Age)

// }

// package main
// import "fmt"

// type Person struct{
// 	Name string 
// 	Age int
// }

// func (p Person) greet(){
// 	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
// }

// func main(){
// 	p1 := Person{
// 		Name: "Prasanna",
// 		Age: 30,
// 	}
// 	p1.greet()
// }


// package main
// import (
// 	"fmt"
// 	"errors"
// )

// func divide(a, b int) (int, error){
// 	if b==0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main(){
// 	result, err := divide(10, 0)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return 
// 	}

// 	fmt.Println("result:", result)
// }

// package main
// import "fmt"

// type MyError struct {
// 	Code int
// 	Message string
// }

// func (e MyError) Error() string {
// 	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
// }

// func main(){
// 	err := MyError{ 404, "page not Found"}
// 	fmt.Println(err.Error())
// }


// package main 
// import "fmt"

// func main(){
// 	marks := []int{85, 90, 78, 92, 88}

// 	total := 0

// 	for _, mark := range marks {
// 		total += mark
// 	}
// 	fmt.Println("Total marks:", total)
// }

// package main 
// import "fmt"

// func main(){
// 	marks := []int{85, 90, 78, 92, 88}
// 	total := 0

// 	for _, mark := range marks {
// 		total += mark
// 	}
// 	average := total /len(marks)
// 	fmt.Println("Average marks:", average)
// 	fmt.Println("Total marks:", total)
// }


// package main
// import "fmt"

// func main(){
// 	var n int 
// 		fmt.Println("Enter a number of subjects:")
// 		fmt.Scan(&n)

// 		marks := make([]int, n)
// 		for i := 0; i<n; i++ {
// 			fmt.Printf("Enter mark %d: ", i+1)
// 			fmt.Scan(&marks[i])
// 		}

// 		total := 0
// 		for _, mark := range marks {
// 			total += mark
// 		}

// 		average := total / n
// 		fmt.Println("Total marks:", total)
// 		fmt.Println("Average marks:", average)
// 	}


