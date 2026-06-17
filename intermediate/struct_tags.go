package intermediate
import (
	"encoding/json"
	"fmt"
)
//Metadata attached to struct fields
type Person struct {
	Name string `json:"name"` //name is the Json skip if empty field is called omitempty
	Age	 int    `json:"age"`
	City string `json:"city"`
}

func main(){
	//Struct tags tell Go how struct fields should appear in JSON, XML, or databases.
	p := Person{
		Name:"",
		Age:22,
		City:"",
	}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}