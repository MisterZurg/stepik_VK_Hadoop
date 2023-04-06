package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func Reducer(txt []string) {
	wordCount := make(map[string]int)
	words := []string{}
	for i := 0; i < len(txt); i++ {
		line := strings.Split(txt[i], "\t")
		word := line[0]
		cnt, _ := strconv.Atoi(line[1])
		if _, ok := wordCount[word]; !ok {
			words = append(words, word)
		}
			wordCount[word] += cnt
		}

		for _, word := range words {
			fmt.Printf("%s\t%d\n", word , wordCount[word])
		}
}

func main() {
	// put your code here
	lines := parseInput()
	Reducer(lines)
}
