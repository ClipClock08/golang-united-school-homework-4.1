package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	var outputSB = strings.Builder{}
	for i := len([]rune(input)) - 1; i >= 0; i-- {
		outputSB.WriteRune([]rune(input)[i])
	}
	output = outputSB.String()
	return output
}
