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

func parseReducerInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)
	lines := [][]string{}
	for sc.Scan() {
		words := strings.Split(sc.Text(), "\t")
		lines = append(lines, words)
	}
	return lines
}

func graphProcessing(in [][]string) {
	documents := make(map[string][2]string)

	reqId := make(map[int]string)
	id := 0
	for _, doc := range in {
		/*
		   1	0	{2,3,4}
		   10	INF	{}
		   10	INF	{}
		   ...
		*/
		node := doc[0]
		dist := doc[1]
		adj := doc[2]

		if _, ok := documents[node]; !ok {
			reqId[id] = node
			documents[node] = [2]string{dist, adj}
			id++
		}

		prevDist := documents[node][0]
		prevAdj := documents[node][1]

		if prevDist == "INF" && prevAdj == "{}" {
			documents[node] = [2]string{dist, adj}
		} else if prevDist == "INF" {
			prevVal, _ := strconv.Atoi(prevDist)
			curVal, _ := strconv.Atoi(dist)

			if prevVal > curVal {
				documents[node] = [2]string{dist, prevAdj}
			}
		}

		prevDist = documents[node][0]
		prevAdj = documents[node][1]

		if prevAdj == "{}" {
			documents[node] = [2]string{prevDist, adj}
		}

	}
	for i := 0; i <= id; i++ {
		v := documents[reqId[i]]
		fmt.Printf("%s\t%s\t%s\n", reqId[i], v[0], v[1])
	}
}

func main() {
	// put your code here
	in := parseReducerInput()
	graphProcessing(in)
}
