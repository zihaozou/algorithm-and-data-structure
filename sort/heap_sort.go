package sort

import "sorting/misc"

func HeapSort(arr []int) {
	heap := misc.Heap{Arr: &arr, HeapSize: 0}
	misc.BuildMaxHeap(&heap)
	for heap.HeapSize > 1 {
		misc.Swap((*heap.Arr), 0, heap.HeapSize-1)
		heap.HeapSize--
		misc.MaxHeapify(&heap, 0)
	}
}
