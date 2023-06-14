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

func Union(A, B Set) Set {
	union := make(Set)
	for k := range A {
		union[k] = true
	}

	for k := range B {
		union[k] = true
	}
	return union
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

	unionAB := Union(A, B)
	unionANS := []int{}
	for v := range unionAB {
		num, _ := strconv.Atoi(v)
		unionANS = append(unionANS, num)
	}

	sort.Ints(unionANS)
	for _, v := range unionANS {
		fmt.Println(v)
	}
}
