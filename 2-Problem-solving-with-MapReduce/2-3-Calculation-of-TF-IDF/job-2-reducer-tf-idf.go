package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	t := make([][]string, 0)
	k := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		text := regexp.MustCompile(`\w+`).FindAllString(line, -1)
		k = append(k, text[:1])
		t = append(t, text[:3])
	}

	for i := 0; i < len(t); i++ {
		count := countOccurrences(k, k[i])
		fmt.Printf("%s#%s\t%s\t%d\n", t[i][0], t[i][1], t[i][2], count)
	}
}

func countOccurrences(slice [][]string, element []string) int {
	count := 0
	for _, item := range slice {
		if areSlicesEqual(item, element) {
			count++
		}
	}
	return count
}

func areSlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
