package selection

import "github.com/nooope/sorts/util"

func Sort(a []int) {
	for sortedI := 0; sortedI < len(a)-1; sortedI++ {
		//find minimum
		min := sortedI
		for i := sortedI + 1; i < len(a); i++ {
			if a[i] < a[min] {
				min = i
			}
		}
		util.SwapValues(a, sortedI, min)
	}
}
