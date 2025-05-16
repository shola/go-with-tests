package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s%s", prefix, name)
}

func main() {
	fmt.Println(Hello("Olusola", ""))
}
