package merge

func Sort(a []int) {
	mergeSort(a, 0, len(a)-1)
}

func mergeSort(a []int, i int, j int) {
	if i >= j {
		return
	}
	mid := (i + j) / 2
	mergeSort(a, i, mid)
	mergeSort(a, mid+1, j)
	merge(a, i, j, mid)
}

func merge(a []int, start int, end int, mid int) {
	leftCopy := make([]int, mid-start+1)
	for i := 0; i < len(leftCopy); i++ {
		leftCopy[i] = a[start+i]
	}

	rightCopy := make([]int, end-mid)
	for i := 0; i < len(rightCopy); i++ {
		rightCopy[i] = a[mid+1+i]
	}

	for i, j := 0, 0; i < len(leftCopy) || j < len(rightCopy); {
		if j >= len(rightCopy) || (i < len(leftCopy) && leftCopy[i] < rightCopy[j]) {
			a[start+i+j] = leftCopy[i]
			i++
		} else {
			a[start+i+j] = rightCopy[j]
			j++
		}
	}
}
