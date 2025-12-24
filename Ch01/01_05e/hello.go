package main

import (
	"fmt"
 "hello/utils"
)

func main() {
	fmt.Println("Hello from Go!")

	name := utils.Input("What is your name? ")
	age := utils.Input("How old are you? ")

	fmt.Printf("Hello %s, you are %s years old!\n", name, age)
}
