package main

// Import required packages
import (
	"strconv" // For string conversion to numbers
	"strings" // For string manipulation functions
)

// ApplyCasing updates the case of words just before the (up), (low), or (cap) markers.
// If a number appears next to the marker, it capitalizes the specified number of preceding words.
func ApplyCasing(input string) string {
	// Split the input string into a slice of words
	words := strings.Fields(input)
	// Iterate over the words slice in reverse
	for i := len(words) - 1; i >= 0; i-- {
		// Declare a variable to hold a function that takes a string and returns a string
		var convertFunc func(string) string

		// Identify which marker is present and apply the corresponding case function
		switch {
		// Check for the presence of the '(up', '(low', or '(cap' marker at the beginning of the current word
		case strings.HasPrefix(words[i], "(up"), strings.HasPrefix(words[i], "(low"), strings.HasPrefix(words[i], "(cap"):
			// Remove the parentheses from the marker, leaving the action and possibly a comma
			words[i] = strings.Trim(words[i], "()")
			// Split the word into the action part and the number part if present
			params := strings.SplitN(words[i], ",", 2)

			// Assign the appropriate function to convertFunc based on the action
			switch params[0] {
			case "up":
				convertFunc = strings.ToUpper // Convert to uppercase
			case "low":
				convertFunc = strings.ToLower // Convert to lowercase
			case "cap":
				convertFunc = strings.Title // Capitalize
			}

			// Extract the number of words to apply the casing to
			numberString := strings.Trim(words[i+1], ")")
			// Try converting numberString to an integer, defaulting to 1 if it fails
			n, err := strconv.Atoi(numberString)
			if err != nil {
				n = 1 // Default to converting just one word if no number is specified
			} else {
				words[i+1] = "" // Clear the array element that contained the number
			}

			// Loop over and apply the conversion function to the number of words specified
			for c := 0; c < n && i-c-1 >= 0; c++ {
				prevIndex := i - c - 1  // Calculate the index of the word to convert
				if convertFunc != nil { // If a conversion function is set
					words[prevIndex] = convertFunc(words[prevIndex]) // Apply the conversion function
				}
			}
			words[i] = "" // Clear the marker array element
		}
	}

	// Initialize a slice to store non-empty words after processing
	var cleanWords []string
	// Iterate over the words and append non-empty words to cleanWords
	for _, w := range words {
		if w == "" { // If the word is empty, skip it
			continue
		}
		cleanWords = append(cleanWords, w) // Append the non-empty word to cleanWords
	}

	// Join the words in cleanWords with spaces to form the processed string
	return strings.Join(cleanWords, " ")
}
