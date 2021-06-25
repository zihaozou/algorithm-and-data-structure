package algorithm

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue1: make([]int, 0), queue2: make([]int, 0)}
}

/** Push element x onto stack. */
func (stack *MyStack) Push(x int) {
	var old, new *[]int
	if len(stack.queue1) == 0 {
		stack.queue1 = append(stack.queue1, x)
		old = &stack.queue2
		new = &stack.queue1
	} else {
		stack.queue2 = append(stack.queue2, x)
		old = &stack.queue1
		new = &stack.queue2
	}
	for len(*old) != 0 {
		*new = append(*new, (*old)[0])
		*old = (*old)[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (stack *MyStack) Pop() int {
	var ret int
	if len(stack.queue1) != 0 {
		ret = stack.queue1[0]
		stack.queue1 = stack.queue1[1:]
	} else {
		ret = stack.queue2[0]
		stack.queue2 = stack.queue2[1:]
	}
	return ret
}

/** Get the top element. */
func (stack *MyStack) Top() int {
	if len(stack.queue1) != 0 {
		return stack.queue1[0]
	} else {
		return stack.queue2[0]
	}
}

/** Returns whether the stack is empty. */
func (stack *MyStack) Empty() bool {
	if len(stack.queue1) == 0 && len(stack.queue2) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
