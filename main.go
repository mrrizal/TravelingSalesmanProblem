package main

import (
	"flag"
	"fmt"

	"github.com/mrrizal/tsp/bruteforce"
	"github.com/mrrizal/tsp/usingformulav1"
	"github.com/mrrizal/tsp/usingformulav2"
)

func main() {
	// data := [][]int{
	// 	[]int{0, 10, 15, 20},
	// 	[]int{5, 0, 9, 10},
	// 	[]int{6, 13, 0, 12},
	// 	[]int{8, 8, 9, 0},
	// }

	data := [][]int{
		[]int{0, 20, 30, 10, 11},
		[]int{15, 0, 16, 4, 2},
		[]int{3, 5, 0, 2, 4},
		[]int{19, 6, 18, 0, 3},
		[]int{16, 4, 7, 16, 0},
	}

	module := flag.String("module", "bruteforce", "choose module")
	flag.Parse()

	switch *module {
	case "bruteforce":
		fmt.Println(*module)
		bruteforce.SolveTSP(data)
	case "formula-v1":
		fmt.Println(*module)
		usingformulav1.SolveTSP(data)
	case "formula-v2":
		fmt.Println(*module)
		usingformulav2.SolveTSP(data)
	}

}
