package sort

import (
	"sorting/misc"
	"sync"
)

/*Merge sort signal thread*/
func MergeSort(arr []int) {
	mergeSortRecur(arr, 0, len(arr)-1)
}

/*Merge sort multi threads*/
func MergeSortMulti(arr []int, numThd int) {
	sem := misc.NewSem(numThd)
	mergeSortMultiRecur(arr, 0, len(arr)-1, sem, nil)
}

func mergeSortRecur(arr []int, l int, r int) {
	if l < r {
		m := (r-l)/2 + l
		mergeSortRecur(arr, l, m)
		mergeSortRecur(arr, m+1, r)
		merge(arr, l, m, r)
	} else {
		return
	}
}

func mergeSortMultiRecur(arr []int, l int, r int, sem *misc.Sem, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	if l < r {
		m := (r-l)/2 + l
		if sem.TryGetSem() {
			var wg sync.WaitGroup
			wg.Add(2)
			go mergeSortMultiRecur(arr, l, m, sem, &wg)
			go mergeSortMultiRecur(arr, m+1, r, sem, &wg)
			wg.Wait()
			sem.GiveSem()
		} else {
			mergeSortMultiRecur(arr, l, m, sem, nil)
			mergeSortMultiRecur(arr, m+1, r, sem, nil)
		}
		merge(arr, l, m, r)
	} else {
		return
	}

}

func merge(arr []int, l int, m int, r int) {
	lArr := make([]int, m+1-l)
	rArr := make([]int, r-m)
	copy(lArr, arr[l:m+1])
	copy(rArr, arr[m+1:r+1])
	for i, j, k := 0, 0, l; k <= r; k++ {
		switch {
		case i >= len(lArr):
			arr[k] = rArr[j]
			j++
		case j >= len(rArr):
			arr[k] = lArr[i]
			i++
		case lArr[i] < rArr[j]:
			arr[k] = lArr[i]
			i++
		default:
			arr[k] = rArr[j]
			j++
		}
	}
}
