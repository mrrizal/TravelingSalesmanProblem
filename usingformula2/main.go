package main

import (
	"fmt"

	"github.com/mrrizal/tsp/utils"
)

var data [][]int

type result struct {
	cost int
	path []int
}

var shortestPath map[int]map[string]int

func generatePermutation(index []int, key int) {
	firstIndex := index[0]
	index = utils.FilterSlice(index, []int{firstIndex})

	permutations := utils.Permutations(index)
	for _, permutation := range permutations {
		cost := data[firstIndex][permutation[0]]

		if len(permutation) == 1 {
			shortestPath[key]["tempCost"] += cost

			if shortestPath[key]["cost"] == 0 {
				shortestPath[key]["cost"] = shortestPath[key]["tempCost"]
			} else if shortestPath[key]["cost"] >= shortestPath[key]["tempCost"] {
				shortestPath[key]["cost"] = shortestPath[key]["tempCost"]
			}

			shortestPath[key]["tempCost"] = 0
		} else {
			shortestPath[key]["tempCost"] += cost
		}

		if len(permutation) > 1 {
			generatePermutation(permutation, key)
		}
	}
}

func countCost() {
	result := struct {
		path []int
		cost int
	}{
		path: []int{},
		cost: 0,
	}

	shortestPath = make(map[int]map[string]int)

	for key, permutation := range utils.Permutations([]int{1, 2, 3}) {
		shortestPath[key] = map[string]int{
			"cost":     0,
			"tempCost": 0,
		}
		generatePermutation(permutation, key)
		shortestPath[key]["cost"] += data[0][permutation[0]]
		shortestPath[key]["cost"] += data[permutation[len(permutation)-1]][0]

		if result.cost == 0 {
			result.path = permutation
			result.cost = shortestPath[key]["cost"]
		} else if result.cost >= shortestPath[key]["cost"] {
			result.path = permutation
			result.cost = shortestPath[key]["cost"]
		}
	}
	result.path = append([]int{0}, result.path...)
	fmt.Println(result.path, result.cost)
}
func main() {
	data = [][]int{
		[]int{0, 10, 15, 20},
		[]int{5, 0, 9, 10},
		[]int{6, 13, 0, 12},
		[]int{8, 8, 9, 0},
	}
	countCost()
}
