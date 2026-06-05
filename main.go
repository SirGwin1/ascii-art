package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Validate arguments
	if err := functions.ValidateArgs(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Argument Error: %v\n", err)
		os.Exit(1)
	}

	// Build font map
	asciiMap, err := functions.BuildFontMap("shadow.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Font map error: %v\n", err)
		os.Exit(1)
	}

	// Get input
	input := strings.Join(os.Args[1:], " ")

	// Render ASCII
	result, err := functions.RenderAscii(input, asciiMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Render Error: %v\n", err)
		os.Exit(1)
	}

	// Print result
	if err := functions.PrintResult(result); err != nil {
		fmt.Fprintf(os.Stderr, "Print Error: %v\n", err)
		os.Exit(1)
	}
}
