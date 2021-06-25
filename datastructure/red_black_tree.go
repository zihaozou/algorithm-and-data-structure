package datastructure

import "fmt"

type RBNode struct {
	parent *RBNode
	left   *RBNode
	right  *RBNode
	black  bool
	key    int
}

type RBTree struct {
	root    *RBNode
	nilnode RBNode
}

func NewRBTree() *RBTree {
	tree := new(RBTree)
	tree.nilnode.black = true
	tree.root = &tree.nilnode
	return tree
}

func NewRBNode(key int) *RBNode {
	return &RBNode{key: key}
}

func (tree *RBTree) InorderPrint() {
	tree.inorderPrintRecur(tree.root)
}

func (tree *RBTree) inorderPrintRecur(node *RBNode) {
	if node.left != &tree.nilnode {
		tree.inorderPrintRecur(node.left)
	}
	if node != &tree.nilnode {
		fmt.Printf("[%d]", node.key)
	}
	if node.right != &tree.nilnode {
		tree.inorderPrintRecur(node.right)
	}
}

func (tree *RBTree) Search(key int) (ret *RBNode) {
	now := tree.root
	for now != &tree.nilnode {
		if now.key == key {
			return now
		} else if now.key > key {
			now = now.left
		} else {
			now = now.right
		}
	}
	return nil
}

func (tree *RBTree) Predecessor(node *RBNode) *RBNode {
	ret := node.left
	if ret == &tree.nilnode {
		return nil
	}
	for ret.right != &tree.nilnode {
		ret = ret.right
	}
	return ret
}

func (tree *RBTree) Successor(node *RBNode) *RBNode {
	ret := node.right
	if ret == &tree.nilnode {
		return nil
	}
	for ret.left != &tree.nilnode {
		ret = ret.left
	}
	return ret
}

func (tree *RBTree) Minimum() (min *RBNode) {
	min = tree.root
	if min == &tree.nilnode {
		return nil
	}
	for min.left != &tree.nilnode {
		min = min.left
	}
	return
}

func (tree *RBTree) Maximun() (max *RBNode) {
	max = tree.root
	if max == &tree.nilnode {
		return nil
	}
	for max.right != &tree.nilnode {
		max = max.right
	}
	return
}

func (tree *RBTree) leftRotate(node *RBNode) {
	y := node.right
	temp := y.left
	if temp != &tree.nilnode {
		temp.parent = node
	}
	node.right = temp
	temp = node.parent
	node.parent = y
	y.left = node
	y.parent = temp
	if temp != &tree.nilnode {
		if temp.right == node {
			temp.right = y
		} else {
			temp.left = y
		}
	} else {
		tree.root = y
	}
}

func (tree *RBTree) rightRotate(node *RBNode) {
	y := node.left
	temp := y.right
	if temp != &tree.nilnode {
		temp.parent = node
	}
	node.left = temp
	temp = node.parent
	node.parent = y
	y.right = node
	y.parent = temp
	if temp != &tree.nilnode {
		if temp.right == node {
			temp.right = y
		} else {
			temp.left = y
		}
	} else {
		tree.root = y
	}
}

func (tree *RBTree) Insert(node *RBNode) {
	now := tree.root
	var parent *RBNode = &tree.nilnode
	for now != &tree.nilnode {
		parent = now
		if now.key > node.key {
			now = now.left
		} else {
			now = now.right
		}
	}
	node.black = false
	node.parent = parent
	if parent != &tree.nilnode {
		if parent.key > node.key {
			parent.left = node
		} else {
			parent.right = node
		}
	} else {
		tree.root = node
	}
	node.right = &tree.nilnode
	node.left = &tree.nilnode
	tree.insertFixUp(node)
}

/*插入修复步骤
首先，插入过程可能会违背三个定义
定义2:根节点一定为黑色
定义4：不能有两个连续的红节点
1，首先插入修复起始于插入的节点
2，插入修复的主循环条件为：当前修复节点的父节点不是黑色
3，进入主循环，首先判断z的父节点是左子节点还是右子节点
左子节点：
修复的方式为：
case1.叔节点和父节点同为红色（则祖父节点一定为黑色）
	设置叔节点和父节点的颜色为黑色，祖父节点颜色为红色
	并设置z为原z的祖父节点，这一步其实已经解决了定义4，
	但是有可能调换的颜色会产生新的定义4冲突，我们进入循环检验
case2.叔节点与父节点不为同色，叔节点为黑色，同时我们也确定了祖父节点一定是黑色
	在这里，我们的做法是确保z处于父节点的左侧，然后调换父节点和祖父节点的颜色，
	则父节点为黑色，祖父节点为红色，然后左旋，此时修复了定义4的冲突
在最后会将根节点设置为黑色，修复定义2
右子节点同理
*/
func (tree *RBTree) insertFixUp(node *RBNode) {
	z := node
	for !z.parent.black {
		if z.parent.parent.left == z.parent { //left
			uncle := z.parent.parent.right
			switch {
			case !uncle.black:
				z.parent.black = true
				uncle.black = true
				z.parent.parent.black = false
				z = z.parent.parent
			case z == z.parent.right:
				z = z.parent
				tree.leftRotate(z)
				fallthrough
			case z == z.parent.left:
				z.parent.black = true
				z.parent.parent.black = false
				tree.rightRotate(z.parent.parent)
			}
		} else {
			uncle := z.parent.parent.left
			switch {
			case !uncle.black:
				z.parent.black = true
				uncle.black = true
				z.parent.parent.black = false
				z = z.parent.parent
			case z == z.parent.left:
				z = z.parent
				tree.rightRotate(z)
				fallthrough
			case z == z.parent.right:
				z.parent.black = true
				z.parent.parent.black = false
				tree.leftRotate(z.parent.parent)
			}
		}
	}
	tree.root.black = true
}

/*
红黑树的移植与传统二叉树的移植的不同是，不论v是否为nilnode，
都将v的父节点设置为u的父节点
*/
func (tree *RBTree) transplate(u, v *RBNode) {
	if u.parent == &tree.nilnode {
		tree.root = v
	} else if u.parent.left == u {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

/*
删除操作
删除分为三种情况：
1&2，要删除的节点只有左子节点或右子节点，这种情况只需
	删除节点后将唯一的子节点顶替被删除的节点的位置
3，要删除的节点有两个子节点，这里比较tricky，我们
	找到需要删除的节点的后继者，此后继者必定最多只
	有一个非nil子节点，我们将此后继者的值拷贝到要删除
	的节点内，然后删除后继者节点，通过此方法，我们只需要
	删除一个只有一个子节点的节点，所以删除方法和1&2相同
删除完毕后，我们需要调用fixup函数修复我们删除后可能引起的
红黑冲突，注意：只有当删除掉的节点为黑色时才需要调用，因为如果
删除的节点是红色那么一定不会引起冲突
*/
func (tree *RBTree) Delete(node *RBNode) {
	//y为将要被删除的节点，x为y的唯一子节点（x也可以为nilnode）
	var y, x *RBNode
	var yOrgnBlk bool
	y = node
	yOrgnBlk = node.black
	if node.left == &tree.nilnode {
		x = node.right
		tree.transplate(node, x)
	} else if node.right == &tree.nilnode {
		x = node.left
		tree.transplate(node, x)
	} else {
		y = tree.Successor(node)
		yOrgnBlk = y.black
		node.key = y.key
		x = y.right
		tree.transplate(y, x)
	}
	if yOrgnBlk {
		tree.deleteFixUp(x)
	}

}

/*
冲突调和：
在这里，我们假设被删掉的节点所持有的颜色被转移到了x节
点上，则x既有原本的颜色，也有被删掉的黑色（我们确保
被删掉的节点是黑色才会进入此函数）
可能引起冲突的定义为1，2，4
我们调和的方式为：首先，如果x节点本来是红色我们可以
直接将x变黑即可。如果x是黑色，我们则将多出的黑色属性
向上传递，知道遇到红色的节点，我们就将此节点变黑，或
者一路传递到根节点，则整个树都少了一个黑色节点，我们
就将此黑色抛弃。
case1：是为了像case2，3，4转换。将兄弟节点通过旋转变黑
case2: 是为了向上传递黑色，在兄弟节点的分径上制造一个黑色
	空缺，然后将x设为x的父节点，即x分径和兄弟分径都需要多一个
	黑色，这个任务就交给了x的父节点
case3&case4: 确保兄弟节点的右子节点是红色，兄弟节点的将右子
	节点变黑。交换x的父节点和兄弟节点的颜色，对父节点作左旋。
	此时x所需的黑色由父节点承担，兄弟节点左子节点的黑色由x的
	父节点承担，兄弟节点的右子节点所需的黑色由其自己承担，至此，定义
	1的冲突就被解决了
*/
func (tree *RBTree) deleteFixUp(x *RBNode) {
	for x != tree.root && !x.black {
		if x == x.parent.left {
			w := x.parent.right
			switch {
			case !w.black:
				x.parent.black = false
				w.black = true
				tree.leftRotate(x.parent)
			case w.left.black && w.right.black:
				w.black = false
				x = x.parent
			case w.right.black:
				w.black = false
				w.left.black = true
				tree.rightRotate(w)
				fallthrough
			default:
				temp := w.black
				w.black = x.parent.black
				x.parent.black = temp
				w.right.black = true
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else {
			w := x.parent.left
			switch {
			case !w.black:
				x.parent.black = false
				w.black = true
				tree.rightRotate(x.parent)
			case w.right.black && w.left.black:
				w.black = false
				x = x.parent
			case w.left.black:
				w.black = false
				w.right.black = true
				tree.leftRotate(w)
				fallthrough
			default:
				temp := w.black
				w.black = x.parent.black
				x.parent.black = temp
				w.right.black = true
				tree.rightRotate(x.parent)
				x = tree.root
			}
		}
	}
	x.black = true
}
