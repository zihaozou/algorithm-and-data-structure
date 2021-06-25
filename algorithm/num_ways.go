package algorithm

import "sorting/misc"

func NumWays(steps int, arrLen int) int {
	mod:=int(1e9+7)
	max := misc.IntMin(steps/2, arrLen-1)
	state := make([][]int, steps+1)
	for i := range state {
		state[i] = make([]int, max+1)
	}
	state[0][0] = 1
	for i := 1; i <= steps; i++ {
		edge:=misc.IntMin(steps-i,max)
		for j := 0; j <= edge; j++ {
			state[i][j] += state[i-1][j]%mod
			if j > 0 {
				state[i][j] += state[i-1][j-1]%mod
			}
			if j < max {
				state[i][j] += state[i-1][j+1]%mod
			}
		}
	}
	return state[steps][0]
}
