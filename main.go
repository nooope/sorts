package main

import (
	"github.com/nooope/algos/graph/minimumspanningtree"
)

func main() {
	//a := []int{1, 2, 7, 3, 14, 20, 10, 4, 5}
	graph := [][]int{{0, 9, 75, 0, 0},
		{9, 0, 95, 19, 42},
		{75, 95, 0, 51, 66},
		{0, 19, 51, 0, 31},
		{0, 42, 66, 31, 0}}
	minimumspanningtree.Prim(graph)
	//fmt.Println(a)
}
