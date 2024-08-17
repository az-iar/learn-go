package main

import "fmt"

const (
	french  = "French"
	malay   = "Malay"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	malayHelloPrefix   = "Salam, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case malay:
		prefix = malayHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Azri", "Malay"))
}
