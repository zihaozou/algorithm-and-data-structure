package datastructure

import (
	"fmt"
)

type PQItf interface {
	Len() int
	Swap(i int, j int)
	Larger(i int, j int) bool
	Less(i int, j int) bool
	Push(itm interface{})
	Seek() interface{}
	DropTail()
}

type PQHeap struct {
	ItmLst PQItf
	mode   uint8 //0: maxheap, otherwise: minheap
}

func (heap *PQHeap) BuildPQ(mode uint8) {
	heap.mode = mode
	for i := heap.ItmLst.Len() / 2; i >= 1; i-- {
		heap.heapify(i - 1)
	}
}

func (heap *PQHeap) heapify(i int) int {
	len := heap.ItmLst.Len()
	temp := i + 1
	l, r := temp<<1, (temp<<1)+1
	var move int
	if heap.mode == 0 {
		if l <= len && heap.ItmLst.Larger(l-1, i) {
			move = l - 1
		} else {
			move = i
		}
		if r <= len && heap.ItmLst.Larger(r-1, move) {
			move = r - 1
		}
	} else {
		if l <= len && heap.ItmLst.Less(l-1, i) {
			move = l - 1
		} else {
			move = i
		}
		if r <= len && heap.ItmLst.Less(r-1, move) {
			move = r - 1
		}
	}
	if move != i {
		heap.ItmLst.Swap(i, move)
		return heap.heapify(move)
	}
	return i
}

func (heap *PQHeap) Pop() (itm interface{}, err error) {
	if heap.ItmLst.Len() < 1 {
		err = fmt.Errorf("heap size underflow")
		return nil, err
	}
	itm = heap.ItmLst.Seek()
	heap.ItmLst.Swap(0, heap.ItmLst.Len()-1)
	heap.ItmLst.DropTail()
	heap.heapify(0)
	return itm, nil
}

func (heap *PQHeap) Seek() (itm interface{}) {
	return heap.ItmLst.Seek()
}

func (heap *PQHeap) Push(itm interface{}) {
	heap.ItmLst.Push(itm)
	heap.up(heap.ItmLst.Len())

}

func (heap *PQHeap) up(i int) int {
	father := (i + 1) >> 1
	if heap.mode == 0 {
		if father > 0 || heap.ItmLst.Larger(i, father-1) {
			heap.ItmLst.Swap(father-1, i)
			return heap.up(father - 1)
		} else {
			return i
		}
	} else {
		if father > 0 || heap.ItmLst.Less(i, father-1) {
			heap.ItmLst.Swap(father-1, i)
			return heap.up(father - 1)
		} else {
			return i
		}
	}

}
func (heap *PQHeap) Update(i int, pri int) {
	heap.heapify(heap.up(i))
}
