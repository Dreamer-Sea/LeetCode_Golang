package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}
	res := findItinerary(tickets)
	fmt.Println(res)
}

type pair struct {
	dst     string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	return p[i].dst < p[j].dst
}

func findItinerary(tickets [][]string) []string {
	res := make([]string, 0, len(tickets)+1)
	ticketsMap := make(map[string]pairs)
	for _, t := range tickets {
		toList := ticketsMap[t[0]]
		if toList == nil {
			toList = make(pairs, 0)
		}
		toList = append(toList, &pair{dst: t[1], visited: false})
		ticketsMap[t[0]] = toList
	}
	for _, toList := range ticketsMap {
		sort.Sort(toList)
	}
	pairList := ticketsMap["JFK"]
	var backtracking func([]string, pairs) bool
	backtracking = func(temp []string, pairList pairs) bool {
		if len(tickets)+1 == len(temp) {
			for _, t := range temp {
				res = append(res, t)
			}
			return true
		}
		for _, to := range pairList {
			if to.visited {
				continue
			}
			to.visited = true
			temp = append(temp, to.dst)
			if backtracking(temp, ticketsMap[to.dst]) {
				return true
			}
			temp = temp[:len(temp)-1]
			to.visited = false
		}
		return false
	}
	backtracking([]string{"JFK"}, pairList)
	return res
}
