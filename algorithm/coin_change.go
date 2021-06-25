package algorithm

import "sorting/misc"

func CoinChange(coins []int, amount int) int {
	state := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		state[i] = -1
	}
	for _, v := range coins {
		if v < amount {
			state[v] = 1
			for i := v + 1; i <= amount; i++ {
				if state[i-v] != -1 {
					if state[i] != -1 {
						state[i] = misc.IntMin(state[i], state[i-v]+1)
					} else {
						state[i] = state[i-v] + 1
					}
				}
			}
		}
	}
	return state[amount]
}
