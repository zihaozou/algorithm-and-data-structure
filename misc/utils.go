package misc

import "time"

/*Semaphore*/
type Sem struct {
	semChan  chan int
	capacity int
}

func NewSem(cap int) *Sem {
	return &Sem{semChan: make(chan int, cap), capacity: cap}
}

func (sem *Sem) GetSem() {
	sem.semChan <- 0
}

func (sem *Sem) GiveSem() {
	<-sem.semChan
}

func (sem *Sem) TryGetSem() bool {
	select {
	case sem.semChan <- 0:
		return true
	default:
		return false
	}
}

func (sem *Sem) TryGetSemTimeout(ms int /*in Milliseconds */) bool {
	select {
	case sem.semChan <- 0:
		return true
	case <-time.After(time.Duration(ms) * time.Millisecond):
		return false
	}
}

func (sem *Sem) CheckSemAvaliable() int {
	return sem.capacity - len(sem.semChan)
}

/*Semaphore*/

/*Heap, father: i>>1, left child: i<<1, right child: (i<<1)+1*/
type Heap struct {
	Arr      *[]int
	HeapSize int
}

func MaxHeapify(heap *Heap, i int) {
	temp := i + 1
	l, r := temp<<1, (temp<<1)+1
	var largest int
	if l <= heap.HeapSize && (*heap.Arr)[l-1] > (*heap.Arr)[i] {
		largest = l - 1
	} else {
		largest = i
	}
	if r <= heap.HeapSize && (*heap.Arr)[r-1] > (*heap.Arr)[largest] {
		largest = r - 1
	}
	if largest != i {
		Swap((*heap.Arr), i, largest)
		MaxHeapify(heap, largest)
	}
}

func BuildMaxHeap(heap *Heap) {
	heap.HeapSize = len((*heap.Arr))
	for i := len((*heap.Arr)) / 2; i >= 1; i-- {
		MaxHeapify(heap, i-1)
	}
}

/*Heap*/

/*Swap function*/
func Swap(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

/*Swap function*/

func IntMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
