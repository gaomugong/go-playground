package hello

import "fmt"

const spanish = "Spanish"
const french = "french"
const englishPrefix = "Hello "
const spanishPrefix = "Hola, "
const frenchPrefix = "Kala, "

func getPrefix(language string) (prefix string) {
	prefix = englishPrefix
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
		//default:
		//	prefix = englishPrefix
	}
	return
}

// Hello go test -v ./hello
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	//if language == spanish {
	//	return spanishPrefix + name
	//}

	//if language == french {
	//	return frenchPrefix + name
	//}

	return getPrefix(language) + name
}
func main() {
	fmt.Println(Hello("", ""))
}
