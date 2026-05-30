package baiscs
import "fmt"

func main(){
	//maps in go are unordered collections of key-value pairs
	//map[keyType]valueType
	//creating a map
	ages := make(map[string]int)
	//adding a key-value pair to the map
	ages["Alices"] = 30
	ages["bob"] = 25
	fmt.Println("ages:", ages)

	//accessing a value from the map using its key 
	aliceAge := ages["Alices"]
	fmt.Println("alices age:", aliceAge)

	//deleting a key-value pair from the map
	delete(ages, "bob")
	fmt.Println("ages after deletion:", ages)

	//checking if key exists in the map
	age, exists := ages["bob"]
	if exists {
		fmt.Println("bob age:", age)
	} else {
		fmt.Println("bob does not exists in the map")
	}
}