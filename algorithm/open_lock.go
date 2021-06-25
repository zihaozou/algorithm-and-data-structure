package algorithm

import "strconv"

func OpenLock(deadends []string, target string) int {
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
