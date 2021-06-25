package algorithm

import "sorting/misc"

func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	mod := int(1e9 + 7)
	numWork := len(group)
	state := make([][][]int, numWork+1)
	for i := range state {
		state[i] = make([][]int, n+1)
		for j := range state[i] {
			state[i][j] = make([]int, minProfit+1)
			state[0][j][0] = 1
		}
	}
	for i := 1; i <= numWork; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= minProfit; k++ {
				if group[i-1] > j {
					state[i][j][k] = state[i-1][j][k]
				} else {
					state[i][j][k] = (state[i-1][j][k] + state[i-1][j-group[i-1]][misc.IntMax(0, k-profit[i-1])]) % mod
				}
			}
		}
	}
	return state[numWork][n][minProfit]
}

func ProfitableSchemes2(n int, minProfit int, group []int, profit []int) int {
	mod := int(1e9 + 7)
	numWork := len(group)
	state := make([][]int, n+1)
	for i := range state {
		state[i] = make([]int, minProfit+1)
		state[i][0] = 1
	}
	for i := 0; i < numWork; i++ {
		for j := n; j >= group[i]; j-- {
			for k := minProfit; k >= 0; k-- {
				state[j][k] = (state[j][k] + state[j-group[i]][misc.IntMax(0, k-profit[i])]) % mod
			}
		}
	}
	return state[n][minProfit]
}
