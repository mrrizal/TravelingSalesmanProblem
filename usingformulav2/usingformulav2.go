package usingformulav2

import "fmt"

var data [][]int

type result struct {
	cost int
	path []int
}

var res result

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func getCost(path []int) int {
	cost := 0
	for i := 0; i < len(path); i++ {
		if i+1 < len(path) {
			cost = cost + data[path[i]][path[i+1]]
		} else {
			cost = cost + data[path[i]][path[0]]
		}

		if res.cost != 0 {
			if cost > res.cost {
				return 0
			}
		}
	}
	return cost
}

func countCost() {
	res.cost = 0
	res.path = []int{}

	index := func() []int {
		result := []int{}
		for i := 0; i < len(data[0])-1; i++ {
			result = append(result, i+1)
		}
		return result
	}()

	for tempPermutation := make([]int, len(index)); tempPermutation[0] < len(tempPermutation); nextPerm(tempPermutation) {
		permutation := getPerm(index, tempPermutation)
		permutation = append([]int{0}, permutation...)
		cost := getCost(permutation)

		if cost == 0 {
			continue
		} else if res.cost == 0 || res.cost > cost {
			res.cost = cost
			res.path = permutation
		}
	}

	fmt.Println(res.path, res.cost)
}

//SolveTSP solve tsp using formula
func SolveTSP(datas [][]int) {
	data = datas
	countCost()
}
