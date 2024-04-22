package main

// Import strings package for string manipulation
import (
	"strings"
)

// FormatPunctuation corrects the spacing around punctuation marks.
func FormatPunctuation(input string) string {
	// Create a Builder to efficiently construct the formatted string.
	var builder strings.Builder

	// Iterate through each character in the input string using range.
	for i, r := range input {
		// Convert the rune 'r' to a string representing the current character.
		currentChar := string(r)
		// Initialize previousChar which will hold the character before currentChar.
		previousChar := " " // Default to a space character.
		// Assign the actual previous character if currentChar is not the first one.
		if i > 0 {
			previousChar = string(input[i-1])
		}

		// Check if the current character is a punctuation character we're interested in.
		if strings.Contains(".!,?:;", currentChar) {
			// If it's not leading group punctuation (e.g., '...') and we're not looking at the first character,
			// the punctuation should immediately follow the previous non-space character.
			if i == 0 || !strings.Contains(".!?", previousChar) {
				if previousChar == " " {
					// If the previous character is a space, remove the last space from the Builder.
					temp := builder.String()                // Store current Builder state in temp.
					builder.Reset()                         // Reset the Builder to empty it.
					builder.WriteString(temp[:len(temp)-1]) // Write back the Builder string without the last character to remove the space.
				}
				// Add the punctuation character to the Builder.
				builder.WriteString(currentChar)
				// If the next character exists and is not punctuation, ensure there's a space afterward.
				if (i+1) < len(input) && !strings.Contains(".!,?:;", string(input[i+1])) {
					builder.WriteRune(' ') // Write a space.
				}
			} else {
				// For cases where group punctuation is detected...
				builder.WriteString(currentChar) // Add the punctuation character to the Builder as is.
			}
		} else {
			// If the current character is not a punctuation mark, add it to the Builder.
			builder.WriteRune(r)
		}
	}

	// Once the loop is done, convert the Builder's contents to a string and return it.
	// This is the resulting formatted text with the corrected punctuation spacing.
	return builder.String()
}
