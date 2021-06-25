package datastructure

type NumArray struct {
	tree []int
}

func Constructor(nums []int) NumArray {

	ret := NumArray{tree: make([]int, 2*len(nums))}
	constructorRecur(&nums, &ret.tree, 0, len(nums)-1, 1)
	return ret
}

func constructorRecur(nums *[]int, tree *[]int, l, r, i int) (ret int) {
	if r == l {
		(*tree)[i] = (*nums)[r]
		return (*nums)[r]
	}
	(*tree)[i] = constructorRecur(nums, tree, l, (r+l)/2, i+1) + constructorRecur(nums, tree, (r+l)/2+1, r, i+r-l+2-((r^l)&1))
	return (*tree)[i]
}

func (this *NumArray) Update(index int, val int) {
	this.updateRecur(index, 0, len(this.tree)/2-1, val, 1)
}
func (this *NumArray) updateRecur(i, l, r, v, t int) (ret int) {
	if r == l {
		ret = v - this.tree[t]
		this.tree[t] = v
		return
	} else if i <= (r+l)/2 {
		ret = this.updateRecur(i, l, (r+l)/2, v, t+1)
	} else {
		ret = this.updateRecur(i, (r+l)/2+1, r, v, t+r-l+2-((r^l)&1))
	}
	this.tree[t] += ret
	return
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sumRangeRecur(left, right, 0, len(this.tree)/2-1, 1)
}
func (this *NumArray) sumRangeRecur(tl, tr, l, r, i int) (ret int) {
	if tr < tl {
		return 0
	} else if (tl == l) && (tr == r) {
		return this.tree[i]
	}
	m := (r + l) / 2
	if m < tl {
		return this.sumRangeRecur(tl, tr, m+1, r, i+r-l+2-((r^l)&1))
	} else if m > tr {
		return this.sumRangeRecur(tl, tr, l, m, i+1)
	} else {
		return this.sumRangeRecur(tl, m, l, m, i+1) + this.sumRangeRecur(m+1, tr, m+1, r, i+r-l+2-((r^l)&1))
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
