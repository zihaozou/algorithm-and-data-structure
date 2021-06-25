package sort

import "sorting/misc"

func InsertSort(arr []int) {
	for x := 1; x < len(arr); x++ {
		for z, y := x, x-1; y >= 0; z, y = z-1, y-1 {
			if arr[z] < arr[y] {
				misc.Swap(arr, y, z)
			} else {
				break
			}
		}
	}
}
