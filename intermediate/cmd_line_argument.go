// package main
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("All Argumetns")
// 	fmt.Println(os.Args) //os.Args is a slice containing all command-line argumnets

// 	fmt.Println("First Arguments")
// 	fmt.Println(os.Args[1])
// }


package intermediate

import (
	"flag"
	"fmt"
)

func main() {
   //Command-line arguments are values passed to a program during execution and can be accessed using os.Args. The flag package provides a structured way to define and parse named command-line options.
	name := flag.String(
		"name",
		"Guest",
		"User Name",
	)

	age := flag.Int(
		"age",
		0,
		"User Age",
	)

	flag.Parse() //read a command line flages

	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
}