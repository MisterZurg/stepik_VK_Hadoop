package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = math.MaxInt

func parseMapperInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)
	lines := [][]string{}
	for sc.Scan() {
		words := strings.Split(sc.Text(), "\t")
		lines = append(lines, words)
	}
	return lines
}

func buildGraph(in [][]string) {
	// map Node -> toNodes -> wight
	// 1	0	{2,3,4}
	// print()
	// 1	0	{2,3,4}
	// 2    1   {}
	// 3    1   {}
	// 4    1   {}
	// g := make(map[])
	dists := make([]int, len(in))
	for i := range dists {
		dists[i] = INF
	}

	for _, document := range in {
		fmt.Printf("%s\t%s\t%s\n", document[0], document[1], document[2])
		node, _ := strconv.Atoi(document[0])
		if document[1] != "INF" {
			dists[node-1], _ = strconv.Atoi(document[1])
		}

		// Replace {}
		clearAdj := document[2][1 : len(document[2])-1]

		if len(clearAdj) > 0 {
			toNodes := strings.Split(clearAdj, ",")
			for _, toNode := range toNodes {
				knowDist := "INF"
				if dists[node-1] != INF {
					knowDist = strconv.Itoa(dists[node-1] + 1)
				}
				fmt.Printf("%s\t%s\t{}\n", toNode, knowDist)
			}
		}
	}
}

func main() {
	// put your code here
	in := parseMapperInput()
	buildGraph(in)
}
