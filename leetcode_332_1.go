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
	res := findItinerary1(tickets)
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

func findItinerary1(tickets [][]string) []string {
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
	beginPairList := ticketsMap["JFK"]
	temp := []string{"JFK"}
	backtracking9(tickets, temp, ticketsMap, beginPairList, &res)
	return res
}

func backtracking9(tickets [][]string, temp []string, ticketsMap map[string]pairs, pairList pairs, res *[]string) bool {
	if len(tickets)+1 == len(temp) {
		for _, t := range temp {
			*res = append(*res, t)
		}
		return true
	}
	for _, to := range pairList {
		if to.visited {
			continue
		}
		to.visited = true
		temp = append(temp, to.dst)
		if backtracking9(tickets, temp, ticketsMap, ticketsMap[to.dst], res) {
			return true
		}
		to.visited = false
		temp = temp[:len(temp)-1]
	}
	return false
}
