package main

import (
	"log"
	"strings"
	"unicode"
)

// ToSnakeCase converts a PascalCase or camelCase string to snake_case.
func ToSnakeCase(str string) string {
	var result strings.Builder
	for i, char := range str {
		if unicode.IsUpper(char) && i > 0 {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(char))
	}
	return result.String()
}

func main() {
	log.Println("executing poc")
}
