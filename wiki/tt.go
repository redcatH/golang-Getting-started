package main

import (
	"fmt"
	"regexp"
)

func main() {
	validPath := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fmt.Println(validPath.FindStringSubmatch("/edit/hello"))
}
