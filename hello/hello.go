package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Setting log prefix and other log properties.
	// setting flag to disable printing
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Kelvin")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
