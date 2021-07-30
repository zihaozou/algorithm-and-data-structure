package main

import "sorting/algorithm"

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
	/*
		data := make([][]int, 6)
		data[0] = []int{0, 1}
		data[1] = []int{0, 2}
		data[2] = []int{2, 1}
		data[3] = []int{1, 2}
		data[4] = []int{1, 0}
		data[5] = []int{2, 0}
		algorithm.NumWaysLCP07(3, data, 5)
	*/
	edges := make([][]int, 6)
	edges[0] = []int{1, 2}
	edges[1] = []int{1, 3}
	edges[2] = []int{1, 7}
	edges[3] = []int{2, 4}
	edges[4] = []int{2, 6}
	edges[5] = []int{3, 5}
	algorithm.FrogPosition(7, edges, 2, 4)
}
