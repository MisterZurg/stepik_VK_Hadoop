package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name    string
	queries []string
	urls    []string
}

func (p *Person) fmtPrintln() {
	if len(p.queries) == 0 || len(p.urls) == 0 {
		fmt.Sprintf("")
	}
	for i := 0; i < len(p.queries); i++ {
		for j := 0; j < len(p.urls); j++ {
			fmt.Println(p.name + "\t" + p.queries[i] + "\t" + p.urls[j])
		}
	}
}

func parseInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)

	var requests [][]string
	for sc.Scan() {
		fields := strings.Split(sc.Text(), "\t")
		markerquery := strings.Split(fields[1], ":")
		requests = append(requests, []string{fields[0], markerquery[0], markerquery[1]})
	}
	return requests
}

func fillPersons(requests [][]string) []Person {
	persons := []Person{}
	personsCache := make(map[string]int)
	idx := 0
	for _, request := range requests {
		if _, ok := personsCache[request[0]]; !ok {
			personsCache[request[0]] = idx
			qrs := []string{}
			urls := []string{}
			persona := Person{name: request[0], queries: qrs, urls: urls}
			persons = append(persons, persona)
			idx++
		}
		currIdx := personsCache[request[0]]
		if request[1] == "query" {
			persons[currIdx].queries = append(persons[currIdx].queries, request[2])
		} else if request[1] == "url" {
			persons[currIdx].urls = append(persons[currIdx].urls, request[2])
		}

	}
	return persons
}

func main() {
	requests := parseInput()
	persons := fillPersons(requests)

	for _, person := range persons {
		person.fmtPrintln()
	}
}
