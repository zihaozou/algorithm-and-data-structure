package datastructure

import (
	"fmt"
	"sorting/misc"
)

//AVL 树的左旋右旋与红黑树的左旋右旋相同
//aVL的左右旋就是先将x左旋，然后对x原本的父节点右旋
//avl的右左旋就是先将x右旋，然后对x原本的父节点左旋
/*
右旋：
		x
	  /
	y        ->			y
  /					  /   \
z					z		x

左旋：
x
  \
	y		 ->			y
	  \				  /	  \
	  	z			x		z
左右旋：
		x
	  /
	y			->			z
	  \					  /	  \
	  	z				y		x

右左旋：
		x
	  	  \
			y			->		z
	  	  /					  /	  \
	  	z					x		y
*/

type AVLNode struct {
	left   *AVLNode
	right  *AVLNode
	height int
	key    int
}

type AVLTree struct {
	root *AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{root: nil}
}

func NewAVLNode(key int) *AVLNode {
	return &AVLNode{key: key}
}

func (tree *AVLTree) InorderPrint() {
	tree.inorderPrintRecur(tree.root)
}

func (tree *AVLTree) inorderPrintRecur(node *AVLNode) {
	if node != nil {
		if node.left != nil {
			tree.inorderPrintRecur(node.left)
		}
		fmt.Printf("[%d]", node.key)
		if node.right != nil {
			tree.inorderPrintRecur(node.right)
		}

	}
}

func (tree *AVLTree) Search(key int) *AVLNode {
	now := tree.root
	for now != nil {
		if key < now.key {
			now = now.left
		} else if key > now.key {
			now = now.right
		} else {
			break
		}
	}
	return now
}

func (tree *AVLTree) Predecessor(node *AVLNode) *AVLNode {
	now := node.left
	for now != nil && now.right != nil {
		now = now.right
	}
	return now
}

func (tree *AVLTree) Successor(node *AVLNode) *AVLNode {
	now := node.right
	for now != nil && now.left != nil {
		now = now.left
	}
	return now
}

func (tree *AVLTree) Minimum() *AVLNode {
	now := tree.root
	for now != nil && now.left != nil {
		now = now.left
	}
	return now
}

func (tree *AVLTree) Maximum() *AVLNode {
	now := tree.root
	for now != nil && now.right != nil {
		now = now.right
	}
	return now
}

func (tree *AVLTree) balanceFactor(node *AVLNode) int8 {
	var lH, rH int
	if node == nil {
		return 0
	}
	if node.left != nil {
		lH = node.left.height
	}
	if node.right != nil {
		rH = node.right.height
	}
	return int8(lH - rH)
}

func (tree *AVLTree) height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (tree *AVLTree) leftRotate(x *AVLNode) *AVLNode {
	y := x.right
	x.right = y.left
	y.left = x
	x.height = misc.IntMax(tree.height(x.left), tree.height(x.right)) + 1
	y.height = misc.IntMax(tree.height(y.left), tree.height(y.right)) + 1
	if x == tree.root {
		tree.root = y
	}
	return y
}

func (tree *AVLTree) rightRotate(x *AVLNode) *AVLNode {
	y := x.left
	x.left = y.right
	y.right = x
	x.height = misc.IntMax(tree.height(x.left), tree.height(x.right)) + 1
	y.height = misc.IntMax(tree.height(y.left), tree.height(y.right)) + 1
	if x == tree.root {
		tree.root = y
	}
	return y
}

func (tree *AVLTree) Insert(node *AVLNode) {
	tree.insertRecur(tree.root, node)
}

func (tree *AVLTree) insertRecur(now *AVLNode, new *AVLNode) *AVLNode {
	if now == nil {
		return new
	} else if now.key > new.key {
		now.left = tree.insertRecur(now.left, new)
	} else if now.key < new.key {
		now.right = tree.insertRecur(now.right, new)
	} else {
		return now
	}
	now.height = 1 + misc.IntMax(tree.height(now.left), tree.height(now.right))
	balance := tree.balanceFactor(now)
	switch {
	case balance > 1 && new.key < now.left.key: //右旋
		return tree.rightRotate(now)
	case balance < -1 && new.key > now.right.key: //左旋
		return tree.leftRotate(now)
	case balance > 1 && new.key > now.left.key: //左右旋
		now.left = tree.leftRotate(now.left)
		return tree.rightRotate(now)
	case balance < -1 && new.key < now.right.key: //右左旋
		now.right = tree.rightRotate(now.right)
		return tree.leftRotate(now)
	default:
		return now
	}
}
