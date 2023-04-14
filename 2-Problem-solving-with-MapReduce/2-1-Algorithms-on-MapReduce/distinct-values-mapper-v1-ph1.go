package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mapper struct{}

func NewMapper() *Mapper {
	return &Mapper{}
}

// Map принимает на вход строку, содержащую значение и через табуляцию список групп, разделенных запятой.
func (mp *Mapper) Map(valueGroups []string) {
	for i := 1; i < len(valueGroups); i++ {
		mp.Emit(valueGroups[0], valueGroups[i], 1)
	}
}

func (mp *Mapper) Emit(g, f string, cnt int) {
	fmt.Printf("%s,%s\t%d", g, f, cnt)
}

func parseMapInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)

	var valueGroups [][]string

	for sc.Scan() {
		txt := sc.Text()
		words := strings.Split(txt, "\t")
		groups := strings.Split(words[1], ",")
		categ := []string{words[0]}
		for _, g := range groups {
			categ = append(categ, g)
		}
		valueGroups = append(valueGroups, categ)
	}
	return valueGroups
}

func main() {
	valueGroups := parseMapInput()

	mp := NewMapper()
	for _, vg := range valueGroups {
		mp.Map(vg)
	}
}
