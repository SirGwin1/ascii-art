package functions

import (
	"fmt"
	"strings"
)

func RenderAscii(input string, asciiMap map[rune][]string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input text cannot be empty")
	}

	if len(asciiMap) == 0 {
		return "", fmt.Errorf("font map is empty")
	}

	var result strings.Builder

	for line := 0; line < 8; line++ {
		for _, ch := range input {

			val, ok := asciiMap[ch]
			if !ok {
				return "", fmt.Errorf("unsupported character: %q", ch)
			}

			result.WriteString(val[line])
		}
		result.WriteString("\n")
	}

	return result.String(), nil
}