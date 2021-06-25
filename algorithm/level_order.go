package algorithm

type LOTreeNode struct {
	Val   int
	Left  *LOTreeNode
	Right *LOTreeNode
}

func LevelOrder(root *LOTreeNode) [][]int {
	if root == nil {
		return nil
	}
	var size int
	orderArray := make([][]int, 0)
	queue := make([]*LOTreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		size = len(queue)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			temp[i] = queue[0].Val
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		orderArray = append(orderArray, temp)
	}
	return orderArray
}
