package main

import "fmt"

const (
	spanish = "Spanish"
	english = "English"
	french  = "French"

	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Chaz", "English"))
}

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	case english:
	default:
		prefix = englishHello
	}
	return
}
