package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:] // Exclude the first argument, which is the program name

	// Find the smallest argument and print it, then remove it from the list
	for len(arguments) > 0 {
		smallestIdx := findSmallestIndex(arguments)
		printString(arguments[smallestIdx])
		printChar('\n')
		arguments = removeElement(arguments, smallestIdx)
	}
}

func findSmallestIndex(slice []string) int {
	smallestIdx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[smallestIdx] {
			smallestIdx = i
		}
	}
	return smallestIdx
}

func removeElement(slice []string, indexToRemove int) []string {
	return append(slice[:indexToRemove], slice[indexToRemove+1:]...)
}

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func printChar(char rune) {
	z01.PrintRune(char)
}
