package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mapper struct {
	predicate map[string][]string
}

func NewMapper() *Mapper {
	prd := make(map[string][]string)
	return &Mapper{predicate: prd}
}

func (mp *Mapper) Map(key string, timestamps map[string]string) {
	if v, ok := mp.predicate[key]; ok {
		for _, timestamp := range v {
			mp.Emit(timestamp, timestamps)
		}
	}
}

func (mp *Mapper) Emit(key string, timestamps map[string]string) {
	fmt.Printf("%s\t%s", key, timestamps[key])
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
	timestamps := make(map[string]string)

	for _, log := range logs {
		timestamps[log[0]] = fmt.Sprintf("%s\t%s\n", log[1], log[2])
		mp.predicate[log[1]] = append(mp.predicate[log[1]], log[0])
	}

	mp.Map("user10", timestamps)
}
