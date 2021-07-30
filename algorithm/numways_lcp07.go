package algorithm

func NumWaysLCP07(n int, relation [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	table := make(map[int][]int, n)
	for i := range relation {
		if table[relation[i][1]] != nil {
			table[relation[i][1]] = append(table[relation[i][1]], relation[i][0])
		} else {
			table[relation[i][1]] = []int{relation[i][0]}
		}
	}
	for i := 1; i < k+1; i++ {
		for j := range dp[i] {
			for _, q := range table[j] {
				dp[i][j] += dp[i-1][q]
			}
		}
	}
	return dp[k][n-1]
}
