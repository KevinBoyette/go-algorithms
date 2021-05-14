package str

import (
	"strings"
	"unicode"
)

func ToCamelCase(word string) string {
	if len(word) < 1 {
		return word
	}

	stringBuilder := strings.Builder{}
	stringBuilder.WriteByte(word[0])
	for i := 1; i < len(word); i++ {
		char := word[i]
		if char == '_' || char == '-' {
			stringBuilder.WriteRune(unicode.ToUpper(rune(word[i+1])))
			i++
			continue
		}
		stringBuilder.WriteByte(char)
	}
	return stringBuilder.String()
}
