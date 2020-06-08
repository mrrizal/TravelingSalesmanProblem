package utils

import (
	"sort"
)

// FilterSlice filter slice x with y
func FilterSlice(x, y []int) []int {
	result := []int{}

	for _, value := range x {
		searchResult := sort.SearchInts(y, value)
		if searchResult == len(y) {
			result = append(result, value)
		} else if y[searchResult] != value {
			result = append(result, value)
		}
	}
	return result

}

// Permutations return permutation from given slice of int
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

//Transpose matrix int
func Transpose(slice [][]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

// MinInts find min value from slice of int
func MinInts(slice []int) int {
	result := -1
	for _, value := range slice {
		if value != -1 {
			if result == -1 {
				result = value
			} else if result > value {
				result = value
			}
		}
	}

	if result == -1 {
		result = 0
	}
	return result
}
