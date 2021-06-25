package algorithm

func TwoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i, v := range nums {
		table[v] = i
	}
	for i, v := range nums {
		ret, ok := table[target-v]
		if ok {
			return []int{ret, i}
		}
	}
	return nil
}
