package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var last string
	sum := 1

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		wordDoc := parts[0]

		if last == "" {
			last = wordDoc
		} else if last != wordDoc {
			word, doc := strings.Split(last, "#")[0], strings.Split(last, "#")[1]
			fmt.Printf("%s\t%s\t%d\n", word, doc, sum)
			last = wordDoc
			sum = 1
		} else {
			sum++
		}
	}

	word, doc := strings.Split(last, "#")[0], strings.Split(last, "#")[1]
	fmt.Printf("%s\t%s\t%d\n", word, doc, sum)
}
