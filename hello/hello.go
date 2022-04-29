package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(1 | 2 | 3)
	for k, v := range []int{1, 2, 3} {
		fmt.Println(k, v)
	}
	// message, err := greetings.Hello("liuyi")
	message, err := greetings.Hellos([]string{"liuyi", "liuyi2"})
	for _, hello := range message {
		fmt.Println(hello)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
