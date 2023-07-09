package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		parsedLine := strings.SplitN(line, ":", 2)
		words := strings.FieldsFunc(parsedLine[1], func(r rune) bool {
			return !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'))
		})
		for _, word := range words {
			fmt.Printf("%s#%s\t1\n", word, parsedLine[0])
		}
	}
}
