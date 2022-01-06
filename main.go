package main

import (
	"fmt"

	"github.com/nooope/sorts/bubble"
)

func main() {
	a := []int{1, 2, 7, 3, 14, 20, 10, 4, 5}
	bubble.Sort(a)
	fmt.Println(a)
}
