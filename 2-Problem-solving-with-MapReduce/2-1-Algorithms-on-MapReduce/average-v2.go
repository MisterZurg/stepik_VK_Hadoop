package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Combiner struct {
	average map[string][]int
	order   []string
}

func NewCombiner() *Combiner {
	average := make(map[string][]int)
	return &Combiner{
		average: average,
		order:   []string{},
	}
}

func (cb *Combiner) Combine(ua []userActivityV2) {
	// average := make(map[string][]int)
	//var order []string
	for _, v := range ua {
		if _, ok := cb.average[v.url]; !ok {
			cb.order = append(cb.order, v.url)
			cb.average[v.url] = make([]int, 2)
		}
		cb.average[v.url][0] += v.seconds
		cb.average[v.url][1] += v.cnt
	}
}

func (cb *Combiner) Emit() {
	for _, k := range cb.order {
		fmt.Printf("%s\t%d;%d\n", k, cb.average[k][0], cb.average[k][1])
	}
}

type userActivityV2 struct {
	url     string
	seconds int
	cnt     int
}

func parseInputCombiner() []userActivityV2 {
	var lines []userActivityV2
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		raw := strings.Split(line, "\t")

		usac := processRawActivityCombiner(raw)

		lines = append(lines, usac)
	}
	return lines
}

func processRawActivityCombiner(raw []string) userActivityV2 {
	nums := strings.Split(raw[1], ";")
	sec, err := strconv.Atoi(nums[0])
	if err != nil {
		fmt.Println("Cannot convert sec=%v", nums[0])
		return userActivityV2{}
	}
	return userActivityV2{url: raw[0], seconds: sec, cnt: 1}
}

func main() {
	ua := parseInputCombiner()
	cmb := NewCombiner()
	cmb.Combine(ua)
	cmb.Emit()
}
