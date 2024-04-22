package main

import (
	"strings"
)

// AdjustArticle adjusts the article 'a' to 'an' where appropriate in a sentence.
func AdjustArticle(input string) string {
	// Split the input into words for analysis.
	words := strings.Fields(input)
	vowels := "aeiou" // String containing vowels for easy checking.

	for i, word := range words {
		if word == "a" || word == "A" { // Check if the word is an article 'a' or 'A'.
			if i+1 < len(words) { // Ensure there is a next word.
				nextWordFirstLetter := strings.ToLower(string(words[i+1][0])) // Get the first letter of the next word, in lowercase.

				if strings.ContainsRune(vowels, rune(nextWordFirstLetter[0])) || nextWordFirstLetter == "h" {
					// If the next word starts with a vowel or 'h', adjust the article.
					if word == "a" {
						words[i] = "an" // Change 'a' to 'an' for lowercase.
					} else {
						words[i] = "An" // Change 'A' to 'An' for uppercase.
					}
				}
			}
		}
	}
	// Join the words back into a single string.
	return strings.Join(words, " ")
}
