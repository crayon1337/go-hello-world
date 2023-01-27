package main

import (
	"log"

	"fmt"

	"crayon/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Mohamed")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
