package functions

import (
	"fmt"
	"os"
	"strings"
)

func BuildFontMap(fileName string) (map[rune][]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("oops, an error has occured while trying to parse file: %w ", err)
	}

	lines := strings.Split(string(data), "\n")
	asciiMaps := make(map[rune][]string)

	spaceChar := rune(32)

	for i := 0; i+8 <= len(lines); i += 9 {
		asciiMaps[spaceChar] = lines[i+1 : i+9]
		spaceChar++
	}

	if len(asciiMaps) == 0 {
		return nil, fmt.Errorf("oops! failed to build map")
	}
	return asciiMaps, nil
}
