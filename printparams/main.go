package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:] // Exclude the first argument, which is the program name

	for _, arg := range arguments {
		printString(arg)
		printChar('\n')
	}
}

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func printChar(char rune) {
	z01.PrintRune(char)
}
