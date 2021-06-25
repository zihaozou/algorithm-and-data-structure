package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != nil || pB != nil {
		if pA == nil {
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
		if pA == pB {
			return pA
		} else {
			pA = pA.Next
			pB = pB.Next
		}
	}
	return nil
}
