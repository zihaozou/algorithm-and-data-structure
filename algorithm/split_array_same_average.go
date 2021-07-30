package algorithm

import "sorting/sort"

func SplitArraySameAverage(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return false
	}
	var sum int
	for i := range nums {
		nums[i] = nums[i] * length
		sum += nums[i]
	}
	avg := sum / length
	for i := range nums {
		nums[i] = nums[i] - avg
	}
	left := enumBin(nums[:length/2])
	for i := 1; i < len(left); i++ {
		if left[i] == 0 {
			return true
		}
	}
	left = left[1 : len(left)-1]
	right := enumBin(nums[length/2+1:])
	for i := 1; i < len(right); i++ {
		if right[i] == 0 {
			return true
		}
	}
	right = right[1 : len(right)-1]
	sort.QuickSort(&left)
	sort.QuickSort(&right)
	l, r := 0, len(right)-1
	for l < len(left) && r >= 0 {
		if left[l]+right[r] == 0 {
			return true
		} else if left[l]+right[r] > 0 {
			r--
		} else {
			l++
		}
	}
	return false
}

func enumBin(lst []int) []int { //二进制枚举，主要用于枚举一个集合的子集
	l := len(lst)
	ret := make([]int, 1<<l) //总共有2^l个子集
	for i := 0; i < l; i++ { //从第0位开始一个个纳入考虑范围
		for j := 0; j < 1<<i; j++ { //例如i=3时，j<100(1<<3),此时j会遍历00,01,10,11
			ret[j+(1<<i)] = lst[i] + ret[j] //i=3的话，则会赋值100,101,110,111
		}
	}
	return ret
}
