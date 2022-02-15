
package main

import (
	"fmt"
)

func partition(a *[10]int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a *[10]int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	array := [10]int{55, 90, 74, 20, 16, 46, 43, 59, 2, 99}
	
	quickSort(&array, 0, len(array)-1)
	fmt.Println(array)
}