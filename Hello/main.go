package hello

import "fmt"

func main() {
	fmt.Println(Hello("Alex", "English"))
}

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return gettingPrefix(language) + name
}

func gettingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return
}
