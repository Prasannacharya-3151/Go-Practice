package main
import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"Name"`
	Age  int    `xml:"Age"`
	City string `xml:"City"`
}

func main(){
	data := `
	<Person>
	<Name>Rahul</Name>
	<Age>25</Age>
	<City>Mysore</City>
	</Person>
    `

	var p Person

	xml.Unmarshal(
		[]byte(data),
		&p,
	)

	fmt.Println(p)
}