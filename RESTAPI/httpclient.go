package RESTAPI

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//create a new http client
	client := &http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil{
		fmt.Println("Error making a get request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		println("Error handling response body:", err)
		return
	}
	fmt.Println(string(body))
}