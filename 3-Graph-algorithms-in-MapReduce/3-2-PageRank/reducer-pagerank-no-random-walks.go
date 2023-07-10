package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	d := make(map[string]float64)
	ag := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		node := parts[0]
		rank, _ := strconv.ParseFloat(parts[1], 64)
		agList := parts[2]

		if agList == "{}" {
			d[node] += rank
		} else {
			ag[node] = agList
		}
	}

	for node := range ag {
		if _, ok := d[node]; !ok {
			d[node] = 0.0
		}
	}

	for node := range d {
		if _, ok := ag[node]; !ok {
			ag[node] = "{}"
		}
	}

	keys := make([]string, 0, len(d))
	for key := range d {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, node := range keys {
		fmt.Printf("%s\t%.3f\t%s\n", node, d[node], ag[node])
	}
}
