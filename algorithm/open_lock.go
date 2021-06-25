package algorithm

import "strconv"

func OpenLock(deadends []string, target string) int {
	queue := make([]int, 0)
	dontGo := make(map[int]int8)
	goal, _ := strconv.Atoi(target)
	queue = append(queue, 0)
	var level int
	dontGo[0] = 1
	for _, s := range deadends {
		if s == "0000" {
			return -1
		}
		index, _ := strconv.Atoi(s)
		dontGo[index] = 1
	}
	if target == "0000" {
		return 0
	}
	for len(queue) != 0 {
		level++
		size := len(queue)
		for i := 0; i < size; i++ {

		}
	}
}
