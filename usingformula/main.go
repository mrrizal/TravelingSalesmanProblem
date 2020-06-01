package main

import (
	"fmt"

	"github.com/mrrizal/tsp/utils"
)

var data [][]int

type shortestPath struct {
	path []int
	cost int
}

func countC(x, y int) int {
	return data[x][y]

}

func generateFormula(i, k int, s []int) int {
	c := countC(i, k)

	if len(s) == 1 {
		return c + countC(s[0], 0)
	}

	tempI := k
	tempS := utils.FilterSlice(s, []int{tempI})
	tempK := tempS[0]

	result := c + generateFormula(tempI, tempK, tempS)
	return result
}

func countCost() shortestPath {
	result := shortestPath{
		path: []int{},
		cost: 0,
	}

	i := 0
	for _, s := range utils.Permutations([]int{1, 2, 3}) {
		k := s[0]
		cost := generateFormula(i, k, s)
		if result.cost == 0 {
			result.path = s
			result.cost = cost
		} else if result.cost >= cost {
			result.path = s
			result.cost = cost
		}
	}

	result.path = append([]int{0}, result.path...)
	return result
}

func main() {
	// try to solve traveling salesman problem using recursive,
	// but i think it's buggy
	data = [][]int{
		[]int{0, 10, 15, 20},
		[]int{5, 0, 9, 10},
		[]int{6, 13, 0, 12},
		[]int{8, 8, 9, 0},
	}

	result := countCost()
	fmt.Println(result.path, result.cost)

}
