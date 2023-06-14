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

func Intersection(A, B Set) Set {
	insertion := make(Set)
	for k := range A {
		if _, ok := B[k]; ok {
			insertion[k] = true
		}
	}
	return insertion
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

	intersectionAB := Intersection(A, B)
	intersectionANS := []int{}
	for v := range intersectionAB {
		num, _ := strconv.Atoi(v)
		intersectionANS = append(intersectionANS, num)
	}

	sort.Ints(intersectionANS)
	for _, v := range intersectionANS {
		fmt.Println(v)
	}
}
