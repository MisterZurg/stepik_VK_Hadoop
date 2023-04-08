package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MapperV1 interface {
	Map(line string)
}

type myMapper []string

func (mp *myMapper) Map(line string) {
	// term : count
	h := make(map[string]int)
	terms := []string{} // Для прохода в порядке добавления в мапку
	for _, term := range strings.Split(line, " ") {
		if _, ok := h[term]; !ok {
			terms = append(terms, term)
		}
		h[term]++
	}

	for _, term := range terms {
		Emit(term, h[term])
	}
}

func Emit(term string, count int) {
	fmt.Printf("%s\t%d\n", term, count)
}

func parseInput() []string {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func main() {
	lines := parseInput()
	var mp myMapper
	for _, line := range lines {
		mp.Map(line)
	}
}
