package sort

import (
	"math/rand"
	"sorting/misc"
	"time"
)

func QuickSort(arr *[]int) {
	rand.Seed(time.Now().UnixNano())
	quickRecur(arr, 0, len(*arr)-1)
}

func quickRecur(arr *[]int, l, r int) {
	if l < r {
		m := partition(arr, l, r)
		quickRecur(arr, l, m-1)
		quickRecur(arr, m+1, r)
	}

}
func partition(arr *[]int, l, r int) int {
	misc.Swap(*arr, r, rand.Int()%(r-l+1)+l)
	i := l - 1
	for j := l; j < r; j++ {
		if (*arr)[j] < (*arr)[r] {
			i++
			misc.Swap(*arr, i, j)
		}
	}
	misc.Swap(*arr, i+1, r)
	return i + 1
}
