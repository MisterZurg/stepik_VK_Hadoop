package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MapperPh2 struct{}

func NewMapperPh2() *MapperPh2 {
	return &MapperPh2{}
}

func (mp *MapperPh2) Map(g string) {
	mp.Emit(g, 1)
}

func (mp *MapperPh2) Emit(g string, cnt int) {
	fmt.Printf("%s\t%d\n", g, cnt)
}

func parseInput4MapperPh2() []string {
	sc := bufio.NewScanner(os.Stdin)

	var valueGroups []string
	for sc.Scan() {
		txt := sc.Text()
		words := strings.Split(txt, ",")
		valueGroups = append(valueGroups, words[1])
	}
	return valueGroups
}

func main() {
	valueGroups := parseInput4MapperPh2()

	mp := NewMapperPh2()
	for _, vg := range valueGroups {
		mp.Map(vg)
	}
}
