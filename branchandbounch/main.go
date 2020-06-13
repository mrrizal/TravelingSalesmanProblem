package main

import (
	"fmt"

	"github.com/mrrizal/tsp/utils"
)

type costResult struct {
	matrix     [][]int
	cost       int
	startIndex int
	index      []int
	path       []int
}

var upper int
var savedCostResult []costResult

func findMinimumValue(data [][]int) []int {
	result := []int{}
	for i := 0; i < len(data); i++ {
		result = append(result, utils.MinInts(data[i]))
	}
	return result
}

func reduceRow(data [][]int, minimumValue []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(data); i++ {
		temp := []int{}
		for j := 0; j < len(data[i]); j++ {
			if j != i {
				if data[i][j] == -1 {
					temp = append(temp, -1)
				} else {
					temp = append(temp, data[i][j]-minimumValue[i])
				}
			} else {
				temp = append(temp, -1)
			}
		}
		result = append(result, temp)
	}
	return result
}

func reducingMatrix(data [][]int) ([][]int, int) {
	minimumValueRow := findMinimumValue(data)
	data = reduceRow(data, minimumValueRow)
	minimumValueColumn := findMinimumValue(utils.Transpose(data))
	data = utils.Transpose(reduceRow(utils.Transpose(data), minimumValueColumn))

	costReduction := func() int {
		result := 0
		for _, value := range append(minimumValueColumn, minimumValueRow...) {
			result += value
		}
		return result
	}()
	return data, costReduction
}

func findCost(matrix [][]int, startPoint, upper int, index, path []int) costResult {
	result := costResult{}

	for _, value := range index {
		tempMatrix := utils.Transpose(matrix)
		for i := 0; i < len(tempMatrix[value]); i++ {
			tempMatrix[value][i] = -1
		}

		tempMatrix = utils.Transpose(tempMatrix)
		tempMatrix[value][startPoint] = -1
		if startPoint != 0 {
			tempMatrix[value][0] = -1
		}

		for i := 0; i < len(tempMatrix); i++ {
			tempMatrix[startPoint][i] = -1
		}

		tempReduceMatrix, tempReduceCost := reducingMatrix(tempMatrix)
		tempCost := matrix[startPoint][value] + upper + tempReduceCost
		tempResult := costResult{
			matrix:     tempReduceMatrix,
			cost:       tempCost,
			startIndex: value,
			index:      index,
			path:       path,
		}

		savedCostResult = append(savedCostResult, tempResult)

		if len(result.matrix) == 0 {
			result = tempResult
		} else if result.cost > tempResult.cost {
			result = tempResult
		}

	}
	return result
}

func printMatrix(data [][]int) {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Print(data[i][j])
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

func findPath(costMatrix costResult) costResult {
	savedCostResult = []costResult{}

	for len(costMatrix.index) != 0 {
		costMatrix = findCost(
			costMatrix.matrix,
			costMatrix.startIndex,
			costMatrix.cost,
			costMatrix.index,
			costMatrix.path)

		costMatrix.path = append(costMatrix.path, costMatrix.startIndex)
		costMatrix.index = utils.FilterSlice(costMatrix.index, costMatrix.path)

		if len(costMatrix.index) == 0 {
			upper = costMatrix.cost
		}
	}
	return costMatrix
}

func findMinCostResult(costMatrix costResult, path []int) costResult {
	for i := len(savedCostResult) - 1; i >= 0; i-- {
		if (savedCostResult[i].cost < costMatrix.cost) && (savedCostResult[i].startIndex != costMatrix.path[1]) && !utils.EqualSlice(savedCostResult[i].path, path) {
			return savedCostResult[i]
		}

	}
	return costResult{}
}

//solveTSP function for solve traveling salesman problem
func solveTSP(datas [][]int) costResult {
	data := datas

	reduceMatrix, reduceCost := reducingMatrix(data)
	upper = reduceCost
	startPoint := 0

	// initial index
	index := func() []int {
		result := []int{}
		for i := 0; i < len(data); i++ {
			result = append(result, i)
		}
		return result
	}()

	path := append([]int{}, startPoint)
	index = utils.FilterSlice(index, path)
	costMatrix := costResult{
		matrix:     reduceMatrix,
		startIndex: startPoint,
		cost:       upper,
		index:      index,
		path:       path,
	}

	costMatrix = findPath(costMatrix)

	temp := costResult{}
	for {
		temp = findMinCostResult(costMatrix, temp.path)
		if temp.cost != 0 {
			tempMatrix := findPath(temp)
			if tempMatrix.cost < costMatrix.cost {
				costMatrix = tempMatrix
			}
		} else {
			break
		}
	}
	return costMatrix
}

func main() {
	// data := [][]int{
	// 	[]int{-1, 20, 30, 10, 11},
	// 	[]int{15, -1, 16, 4, 2},
	// 	[]int{3, 5, -1, 2, 4},
	// 	[]int{19, 6, 18, -1, 3},
	// 	[]int{16, 4, 7, 16, -1},
	// }

	// data := [][]int{
	// 	[]int{0, 2451, 713, 1018, 1631, 1374, 2408, 213, 2571, 875, 1420, 2145, 1972},
	// 	[]int{2451, 0, 1745, 1524, 831, 1240, 959, 2596, 403, 1589, 1374, 357, 579},
	// 	[]int{713, 1745, 0, 355, 920, 803, 1737, 851, 1858, 262, 940, 1453, 1260},
	// 	[]int{1018, 1524, 355, 0, 700, 862, 1395, 1123, 1584, 466, 1056, 1280, 987},
	// 	[]int{1631, 831, 920, 700, 0, 663, 1021, 1769, 949, 796, 879, 586, 371},
	// 	[]int{1374, 1240, 803, 862, 663, 0, 1681, 1551, 1765, 547, 225, 887, 999},
	// 	[]int{2408, 959, 1737, 1395, 1021, 1681, 0, 2493, 678, 1724, 1891, 1114, 701},
	// 	[]int{213, 2596, 851, 1123, 1769, 1551, 2493, 0, 2699, 1038, 1605, 2300, 2099},
	// 	[]int{2571, 403, 1858, 1584, 949, 1765, 678, 2699, 0, 1744, 1645, 653, 600},
	// 	[]int{875, 1589, 262, 466, 796, 547, 1724, 1038, 1744, 0, 679, 1272, 1162},
	// 	[]int{1420, 1374, 940, 1056, 879, 225, 1891, 1605, 1645, 679, 0, 1017, 1200},
	// 	[]int{2145, 357, 1453, 1280, 586, 887, 1114, 2300, 653, 1272, 1017, 0, 504},
	// 	[]int{1972, 579, 1260, 987, 371, 999, 701, 2099, 600, 1162, 1200, 504, 0},
	// }

	data := [][]int{
		[]int{0, 29, 20, 21, 16, 31, 100, 12, 4, 31, 18},
		[]int{29, 0, 15, 29, 28, 40, 72, 21, 29, 41, 12},
		[]int{20, 15, 0, 15, 14, 25, 81, 9, 23, 27, 13},
		[]int{21, 29, 15, 0, 4, 12, 92, 12, 25, 13, 25},
		[]int{16, 28, 14, 4, 0, 16, 94, 9, 20, 16, 22},
		[]int{31, 40, 25, 12, 16, 0, 95, 24, 36, 3, 37},
		[]int{100, 72, 81, 92, 94, 95, 0, 90, 101, 99, 84},
		[]int{12, 21, 9, 12, 9, 24, 90, 0, 15, 25, 13},
		[]int{4, 29, 23, 25, 20, 36, 101, 15, 0, 35, 18},
		[]int{31, 41, 27, 13, 16, 3, 99, 25, 35, 0, 38},
		[]int{18, 12, 13, 25, 22, 37, 84, 13, 18, 38, 0},
	}
	result := solveTSP(data)
	fmt.Println(result)

}
