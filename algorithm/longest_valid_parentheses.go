package algorithm

import "sorting/misc"

func LongestValidParentheses(s string) int {
	var max int
	dp := make([]int, len(s)) //dp[i]子字符串内以i结尾最长有效括号
	for i := 1; i < len(s); i++ {
		if s[i] == ')' && (s[i-1] == '(' || dp[i-dp[i-1]-1] == '(') {
			dp[i] = dp[i-1] + 2
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func LongestValidParentheses2(s string) int {
	var max int
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i, v := range s {
		if v == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = misc.IntMax(max, i-stack[len(stack)-1]+1)
			}
		}
	}
	return max
}
