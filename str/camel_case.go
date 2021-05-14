package str

import (
	"strings"
	"unicode"
)

// ToCamelCase converts a word that is in either kebob-case or
// snake_case into camelCase.
func ToCamelCase(word string) string {
	if len(word) < 1 {
		return word
	}

	stringBuilder := strings.Builder{}
	// We could loop through the word and count the hyphens and
	// underscores.  However, for most inputs, we can grow a little
	// more than we should.
	stringBuilder.Grow(len(word))

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
