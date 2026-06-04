package intermediate

import "fmt"

// Interface
type Animal interface {
	MakeSound()
}

// Dog Struct
type Dog struct{}

// Dog Method
func (d Dog) MakeSound() {
	fmt.Println("Woof Woof")
}

// Cat Struct
type Cat struct{}

// Cat Method
func (c Cat) MakeSound() {
	fmt.Println("Meow Meow")
}

func main() {

	var animal Animal

	animal = Dog{}
	animal.MakeSound()

	animal = Cat{}
	animal.MakeSound()
}