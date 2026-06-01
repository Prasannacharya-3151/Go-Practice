package main
import "fmt"

func main() {
	//string and runes
	str := "hello world"
	fmt.Println("string:", str)

	//runes
	runes := []rune(str)
	fmt.Println("runes:", runes)
	fmt.Printf("%c\n", 104)

	greeting := "hello"
	name := "world"
	fmt.Println(greeting + "" + name)

	str1 := "Apple"
	
}