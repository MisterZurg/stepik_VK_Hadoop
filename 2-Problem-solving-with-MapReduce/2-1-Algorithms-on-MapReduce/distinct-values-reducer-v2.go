package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Group struct {
	group string
	val   int
}

type ReducerV2 struct {
	// categoty -> count
	h map[string]int
}

func NewReducerV2() *ReducerV2 {
	return &ReducerV2{}
}

func (rd *ReducerV2) Initialize() {
	rd.h = make(map[string]int)
}

func (rd *ReducerV2) Reduce(categories []Group) {
	uniqueCategories := ExcludeDuplicates(categories)
	for _, cat := range uniqueCategories {
		rd.h[cat.group]++
	}
}

func (rd *ReducerV2) Close() {
	for g, v := range rd.h {
		fmt.Printf("%s\t%d\n", g, v)
	}
}

func ExcludeDuplicates(categories []Group) []Group {
	set := make(map[Group]bool)
	var unique []Group
	for _, cat := range categories {
		if _, ok := set[cat]; !ok {
			unique = append(unique, cat)
		}
		set[cat] = true
	}
	return unique
}

func parseInput4ReducerV2() []Group {
	sc := bufio.NewScanner(os.Stdin)

	var valueGroups []Group
	for sc.Scan() {
		txt := sc.Text()
		words := strings.Split(txt, "\t")
		v, _ := strconv.Atoi(words[0])
		valueGroups = append(valueGroups, Group{group: words[1], val: v})
	}
	return valueGroups
}

func main() {
	valueGroups := parseInput4ReducerV2()

	reducer := NewReducerV2()
	reducer.Initialize()
	reducer.Reduce(valueGroups)
	reducer.Close()
}
