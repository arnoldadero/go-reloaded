# TextModifier

TextModifier is a highly efficient text processing tool designed to automate text transformations in Go. It functions by receiving a source text file with its content requiring modifications, and outputs the modified text to a target file.

## Features
This program features a robust set of text transformations, including:
- **Hexadecimal to Decimal**: Converts words representing hex numbers directly preceding `(hex)` into their decimal counterpart.
- **Binary to Decimal**: Converts words representing binary numbers directly preceding `(bin)` into their decimal counterpart.
- **Uppercase Conversion**: Transforms the word before `(up)` into uppercase, or multiple words in case of `(up, <number>)`.
- **Lowercase Conversion**: Transforms the word before `(low)` into lowercase, or multiple words in case of `(low, <number>)`.
- **Capitalization**: Capitalizes the word before `(cap)` or multiple words when `(cap, <number>)` is used.
- **Punctuation Formatting**: Corrects misplaced spacing around punctuation marks `.`, `,`, `!`, `?`, `:`, and `;`.
- **Grouped Punctuation Formatting**: Adjusts spacing for punctuations groups like `...` or `!?`.
- **Apostrophe Formatting**: Ensures apostrophes are correctly positioned around single or multiple words.
- **Indefinite Article Correction**: Changes `a` to `an` when followed by a word starting with a vowel or `h`.

## Installation
To set up TextModifier on your system, ensure you have Go installed and follow these steps:
1. Clone the repository to your local machine.
2. Navigate to the cloned directory.

## Usage
Execute the program by running:
```bash
go run . sample.txt result.txt

