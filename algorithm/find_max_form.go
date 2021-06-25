package algorithm

import (
	"sorting/misc"
	"strings"
)

//背包问题变种
/*设k为字符串集strs的前k项字符串，i为0的数量，j为1的数量
state[k][i][j]为在前k个字符串中，可以放下最多的子字符串的个数，且其中的0不超过i，
1不超过j。
*/
func FindMaxForm(strs []string, m int, n int) int {
	state := make([][][]int, len(strs)+1)
	for i := range state {
		state[i] = make([][]int, m+1)
		for j := range state[i] {
			state[i][j] = make([]int, n+1)
		}
	}
	for k := 1; k < len(strs)+1; k++ {
		zerocnt := strings.Count(strs[k-1], "0")
		onecnt := len(strs[k-1]) - zerocnt
		for i := 0; i < m+1; i++ {
			for j := 0; j < n+1; j++ {
				state[k][i][j] = state[k-1][i][j]
				if zerocnt <= i && onecnt <= j {
					/*
						为什么要[i-zerocnt][j-onecnt]:
						在if中我们已经判断了当前的字符串是可以放入背包了，
						即满足了放入背包的基本要求：字符串中的0和1不超过总限制，
						那么我们就要选择要放入哪一个状态的背包中，
						i-zerocnt和j-onecnt确保了此状态的背包有足够的空间放入此字符串
					*/
					state[k][i][j] = misc.IntMax(state[k-1][i][j], state[k-1][i-zerocnt][j-onecnt])
				}
			}
		}
	}
	return state[len(strs)][m][n]
}

//滚动数组以减少空间维度
func FindMaxForm2(strs []string, m int, n int) int {
	state := make([][]int, m+1)
	for i := range state {
		state[i] = make([]int, n+1)
	}
	for _, v := range strs {
		zerocnt := strings.Count(v, "0")
		onecnt := len(v) - zerocnt
		for i := m; i >= 0; i++ {
			for j := n; j >= 0; j++ {
				state[i][j] = misc.IntMax(state[i][j], state[i-zerocnt][j-onecnt])
			}
		}
	}
	return state[m][n]
}
