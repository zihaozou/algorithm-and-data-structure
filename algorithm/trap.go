//接雨水
package algorithm

import "sorting/misc"

func Trap(height []int) int {
	var lHighest, rHighest, l, r, ret int
	l, r = 0, len(height)-1

	for l != r {
		lHighest = misc.IntMax(lHighest, height[l])
		rHighest = misc.IntMax(rHighest, height[r])
		if lHighest < rHighest {
			ret += lHighest - height[l]
			l++
		} else {
			ret += rHighest - height[r]
			r--
		}
	}
	return ret
}
