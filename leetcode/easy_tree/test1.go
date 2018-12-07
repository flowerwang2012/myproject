package main

import (
	"fmt"
	"math"
)

type Tree struct {
	root *Node
}
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

func newNode(data interface{}) *Node {
	node := new(Node)
	node.data = data
	return node
}
// 完全二叉树
func initTree(datas []interface{}) *Tree {
	t := new(Tree)
	nodes := make([]*Node, 0)
	for _, data := range datas {
		node := newNode(data)
		nodes = append(nodes, node)
	}
	t.root = nodes[0]
	for i := 0; i < len(nodes)/2; i++ {
		if nodes[i*2+1].data != nil {
			nodes[i].left = nodes[i*2+1]
		}
		if i*2+2 < len(nodes) {
			if nodes[i*2+2].data != nil {
				nodes[i].right = nodes[i*2+2]
			}
		}
	}
	return t
}
// 二叉搜索树
func (t *Tree) add(node *Node) {
	if t.root == nil {
		t.root = node
	} else {
		cur := t.root
		for cur != nil {
			if cur.data.(int) > node.data.(int) {
				if cur.left == nil {
					cur.left = node
					return
				}
				cur = cur.left
			} else {
				if cur.right == nil {
					cur.right = node
					return
				}
				cur = cur.right
			}
		}
	}
}
// 先序遍历
func preorder(root *Node) {
	if root != nil {
		fmt.Print(root.data)
		fmt.Print(" ")
		preorder(root.left)
		preorder(root.right)
	}
}
// 中序遍历
func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Print(root.data)
		fmt.Print(" ")
		inorder(root.right)
	}
}
// 后序遍历
func afterorder(root *Node) {
	if root != nil {
		afterorder(root.left)
		afterorder(root.right)
		fmt.Print(root.data)
		fmt.Print(" ")
	}
}
// 最大深度
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	} else {
		leftDepth := maxDepth(root.left)
		rightDepth := maxDepth(root.right)
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	}
}
var (
	firstNode = true
	lastVal = math.MaxInt32
)
// 验证二叉搜索树
func isValidBST(root *Node) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.left) {
		return false
	}
	if !firstNode && lastVal >= root.data.(int) {
		return false
	}
	// 此时 root.val>=lastval 是右子树
	firstNode = false
	lastVal = root.data.(int)
	if !isValidBST(root.right) {
		return false
	}
	return true
}
// 二叉树的层次遍历 借助队列来实现。相当于广度优先搜索，使用队列（深度优先搜索的话，使用栈）
func levelOrder(root *Node) {
	queue := make([]*Node, 0)
	if root == nil {
		return
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		nextQueue := make([]*Node, 0)
		for _, n := range queue {
			if n.left != nil {
				nextQueue = append(nextQueue, n.left)
			}
			if n.right != nil {
				nextQueue = append(nextQueue, n.right)
			}
			fmt.Print(n.data)
			fmt.Print(" ")
		}
		fmt.Println()
		queue = nil
		queue = append(queue, nextQueue...)
	}
}
func main() {
	nums := []interface{}{3, 9, 20, nil, nil, 15, 7}
	t := initTree(nums)
	preorder(t.root)
	fmt.Println()
	inorder(t.root)
	fmt.Println()
	afterorder(t.root)
	fmt.Println()
	fmt.Println(maxDepth(t.root))

	nums2 := []interface{}{2, 1, 3, 4}
	t2 := new(Tree)
	for _, num := range nums2 {
		node := newNode(num)
		t2.add(node)
	}
	preorder(t2.root)
	fmt.Println()
	nums3 := []interface{}{3, 2, 15, nil, nil, 7, 20}
	t3 := initTree(nums3)
	inorder(t3.root)
	fmt.Println()
	fmt.Println(isValidBST(t3.root))
	levelOrder(t3.root)
}
