package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reducer struct{}

func NewReducerDV1() *Reducer {
	return &Reducer{}
}

func (rd *Reducer) Reduce(mapOut []string) {
	rd.Emit(mapOut[0], mapOut[1], 0)
}

func (mp *Reducer) Emit(g, f string, cnt int) {
	fmt.Printf("%s,%s\n", g, f)
}

func parseMapOutput() [][]string {
	sc := bufio.NewScanner(os.Stdin)

	var valueGroups [][]string
	procecedGroup := make(map[string]bool)
	for sc.Scan() {
		txt := sc.Text()
		words := strings.Split(txt, "\t")

		if _, ok := procecedGroup[words[0]]; !ok {
			groups := strings.Split(words[0], ",")
			var categ []string
			categ = append(categ, groups...)
			categ = append(categ, words[1])

			valueGroups = append(valueGroups, categ)

			procecedGroup[words[0]] = true
		}
	}
	return valueGroups
}

func main() {
	mapOutput := parseMapOutput()
	rd := NewReducerDV1()

	for _, out := range mapOutput {
		rd.Reduce(out)
	}
}
