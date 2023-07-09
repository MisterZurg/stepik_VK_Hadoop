package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		reversedLine := reverseString(line)
		trimmedLine := strings.TrimSpace(reversedLine)
		replacedLine := strings.Replace(trimmedLine, "\t", ";", 1)
		finalLine := reverseString(replacedLine) + ";1"
		fmt.Println(finalLine)
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
