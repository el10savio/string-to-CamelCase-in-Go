package main

import (
	"fmt"
	"strings"
)

func toCamelCase(input string) string {
	// 1. Trim the input string

	inputStringTrimmed := strings.Trim(input, " ")

	var result []string

	// 2. Iterate through all input string characters
	for i := 0; i < len(inputStringTrimmed); i++ {
		character := inputStringTrimmed[i]
		// 3. Make first character lowercase
		if i == 0 {
			result = append(result, strings.ToLower(string(character)))
		} else if string(character) == " " {
			// 4. Find occurrences of space
			// 5. Convert character next to space to Uppercase
			// 6. Remove all occurrences of spaces
			i = i + 1
			result = append(result, strings.ToUpper(string(inputStringTrimmed[i])))
		} else {
			result = append(result, string(character))
		}
	}

	resultString := strings.Join(result, "")
	return resultString
}

func main() {
	inputString := "HEllo world"
	fmt.Println(toCamelCase(inputString))
}
