package main

import (
	"fmt"
	"time"
)

func main() {

	// Current Time
	now := time.Now()

	fmt.Println("Normal Time:")
	fmt.Println(now)

	// Convert Time => Epoch
	epoch := now.Unix()

	fmt.Println("\nEpoch Time:")
	fmt.Println(epoch)

	// Convert Epoch => Normal Time
	converted := time.Unix(epoch, 0)

	fmt.Println("\nConverted Back:")
	fmt.Println(converted)
}