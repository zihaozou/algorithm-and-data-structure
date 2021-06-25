package algorithm

func LongestPalindrome(s string) string {
	length := len(s)
	state := make([][]bool, length)
	for i := range state {
		state[i] = make([]bool, length)
	}
	var l int
	var r int
	for i := length - 1; i >= 0; i-- {
		for j := i; j <= length-1; j++ {
			if i == j {
				state[i][j] = true
			} else if s[i] == s[j] {
				if j-i == 1 {
					state[i][j] = true
				} else {
					state[i][j] = state[i+1][j-1]
				}
			}
			if state[i][j] && j-i > r-l {
				l = i
				r = j
			}
		}
	}
	return s[l : r+1]
}
