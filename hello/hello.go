package main

import (
	"fmt"
	"github.com/az-iar/learn-go/greetings"
)

func main() {
	fmt.Println("Hello, World!")

	message := greetings.Hello("Azri")
	fmt.Println(message)
}
