package util
func SwapValues( a []int,i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}