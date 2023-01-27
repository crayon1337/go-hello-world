package main

import (
	"fmt"

	"crayon/greetings"
)

func main() {
	message := greetings.Hello("Mohammed")

	fmt.Println(message)
}
