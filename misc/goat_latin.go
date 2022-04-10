package misc

import "strings"

// ToGoatLatin a string (as seen in interviews and codewars).
func ToGoatLatin(s string) string {
	vowels := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	words := strings.Split(s, " ")
	a := "a"
	for i := 0; i < len(words); i++ {
		if !vowels[words[i][0]] {
			words[i] = words[i][1:] + string(words[i][0])
		}
		words[i] += "ma" + a
		a += "a"
	}
	return strings.Join(words, " ")
}
