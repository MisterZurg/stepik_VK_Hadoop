package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MapperV2 interface {
	Initialize()
	Map(line string)
	Close()
}

type myMapperV2 struct {
	lines []string
	h     map[string]int
}

func (mp *myMapperV2) Initialize() {
	mp.h = make(map[string]int)
}

func (mp *myMapperV2) Map(line string) {
	for _, term := range strings.Split(line, " ") {
		//if _, ok := mp.h[term]; !ok {
		//	terms = append(terms, term)
		//}
		mp.h[term]++
	}
}

func (mp *myMapperV2) Close() {
	for term, count := range mp.h {
		Emit(term, count)
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
	var mp myMapperV2
	mp.Initialize()

	for _, line := range lines {
		mp.Map(line)
	}

	mp.Close()
}
