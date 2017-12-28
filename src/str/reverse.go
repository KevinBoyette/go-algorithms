package str

// Reverse takes a word or string and reverses the word
// Example:
// 		Reverse("hello") -> olleh
// 		Reverse("racecar") -> racecar
// 		Reverse("") -> ""
func Reverse(word string) string {
	characters := []rune(word)
	wordLen := len(characters) - 1
	if wordLen <= 1 {
		return word
	}
	for i, j := 0, wordLen; i < j; i, j = i+1, j-1 {
		characters[i], characters[j] = characters[j], characters[i]
	}
	return string(characters)
}
