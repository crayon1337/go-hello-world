package main

import (
	"log"

	"fmt"

	"crayon/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Mohamed", "Menna", "Foo", "Baz"}

	messages, err := greetings.MultipleHello(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
