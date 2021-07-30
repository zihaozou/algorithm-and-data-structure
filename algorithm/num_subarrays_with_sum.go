package algorithm

func NumSubarraysWithSum(nums []int, goal int) int {
	var l1, l2, cnt, s1, s2 int
	for r, v := range nums {
		s1 += v
		s2 += v
		for s1 > goal && l1 <= r {
			s1 -= nums[l1]
			l1++
		}
		for s2 >= goal && l1 <= r {
			s2 -= nums[l2]
			l2++
		}
		cnt += l2 - l1
	}
	return cnt
}
