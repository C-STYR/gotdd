package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("Chris", ""))
}

func Hello(str, language string) string {
	if str == "" {
		return "Hello, World"
	}

	return greetingPrefix(language) + str
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
