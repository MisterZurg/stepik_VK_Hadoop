package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput() [][]string {
	sc := bufio.NewScanner(os.Stdin)
	var pairs [][]string
	for sc.Scan() {
		fields := strings.Split(sc.Text(), "\t")
		pairs = append(pairs, fields)
	}
	return pairs
}

type Set map[string]bool

func Difference(A, B Set) Set {
	difference := make(Set)
	for k := range A {
		if _, ok := B[k]; !ok {
			difference[k] = true
		}
	}
	return difference
}

func main() {
	pairs := parseInput()
	A := make(Set)
	B := make(Set)

	for _, kv := range pairs {
		switch kv[1] {
		case "A":
			A[kv[0]] = true
		case "B":
			B[kv[0]] = true
		}
	}

	differenceAB := Difference(A, B)
	differenceANS := []int{}
	for v := range differenceAB {
		num, _ := strconv.Atoi(v)
		differenceANS = append(differenceANS, num)
	}

	sort.Ints(differenceANS)
	for _, v := range differenceANS {
		fmt.Println(v)
	}
}
