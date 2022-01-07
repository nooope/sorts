package main

import (
	"fmt"

	"github.com/nooope/sorts/quick"
)

func main() {
	a := []int{1, 2, 7, 3, 14, 20, 10, 4, 5}
	quick.Sort(a)
	fmt.Println(a)
}
