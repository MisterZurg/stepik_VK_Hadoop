package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)
	var pairs [][]string
	for sc.Scan() {
		fields := strings.Split(sc.Text(), "\t")
		pairs = append(pairs, fields)
	}
	return pairs
}

func main() {
	pairs := parseInput()
	lastKey := ""
	size := 0

	for _, pair := range pairs {
		if lastKey == "" || lastKey == pair[0] {
			size++
		} else {
			if size == 1 {
				fmt.Println(lastKey)
			}
			size = 1
		}
		lastKey = pair[0]
	}
	if size == 1 {
		fmt.Println(lastKey)
	}
}
