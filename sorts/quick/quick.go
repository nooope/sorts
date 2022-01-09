package quick

import "github.com/nooope/algos/util"

func Sort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, start int, end int) {
	if start >= end {
		return
	}
	partitionIndex := partition(a, start, end)
	quickSort(a, start, partitionIndex-1)
	quickSort(a, partitionIndex, end)
}

func partition(a []int, start int, end int) int {
	partitionIndex := end //or whatever
	bigger := start - 1
	for i := start; i < end; i++ {
		if a[i] <= a[partitionIndex] { //we're looking for a smaller element and we've found it
			bigger = bigger + 1
			util.SwapValues(a, bigger, i)
		}
	}
	bigger = bigger + 1
	util.SwapValues(a, bigger, partitionIndex)
	return bigger
}
