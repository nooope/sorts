package bubble

import "github.com/nooope/algos/util"

func Sort(a []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				util.SwapValues(a, i, i+1)
				sorted = false
			}
		}
	}
}
