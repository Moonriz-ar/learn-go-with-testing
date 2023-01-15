package main

import "fmt"

// how to test this?
// it is good to separate the "domain" code from the ouside world (side effects). The fmt.Println is a side effect (printing to stdout)
// and the stroing we send in is our domain
// so lets separate these concerns so it's easier to test

// the CYCLE
// 1/ write a test
// 2/ make the compiler pass
// 3/ run the test, see that it fails and check if the error message is meaningful
// 4/ write enough code to make the test pass
// 5/ refactor

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Mario", "French"))
}

// Hello returns a string saying "Hello, world"
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// private function, with named return value (prefix string)
// this will create a variable called prefix in function, that will be assigned the zero value
func greetingPrefix(language string) (prefix string) {
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
