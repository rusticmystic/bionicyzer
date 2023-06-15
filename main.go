// / 2>/dev/null ; gorun "$0" "$@" ; exit $?
package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/bitfield/script"
)

func bionic(text string) string {
	words := strings.Fields(text)
	transformedWords := make([]string, len(words))
	for i, word := range words {
		transformedWords[i] = transformWord(word)
	}
	return strings.Join(transformedWords, " ")
}

func boldChars(text string, num int) string {
	n := int(num)

	if n == 1 && len(text) == 1 || n == 0 && len(text) == 1 {
		// if string has length 1
		// n can be 0, for single char like : a
		// n can be 1, for 2 character strings like: an
		return fmt.Sprintf("**%s**", text)
	}

	// Check if first character is alphanumeric
	if !unicode.IsLetter(rune(text[0])) && !unicode.IsNumber(rune(text[0])) {
		// strings starting with `,",( are unchanged)
		return text
	}

	return fmt.Sprintf("**%s**%s", text[:n], text[n:])
}

func transformWord(s string) string {
	textLen := len(s)

	if textLen/2 > 7 {
		// bold at max 5 characters
		return boldChars(s, 7)
	}
	if textLen%2 == 0 {
		// even
		return boldChars(s, textLen/2)
	}

	// odd
	return boldChars(s, (textLen+1)/2)
}

func main() {
	// read from stdin and write to stdin
	script.Stdin().FilterScan(func(line string, w io.Writer) {
		fmt.Fprintf(w, "%s\n", bionic(line))
	}).Stdout()
}
