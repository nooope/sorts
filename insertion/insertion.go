package insertion

func Sort(a []int) {
	for sortedI := 1; sortedI < len(a); sortedI++ {
		position := 0
		for p := sortedI - 1; p >= 0; p-- {
			if a[sortedI] >= a[p] { //stop here
				position = p + 1
				break
			}
		}
		temp := a[sortedI]
		for i := sortedI; i > position; i-- {
			a[i] = a[i-1]
		}
		a[position] = temp
	}

}
