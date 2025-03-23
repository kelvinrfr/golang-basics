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

	names := []string{"Kelvin", "Nubia", "Gladys"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
