package algorithm

func Change(amount int, coins []int) int {
	state := make([]int, amount+1)
	state[0] = 1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			state[i] += state[i-c]
		}
	}
	return state[amount]
}
