package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:] // Exclude the first argument, which is the program name

	// Reverse the order of the arguments
	reversedArguments := reverseStringSlice(arguments)

	// Print the reversed arguments
	for _, arg := range reversedArguments {
		printString(arg)
		printChar('\n')
	}
}

func reverseStringSlice(slice []string) []string {
	length := len(slice)
	reversed := make([]string, length)
	for i, item := range slice {
		reversed[length-i-1] = item
	}
	return reversed
}

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func printChar(char rune) {
	z01.PrintRune(char)
}
