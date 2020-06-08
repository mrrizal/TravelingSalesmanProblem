package main

import (
	"fmt"

	"github.com/mrrizal/tsp/utils"
)

type costResult struct {
	matrix     [][]int
	cost       int
	startIndex int
}

var currentIndex []int

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

func findCost(matrix [][]int, startPoint, upper int, index []int) costResult {
	result := costResult{}

	for _, value := range index {
		tempMatrix := utils.Transpose(matrix)
		for i := 0; i < len(tempMatrix[value]); i++ {
			tempMatrix[value][i] = -1
		}

		tempMatrix = utils.Transpose(tempMatrix)
		tempMatrix[value][startPoint] = -1
		for i := 0; i < len(tempMatrix); i++ {
			tempMatrix[startPoint][i] = -1
		}

		tempReduceMatrix, tempReduceCost := reducingMatrix(tempMatrix)
		tempCost := matrix[startPoint][value] + upper + tempReduceCost

		tempResult := costResult{
			matrix:     tempReduceMatrix,
			cost:       tempCost,
			startIndex: value,
		}

		if len(result.matrix) == 0 {
			result = tempResult
		} else if result.cost > tempResult.cost {
			result = tempResult
		}

	}
	return result
}

func main() {
	data := [][]int{
		[]int{-1, 20, 30, 10, 11},
		[]int{15, -1, 16, 4, 2},
		[]int{3, 5, -1, 2, 4},
		[]int{19, 6, 18, -1, 3},
		[]int{16, 4, 7, 16, -1},
	}

	reduceMatrix, reduceCost := reducingMatrix(data)

	startPoint := 0

	// initial index
	index := func() []int {
		result := []int{}
		for i := 0; i < len(data); i++ {
			result = append(result, i)
		}
		return result
	}()

	currentIndex = append(currentIndex, startPoint)
	index = utils.FilterSlice(index, currentIndex)

	costMatrix := findCost(reduceMatrix, startPoint, reduceCost, index)
	currentIndex = append(currentIndex, costMatrix.startIndex)
	index = utils.FilterSlice(index, currentIndex)
	for i := 0; i < len(costMatrix.matrix); i++ {
		fmt.Println(costMatrix.matrix[i])
	}
	fmt.Println()
	fmt.Println(currentIndex)
	fmt.Println(index)

}
