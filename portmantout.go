package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Read the input
	lines := getStdin()

	// Number of words
	n := len(lines)
	if n < 2 {
		fmt.Println(-1)
		return
	}

	// Store the words
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = lines[i]
	}

	// Find the longest coined word
	maxLength := -1

	// Iterate over all pairs of words
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && len(words[i]) >= 2 && len(words[j]) >= 2 && strings.HasSuffix(words[i], words[j][:2]) {
				coinedWord := words[i] + words[j][2:]
				if len(coinedWord) > maxLength {
					maxLength = len(coinedWord)
				}
			}
		}
	}

	// Output the result
	if maxLength == -1 {
		fmt.Println(-1)
	} else {
		fmt.Println(maxLength)
	}
}

func getStdin() []string {
	stdin, _ := ioutil.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
