package main

import (
	"fmt"
	"sorting/algorithm"
)

func main() {
	/*RB Tree test
	rand.Seed(time.Now().UnixNano())
	tree := datastructure.NewRBTree()
	for _, x := range []int{1, 4, 5, 14, 26, 7, 39, 54, 9, 19, 29, 58, 40, 16} {
		node := datastructure.NewRBNode(x)
		tree.Insert(node)
	}
	tree.InorderPrint()
	RB Tree test*/
	/* get Intersect node
	nodei5 := algorithm.ListNode{Val: 5, Next: nil}
	nodei4 := algorithm.ListNode{Val: 4, Next: &nodei5}
	nodei8 := algorithm.ListNode{Val: 8, Next: &nodei4}
	nodeA1 := algorithm.ListNode{Val: 1, Next: &nodei8}
	nodeA4 := algorithm.ListNode{Val: 4, Next: &nodeA1}
	nodeB1 := algorithm.ListNode{Val: 1, Next: &nodei8}
	nodeB6 := algorithm.ListNode{Val: 6, Next: &nodeB1}
	nodeB5 := algorithm.ListNode{Val: 5, Next: &nodeB6}
	algorithm.GetIntersectionNode(&nodeA4, &nodeB5)
	*/
	/*
		fmt.Printf("%s", algorithm.LongestPalindrome("ababd"))
	*/
	/*
		node5 := algorithm.ReListNode{Val: 5, Next: nil}
		node4 := algorithm.ReListNode{Val: 4, Next: &node5}
		node3 := algorithm.ReListNode{Val: 3, Next: &node4}
		algorithm.ReverseList(&node3)
	*/
	step := algorithm.OpenLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	fmt.Printf("step: %v\n", step)
}
