package main

import (
	"fmt"
	"kubernetes/pkg/util/slice"
	"math"
)

func dijkstra() string {
	graph := map[string]map[string]int{
		"start": {
			"A": 6,
			"B": 2,
		},
		"A": {
			"fin": 1,
		},
		"B": {
			"A":   3,
			"fin": 5,
		},
		"fin": {},
	}
	costs := map[string]int{
		"A":   6,
		"B":   2,
		"fin": math.MaxInt,
	}
	parents := map[string]string{
		"A":   "start",
		"B":   "start",
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
	return fmt.Sprintf("%v", parents)
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
