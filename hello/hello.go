package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	// prefix := englishHelloPrefix

	// switch language {
	// 	case "Spanish":
	// 		prefix = spanishHelloPrefix
	// 	case "French":
	// 		prefix = frenchHelloPrefix
	// }
	return greetingPrefix(language) + name

	// if language == "Spanish" {
	// 	return spanishHelloPrefix + name
	// }

	// if language == "French" {
	// 	return frenchHelloPrefix + name
	// }
	// return englishHelloPrefix + name
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

func main() {
	fmt.Println(Hello("world", "english"))
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("test", ""))
	fmt.Println(Hello("", "english"))
}
