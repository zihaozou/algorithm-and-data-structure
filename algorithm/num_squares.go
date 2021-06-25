package algorithm

import (
	"math"
	"sorting/misc"
)

func NumSquares(n int) int { //广度优先
	if n == 0 {
		return 0
	}
	cache := make([]int, int(math.Sqrt(float64(n))+1))
	square := func(num int) int {
		if cache[num] != 0 {
			return cache[num]
		} else {
			cache[num] = num * num
			return cache[num]
		}
	}
	queue := make([]int, 1)
	queue[0] = 0
	var level int
	for len(queue) != 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			for j := 1; square(j) <= n; j++ {
				new := square(j) + queue[0]
				if new == n {
					return level
				} else if new < n {
					queue = append(queue, new)
				}
			}
			queue = queue[1:]
		}
	}
	return n
}

func NumSquares2(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minimum := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minimum = misc.IntMin(minimum, dp[i-j*j])
		}
		dp[i] = minimum
	}
	return dp[n]
}
