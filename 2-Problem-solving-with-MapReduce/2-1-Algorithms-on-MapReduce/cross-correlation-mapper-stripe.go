package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stripe map[string]int

func (st Stripe) String() string {
	res := ""

	for k, v := range st {
		res += fmt.Sprintf("%s:%d,", k, v)
	}
	return res[:len(res)-1]
}

type MapperStripe struct{}

func NewStripeMapper() *MapperStripe {
	return new(MapperStripe)
}

func (mp *MapperStripe) Map(items []string) {
	for _, i := range items {
		// stripe
		h := make(map[string]int)
		for _, j := range items {
			if i != j {
				h[j]++
			}
		}
		mp.Emit(i, h)
	}
}

func (mp *MapperStripe) Emit(item string, stripe Stripe) {
	fmt.Printf("%s\t%s\n", item, stripe)
}

func ParseCrossCorrelationStripe() [][]string {
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
	objects := ParseCrossCorrelationStripe()
	mp := NewStripeMapper()

	for _, batch := range objects {
		mp.Map(batch)
	}
}
