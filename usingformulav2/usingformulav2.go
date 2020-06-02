package usingformulav2

import (
	"fmt"

	"github.com/mrrizal/tsp/utils"
)

var data [][]int

type result struct {
	cost int
	path []int
}

func getCost(path []int) int {
	cost := 0
	for i := 0; i < len(path); i++ {
		if i+1 < len(path) {
			cost = cost + data[path[i]][path[i+1]]
		} else {
			cost = cost + data[path[i]][path[0]]
		}
	}
	return cost
}

func countCost() {
	result := struct {
		path []int
		cost int
	}{
		path: []int{},
		cost: 0,
	}

	index := func() []int {
		result := []int{}
		for i := 0; i < len(data[0])-1; i++ {
			result = append(result, i+1)
		}
		return result
	}()

	for _, permutation := range utils.Permutations(index) {
		permutation = append([]int{0}, permutation...)
		cost := getCost(permutation)
		if result.cost == 0 {
			result.cost = cost
			result.path = permutation
		} else if result.cost >= cost {
			result.cost = cost
			result.path = permutation
		}
	}

	fmt.Println(result.path, result.cost)
}

//SolveTSP solve tsp using formula
func SolveTSP(datas [][]int) {
	data = datas
	countCost()
}
