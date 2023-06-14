package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mapper struct {
	predicate []string
}

func NewMapper() *Mapper {
	var prd []string
	return &Mapper{predicate: prd}
}

func (mp *Mapper) Map() {
	for i := range mp.predicate {
		mp.Emit(mp.predicate[i])
	}
}

func (mp *Mapper) Emit(url string) {
	fmt.Printf("%s\n", url)
}

func parseInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)
	// unixTimestamp, ID пользователя, URL string
	var logs [][]string
	for sc.Scan() {
		fields := strings.Split(sc.Text(), "\t")
		logs = append(logs, fields)
	}
	return logs
}

func main() {
	logs := parseInput()
	mp := NewMapper()

	for _, log := range logs {
		mp.predicate = append(mp.predicate, log[2])
	}
	// fmt.Println(mp.predicate)
	mp.Map()
}
