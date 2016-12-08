package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

func (p Pair) String() string {
	return fmt.Sprintf("%s: %d", p.Key, p.Value)
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}

func main() {
	input, _ := ioutil.ReadFile("./day06.txt")
	raw_lines := strings.TrimSpace(string(input))
	lines := strings.Split(raw_lines, "\n")

	columns := [8]map[rune][]int{}
	for i := 0; i < len(columns); i++ {
		columns[i] = make(map[rune][]int)
	}

	for _, line := range lines {
		for i, r := range line {
			columns[i][r] = append(columns[i][r], 1)
		}
	}

	for i := 0; i < len(columns); i++ {
		list := PairList{}

		for k, v := range columns[i] {
			list = append(list, Pair{string(k), len(v)})
		}

		sort.Sort(list)
		fmt.Printf("%v", list[0].Key)
	}
}
