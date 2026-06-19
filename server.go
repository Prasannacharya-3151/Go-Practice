package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello server!")
	})
	const serverAddr string = "127.0.0.1:3000" //address:127.0.0.1 and gate number 3000

	fmt.Println("Server listening on port 3000")
	err := http.ListenAndServe(serverAddr, nil)
	if err !=nil {
		log.Fatalln("error starting server", err)
	}
}