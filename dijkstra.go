package main

import (
	"kubernetes/pkg/util/slice"
	"math"
)

func dijkstra() (string, int) {
	graph := map[string]map[string]int{
		"start": {
			"A": 5,
			"B": 2,
		},
		"A": {
			"C": 4,
			"D": 2,
		},
		"B": {
			"A": 8,
			"D": 7,
		},
		"C": {
			"D":   6,
			"fin": 3,
		},
		"D": {
			"fin": 1,
		},
		"fin": {},
	}
	costs := map[string]int{
		"A":   5,
		"B":   2,
		"C":   math.MaxInt,
		"D":   math.MaxInt,
		"fin": math.MaxInt,
	}
	parents := map[string]string{
		"A":   "start",
		"B":   "start",
		"C":   "",
		"D":   "",
		"fin": "",
	}
	processed := make([]string, 0)

	for node := unprocessedLowCost(costs, processed); node != ""; node = unprocessedLowCost(costs, processed) {
		cost := costs[node]
		neighbours := graph[node]
		for nn, nc := range neighbours {
			new_cost := cost + nc
			if new_cost < costs[nn] {
				costs[nn] = new_cost
				parents[nn] = node
			}
		}
		processed = append(processed, node)
	}
	route := "fin"
	for prev, ok := parents["fin"]; ok; prev, ok = parents[prev] {
		route = prev + " -> " + route
	}

	return route, costs["fin"]
}

func unprocessedLowCost(costs map[string]int, processed []string) string {
	lowestCost := 0
	lowestNode := ""
	for k, c := range costs {
		if slice.ContainsString(processed, k, nil) {
			continue
		}
		if lowestNode == "" {
			lowestNode = k
			lowestCost = c
			continue
		}
		if c < lowestCost {
			lowestNode = k
			lowestCost = c
		}
	}
	return lowestNode
}
