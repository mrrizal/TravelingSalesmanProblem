package bruteforce

import (
	"fmt"

	"github.com/mrrizal/tsp/utils"
)

// ShortestPath min cost and its path
type shortestPath struct {
	path []int
	cost int
}

// CountCost counting bruteforce way
func countCost(data, permutation [][]int) shortestPath {
	result := shortestPath{
		path: []int{0},
		cost: 0,
	}

	for i := 0; i < len(permutation); i++ {
		cost := data[0][permutation[i][0]]
		currentState := 0
		for j := len(permutation[i]) - 1; j >= 0; j-- {
			cost += data[permutation[i][j]][currentState]
			currentState = permutation[i][j]
		}

		if result.cost == 0 || result.cost >= cost {
			result.path = permutation[i]
			result.cost = cost
		}
	}

	result.path = append([]int{0}, result.path...)
	result.path = append(result.path, 0)
	return result
}

//SolveTSP solve tsp brute foce way
func SolveTSP(data [][]int) {
	index := func() []int {
		result := []int{}
		for i := 0; i < len(data[0])-1; i++ {
			result = append(result, i+1)
		}
		return result
	}()

	permutation := utils.Permutations(index)
	resultCost := countCost(data, permutation)
	fmt.Println(resultCost.path, resultCost.cost)
}
