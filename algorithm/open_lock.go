package algorithm

import (
	"container/heap"
)

/*
func OpenLock(deadends []string, target string) int {//广度优先算法，不想继续写下去了
	//主要思路就是用广度优先算每次可以变化的数字，再用一个哈希表算已经算过的数字和不能走的数字
	queue := make([]int16, 0)
	dontGo := make(map[int16]bool)
	queue = append(queue, 0)
	var level int
	for _, s := range deadends {
		i, _ := strconv.Atoi(s)
		dontGo[int16(i)] = true
	}
	if dontGo[0] {
		return -1
	} else if target == "0000" {
		return 0
	}
	dontGo[0] = true
	level, _ = strconv.Atoi(target)
	dest := int16(level)
	level = 0
	moveList := func(now int16) []int16 {
		temp := now
		base := 1
		for base < 10000 {
			num := temp % 10
			if num == 0 {

			}
		}
	}
}
*/
type item struct {
	val      string
	priority int
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq *priorityQueue) insert(val string, pri int) {
	newItem := &item{val: val, priority: pri}
	heap.Push(pq, newItem)
}

func distance(now, dest string) int {
	ret := singleDistance(now[0], dest[0])
	ret += singleDistance(now[1], dest[1])
	ret += singleDistance(now[2], dest[2])
	ret += singleDistance(now[3], dest[3])
	return ret
}
func singleDistance(now, dest byte) int {
	var diff int
	if now > dest {
		diff = int(now - dest)
	} else {
		diff = int(dest - now)
	}
	if diff > 5 {
		diff = 10 - diff
	}
	return diff
}

func OpenLock(deadends []string, target string) int {
	costSoFar := make(map[string]int)
	frontier := make(priorityQueue, 0)
	barrier := make(map[string]bool)
	neighbor := func(now string) []string {
		var nblst []string
		var temp string
		for i, b := range now {
			if b == '0' {
				temp = now[:i] + "9" + now[i+1:]
				if !barrier[temp] {
					nblst = append(nblst, temp)
				}
				temp = now[:i] + string(b+1) + now[i+1:]
				if !barrier[temp] {
					nblst = append(nblst, temp)
				}
			} else if b == '9' {
				temp = now[:i] + "0" + now[i+1:]
				if !barrier[temp] {
					nblst = append(nblst, temp)
				}
				temp = now[:i] + string(b-1) + now[i+1:]
				if !barrier[temp] {
					nblst = append(nblst, temp)
				}
			} else {
				temp = now[:i] + string(b+1) + now[i+1:]
				if !barrier[temp] {
					nblst = append(nblst, temp)
				}
				temp = now[:i] + string(b-1) + now[i+1:]
				if !barrier[temp] {
					nblst = append(nblst, temp)
				}
			}
		}
		return nblst
	}
	for _, s := range deadends {
		barrier[s] = true
	}
	if barrier["0000"] {
		return -1
	} else if target == "0000" {
		return 0
	}
	frontier.insert("0000", 0)
	costSoFar["0000"] = 0
	for frontier.Len() != 0 {
		curr := heap.Pop(&frontier).(*item)
		if curr.val == target {
			return costSoFar[curr.val]
		}
		for _, s := range neighbor(curr.val) {
			newCost := costSoFar[curr.val] + 1
			v, ok := costSoFar[s]
			if !ok || newCost < v {
				costSoFar[s] = newCost
				frontier.insert(s, newCost+distance(s, target))
			}
		}
	}
	return -1
}
