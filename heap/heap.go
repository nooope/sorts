package heap

import "fmt"

func buildMaxHeap(a []int) []int {
	for i := len(a) / 2; i >= 0; i-- {
		maxHeapify(a, i)
	}
	return a
}

func maxHeapify(a []int, i int) { //make subtree with root node i a maxheap
	l := left(i)
	r := right(i)
	largest := i
	if l < len(a) && a[l] > a[i] {
		largest = l //the left leaf is bigger so should be the root
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		swapValues(a, i, largest)
		maxHeapify(a, largest) //our previous i bubbles down
	}

}

func swapValues(a []int, i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func heapSort(a []int) []int {
	buildMaxHeap(a)
	b := a
	for i := len(a) - 1; i >= 0; i-- {
		swapValues(b, 0, i)
		b = b[0 : len(b)-1]
		maxHeapify(b, 0) //TODO: keep track of length
	}
	return a
}

func main() {
	a := []int{1, 2, 7, 3, 14, 20, 10, 4, 5}
	heapSort(a)
	fmt.Println(a)
}
