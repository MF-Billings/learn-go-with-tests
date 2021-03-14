package main

import "fmt"

const french = "French"
const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const japaneseHelloPrefix = "今日は"

// Hello says Hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}