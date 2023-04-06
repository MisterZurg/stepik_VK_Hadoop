package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func Mapper(txt []string) { //([]string, map[string]int){
	for i := 0; i < len(txt); i++ {
		// wordCount := make(map[string]int)
		words := []string{}

		for _, word := range strings.Split(txt[i], " ") {
			//if _, ok := wordCount[word]; !ok {
			words = append(words, word)
			//}
			//wordCount[word]++
		}

		for _, word := range words {
			fmt.Printf("%s\t1\n", word) //, wordCount[word])
		}
	}
}

func main() {
	// put your code here
	lines := parseInput()
	Mapper(lines)
}
