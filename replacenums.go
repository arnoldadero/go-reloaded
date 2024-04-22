package main

// Import necessary standard library packages
import (
	"strconv" // for converting strings to numbers
	"strings" // for manipulating strings
)

// replaceNums replaces each instance of (hex) and (bin) in the text with their decimal equivalents.
func ReplaceNums(text string) string {
	// Split the input text into individual words, based on whitespace.
	words := strings.Fields(text)

	// Iterate through each word starting from the second position, since a marker can't be at the start.
	for i := 1; i < len(words); i++ {
		// Check if the current word is "(hex)", indicating the previous word is a hexadecimal.
		if words[i] == "(hex)" && i > 0 {
			// Parse the previous word, assumed to be a hex number, converting it to a decimal value.
			decimalValue, err := strconv.ParseInt(words[i-1], 16, 64)
			// Check if the conversion was successful (err is nil).
			if err == nil {
				// Replace the hexadecimal string with its decimal representation as a string.
				words[i-1] = strconv.FormatInt(decimalValue, 10)
				// Empty the current "(hex)" marker to remove it from the result.
				words[i] = ""
			}
		} else if words[i] == "(bin)" && i > 0 { // Similarly, check if the current word is "(bin)".
			// Convert the previous word from binary (base 2) to its decimal (base 10) value.
			decimalValue, err := strconv.ParseInt(words[i-1], 2, 64)
			// If no error occurred, the conversion was successful.
			if err == nil {
				// Replace the binary number with its decimal equivalent in the list of words.
				words[i-1] = strconv.FormatInt(decimalValue, 10)
				// Again, set the "(bin)" marker as an empty string to effectively delete it.
				words[i] = ""
			}
		}
		// No else clause is necessary because we only want to act when we find (hex) or (bin) markers.
	}

	// Join the modified list of words into a single string with spaces between words.
	return strings.Join(words, " ")
}
