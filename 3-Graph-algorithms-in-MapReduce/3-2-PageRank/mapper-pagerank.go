package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseMapperInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)
	in := [][]string{}
	for sc.Scan() {
		line := sc.Text()
		in = append(in, strings.Split(line, "\t"))
	}

	return in
}

func mapperPageRank(in [][]string) {

	for _, req := range in {
		fmt.Printf("%s\t%s\t%s\n", req[0], req[1], req[2])
		PageRank, _ := strconv.ParseFloat(req[1], 64)

		clearAdj := req[2][1 : len(req[2])-1]
		toNodes := strings.Split(clearAdj, ",")
		for _, val := range toNodes {
			fmt.Printf("%s\t%.3f\t{}\n", val, PageRank/float64(len(toNodes)))
		}
	}
}

func main() {
	in := parseMapperInput()
	mapperPageRank(in)
}
