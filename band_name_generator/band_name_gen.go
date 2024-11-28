package main

import (
	"fmt"
)

func main() {
	var name string
	var pet_name string
	fmt.Println("welcome to band name generator")
	fmt.Print("what is your name: ")
	fmt.Scanf("%s", &name)
	// name := fmt.Scanf("%d", "what is your name ")
	fmt.Print("what is your pet name: ")
	fmt.Scanf("%s", &pet_name)

	fmt.Println("your rock band name is", name, pet_name)
}
