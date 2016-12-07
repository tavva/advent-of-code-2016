package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[1].Key < p[j].Key
	}
	return p[i].Value < p[j].Value
}

func main() {
	input, _ := ioutil.ReadFile("./day04.txt")
	room_strings := strings.TrimSpace(string(input))
	lines := strings.Split(room_strings, "\n")

	re := regexp.MustCompile("^([a-z-]+)([0-9]+)\\[([a-z]+)\\]$")

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		name := strings.Replace(matches[1], "-", "", -1)
		sector_id := matches[2]
		checksum := matches[3]

		letter_counts := make(map[rune][]int)
		for _, x := range name {
			letter_counts[x] = append(letter_counts[x], 1)
		}

		list := make(PairList, len(letter_counts))

		for k, v := range letter_counts {
			list = append(list, Pair{string(k), len(v)})
		}
		sort.Sort(sort.Reverse(list))

		for _, p := range list {
			fmt.Printf("%v: %v\n", p.Key, p.Value)
		}

		println(name)
		println(sector_id)
		println(checksum)
	}
}
