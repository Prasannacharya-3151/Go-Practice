package intermediate

//text templates in Go can be created using the text/teplate package.this package provides a way to create and execute templates that can generate text output based on data.Here is an example of how to use text templates in Go:

import "text/template"
import "os"

// type Person struct {
// 	Name string
// 	Age int
// }

func main() {
	// person := Person{Name:"Alice", Age: 30}
	// tmpl, err := template.New("Person").Parse("My name is {{.Name}} and I am {{.Age}} years old.")
	// if err != nil {
	// 	panic(err)
	// }
	// err = tmpl.Execute(os.Stdout, person)
	// if err != nil {
	// 	panic(err)
	// }


	tmpl := template.Must(
		template.New("test").Parse(
			"Hello {{.Name}}",
		),
	)

	data := struct {
		Name string
	}{
		Name:"Prasanna",
	}

	tmpl.Execute(os.Stdout, data)
}