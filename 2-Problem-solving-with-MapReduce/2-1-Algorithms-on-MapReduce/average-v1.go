package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type userActivity struct {
	url     string
	seconds int
}

func processRawActivity(raw []string) userActivity {
	sec, err := strconv.Atoi(raw[1])
	if err != nil {
		fmt.Println("Cannot convert sec=%v", raw[1])
		return userActivity{}
	}
	return userActivity{url: raw[0], seconds: sec}
}

type MapperAvgV1 interface {
	Map(t string, r int)
}

type ReducerAvgV1 interface {
	Reduce(t string, rs []int)
}

type myMapperAvgV1 []userActivity

type myReducer struct {
	uaFromMapper []userActivity
	average      map[string][]int
}

func NewReducer() *myReducer {
	average := make(map[string][]int)
	return &myReducer{uaFromMapper: []userActivity{}, average: average}
}

func (rd myReducer) Reduce() {
	order := []string{}
	for _, ua := range rd.uaFromMapper {
		if _, ok := rd.average[ua.url]; !ok {
			order = append(order, ua.url)
			rd.average[ua.url] = make([]int, 2)
		}
		rd.average[ua.url][0] += ua.seconds
		rd.average[ua.url][1]++
	}
	rd.Emit(order)
}
func (rd myReducer) Emit(order []string) {
	for _, k := range order {
		fmt.Printf("%s\t%d\n", k, rd.average[k][0]/rd.average[k][1])
	}
}

func parseInputReducer() []userActivity {
	lines := []userActivity{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		raw := strings.Split(line, "\t")

		usac := processRawActivity(raw)

		lines = append(lines, usac)
	}
	return lines
}

func main() {
	procUserActivity := parseInputReducer()
	fmt.Println(procUserActivity)
	var mp myMapperAvgV1 = procUserActivity
	rd := NewReducer()
	rd.uaFromMapper = mp
	fmt.Println(rd)
	rd.Reduce()
}
