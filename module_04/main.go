package main

import (
	"fmt"
	"strings"
)

func main() {
	sayHi()

	customMessage(func(m string) {
		fmt.Println(strings.ToUpper(m))
	}, "Hey brother how are you?")
}

func sayHi() {
	fmt.Println("hey...")

	sayMessage := func() {
		fmt.Println("how are you")
	}

	sayMessage()
	sayMessage()
}

func customMessage(fn func(m string), message string) {
	fn(message)
}
