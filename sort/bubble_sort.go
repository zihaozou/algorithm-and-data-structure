package sort

import "sorting/misc"

func BubbleSort(arr []int) {
	for i := range arr {
		for j := len(arr) - 1; j >= i+1; j-- {
			if arr[j] < arr[i] {
				misc.Swap(arr, i, j)
			}
		}
	}
}
