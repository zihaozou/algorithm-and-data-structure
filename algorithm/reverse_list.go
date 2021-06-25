package algorithm

type ReListNode struct {
	Val  int
	Next *ReListNode
}

func ReverseList(head *ReListNode) *ReListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
