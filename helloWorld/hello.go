// This will display in the Go Doc for your function so it can make the intent of your code clearer.package main
package main

// Constants

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = " Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// In our function signature we have made a named return value (prefix string).

// This will create a variable called prefix in your function.
// It will be assigned the "zero" value. This depends on the type, for example ints are 0 and for strings it is "".
// You can return whatever it's set to by just calling return rather than return prefix.

//This will display in the Go Doc for your function The function name starts with a lowercase letter. In Go, public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.so it can make the intent of your code clearer.

// The function name starts with a lowercase letter.
// In Go, public functions start with a capital letter and private ones start with// The function name starts with a lowercase letter.The function name starts with a lowercase letter. In Go, public functions start with a capital letter and private ones start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.
//  a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this function private.

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

func main() {

}
