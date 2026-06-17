//XML is a markup language used for storing and exchanging structured data. In Go, the encoding/xml package is used to convert structs into XML using Marshal() and XML into structs using Unmarshal().
//unmarshal
// package main
// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Person struct {
// 	Name string `xml:"Name"`
// 	Age  int    `xml:"Age"`
// 	City string `xml:"City"`
// }

// func main(){
// 	//XML (eXtensible Markup Language) is a markup language used to store and exchange structured data.
// 	data := `
// 	<Person>
// 	<Name>Rahul</Name>
// 	<Age>25</Age>
// 	<City>Mysore</City>
// 	</Person>
//     `

// 	var p Person

// 	xml.Unmarshal(
// 		[]byte(data),
// 		&p,
// 	)

// 	fmt.Println(p)
// }

//marshal(struct-xml)
package intermediate
import (
	"encoding/xml"
	"fmt"
)

type Person struct{
	XMLName xml.Name `xml:"Person"`

	Name string `xml:"Name"`
	Age  int    `xml:"Age"`
	City string `xml:"City"`
}

func main(){
	p := Person{
		Name:"Prasanna",
		Age:22,
		City:"Banglore",
	}

	data, _ := xml.MarshalIndent(
		p,
		"",
		"   ",
	)

	fmt.Println(string(data))
}