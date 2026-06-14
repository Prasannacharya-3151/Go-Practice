package main
import (
	"encoding/json"
	"fmt"
)

type Person struct { //go struct foramte
	Name string `json:"name"`
	Age int `json:"age"`
	City string `json:"city"`
}

func main() {

	//struct
	person := Person{ //go struct
		Name: "Prasanna",
		Age: 22,
		City: "Bangalore",
	}

	//marshal
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Json:")
	fmt.Println(string(jsonData))

	//unmarshal
	jsonString := `{
	"name":"Rahul",
	"age":25,
	"city":"Mysore"
	}`

	var p Person

	err = json.Unmarshal(
		[]byte(jsonString),
		&p,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nStruct:")
	fmt.Println(p)
}