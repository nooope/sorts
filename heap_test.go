package main

import "testing"
import "fmt"

func TestHeap(t *testing.T) {
	a := []int{1,4,3,5,2,6,36,10,23}
	buildMaxHeap(a)
	for i := 0; i < len(a);i++{
		l := left(i)
		r := right(i)
		if l < len(a) && a[l] > a[i]{
			t.Error(fmt.Sprintf("Invalid Heap at index %d, bigger left node: %d > %d", i,a[l], a[i]))
		}
		if r < len(a) && a[r] > a[i]{
			t.Error(fmt.Sprintf("Invalid Heap at index %d, bigger righ node: %d > %d", i,a[r], a[i]))
		}
	}
}