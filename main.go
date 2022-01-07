package main

import (
	"fmt"

	"github.com/nooope/sorts/merge"
)

func main() {
	a := []int{1, 2, 7, 3, 14, 20, 10, 4, 5}
	merge.Sort(a)
	fmt.Println(a)
}
