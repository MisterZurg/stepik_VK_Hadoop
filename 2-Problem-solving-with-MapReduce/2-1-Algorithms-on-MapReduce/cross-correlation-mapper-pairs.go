package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	i string
	j string
}

func (pr Pair) String() string {
	return fmt.Sprintf("%s,%s", pr.i, pr.j)
}

type MapperPairs struct{}

func NewPairsMapper() *MapperPairs {
	return new(MapperPairs)
}

func (mp *MapperPairs) Map(items []string) {
	for _, i := range items {
		for _, j := range items {
			if i != j {
				mp.Emit(Pair{i, j}, 1)
			}
		}
	}
}

func (mp *MapperPairs) Emit(pr Pair, cnt int) {
	fmt.Printf("%s\t%d\n", pr, cnt)
}

func ParseCrossCorrelation() [][]string {
	var objects [][]string
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		items := strings.Split(line, " ")
		objects = append(objects, items)
	}

	return objects
}

func main() {
	objects := ParseCrossCorrelation()
	mp := NewPairsMapper()

	for _, batch := range objects {
		mp.Map(batch)
	}
}
