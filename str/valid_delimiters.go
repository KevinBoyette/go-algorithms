package str

// ValidDelimiters takes a word as input and computes
// whether the delimiters are balanced or not.
//
// Examples:
//
//	ValidDelimiters("()") -> true
//	ValidDelimiters("[](){}") -> true
//	ValidDelimiters("[}") -> false
func ValidDelimiters(word string) bool {
	var left []rune
	delims := map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}
	for _, s := range word {
		if _, ok := delims[s]; ok {
			left = append(left, s)
		} else if len(left) == 0 || delims[left[len(left)-1]] != s {
			return false
		} else {
			left = left[:len(left)-1]
		}
	}
	return len(left) == 0
}
