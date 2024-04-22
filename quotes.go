package main

// Import the strings package for string manipulation functions.
import "strings"

// FormatQuotes finds single quote characters and ensures they are correctly placed next to the quoted words.
func FormatQuotes(input string) string {
	// Split the input text into a slice of words.
	words := strings.Fields(input)
	// Initialize a flag that indicates whether we're looking for the start of a quote.
	waitingForStartQuote := true

	// Iterate over the words.
	for i, w := range words {
		// Check if the current word is a single quote.
		if w == "'" {
			// If we are waiting for the start quote...
			if waitingForStartQuote {
				// Clear the current quote mark as we will prepend it to the next word.
				words[i] = ""
				// Prepend the single quote to the next word.
				words[i+1] = "'" + words[i+1]

				// We've found the start quote, now we're looking for the end quote.
				waitingForStartQuote = false
			} else {
				// If we were not waiting for the start quote, this must be the end quote.
				// Clear the current end quote mark as we will append it to the previous word.
				words[i] = ""
				// Append the single quote to the previous word.
				words[i-1] = words[i-1] + "'"

				// We've found the end quote, now we'll look for the next start quote.
				waitingForStartQuote = true
			}
		}
	}

	// Initialize an empty slice to store non-empty words after processing.
	var cleanWords []string
	// Loop through the slice words to remove any empty strings resulting from the quote formatting.
	for _, w := range words {
		// If the current word is not empty, append it to cleanWords.
		if w != "" {
			cleanWords = append(cleanWords, w)
		}
	}

	// Join all the words in cleanWords with spaces and return the resulting string, which is the input string with correctly formatted quotes.
	return strings.Join(cleanWords, " ")
}
