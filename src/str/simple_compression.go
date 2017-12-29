package str

import (
	"bytes"
	"strconv"
	"strings"
)

// Takes a sequence as input and returns
// a compressed form
// Example:
// SimpleCompression("aaaabbbccdaa") -> a4b3c2d1a2
func SimpleCompression(sequence string) string {
	sequenceLen := len(sequence)
	sequence = strings.ToLower(sequence)
	if sequenceLen < 1 {
		return sequence
	} else if sequenceLen == 1 {
		return sequence + "1"
	}
	currentChar := sequence[0]
	count := 1
	var list bytes.Buffer
	for index := 1; index < sequenceLen; index++ {
		if sequence[index] == currentChar {
			count++
		} else {
			list.WriteString(string(currentChar))
			list.WriteString(strconv.Itoa(count))
			currentChar = sequence[index]
			count = 1
		}
	}
	list.WriteString(string(currentChar))
	list.WriteString(strconv.Itoa(count))

	return list.String()
}
