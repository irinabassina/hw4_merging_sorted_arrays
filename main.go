package main

import (
	"fmt"
	"sort"
)

func mergeSort(a, b []int) (res []int) {
	lenA, lenB := len(a), len(b)
	res = make([]int, 0, lenA+lenB)

	idxA, idxB := 0, 0
	for idxA+idxB < lenA+lenB {
		if idxB == lenB || (idxA < lenA && a[idxA] < b[idxB]) {
			res = append(res, a[idxA])
			idxA++
		} else {
			res = append(res, b[idxB])
			idxB++
		}

	}

	return res
}

func sortArr(c []int, d []int) ([]int, []int) {
	sort.Ints(c[:])
	sort.Ints(d[:])

	return c, d
}

func main() {

	array1 := []int{4, 25, 132, 43}
	array2 := []int{1, 27, 13, 142, 90}

	mergedArray := mergeSort(sortArr(array1, array2))
	fmt.Println("Merged array is:", mergedArray)
}
