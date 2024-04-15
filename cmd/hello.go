package main

import "fmt"

const (
    spanish = "Spanish"
    french = "French"
    englishPrefix = "Hello, "
    frenchPrefix = "Bonjour, "
    spanishPrefix = "Hola, "
)

func main() {
    fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {

    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

    switch language {
    case french:
        prefix = frenchPrefix

    case spanish:
        prefix = spanishPrefix
        
    default:
        prefix = englishPrefix
    }

    return
}