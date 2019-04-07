package main

import (
	"fmt"
	"math"
	"container/list"
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
func initBSTree(datas []interface{}) *Tree {
	t := new(Tree)
	nodes := make([]*Node, len(datas))
	for i, v := range datas {
		n := newNode(v)
		nodes[i] = n
	}
	t.root = nodes[0]
	for i := 1; i < len(nodes); i++ {
		cur := t.root
		node := nodes[i]
		for cur != nil {
			if cur.data.(int) > node.data.(int) {
				if cur.left == nil {
					cur.left = node
					break
				}
				cur = cur.left
			} else {
				if cur.right == nil {
					cur.right = node
					break
				}
				cur = cur.right
			}
		}
	}
	return t
}
//先序遍历 先根后左最后右
//递归
func preorder(root *Node) {
	if root != nil {
		fmt.Print(root.data)
		fmt.Print(" ")
		preorder(root.left)
		preorder(root.right)
	}
}
//非递归 解决思路：假设只有左子树的情况，如何遍历，其他两种遍历也是参考这个思路
func preorder2(root *Node) {
	t := root
	l := list.New()
	for t != nil || l.Len() != 0 {
		for t != nil {
			fmt.Print(t.data)
			fmt.Print(" ")
			l.PushBack(t)
			t = t.left
		}
		if l.Len() != 0 {
			v := l.Back()
			t = v.Value.(*Node)
			t = t.right
			l.Remove(v)
		}
	}
}
//中序遍历 先左后根最后右
//递归
func inorder(root *Node) {
	if root != nil {
		inorder(root.left)
		fmt.Print(root.data)
		fmt.Print(" ")
		inorder(root.right)
	}
}
//非递归
func inorder2(root *Node) {
	t := root
	l := list.New()
	for t != nil || l.Len() != 0 {
		for t != nil {
			l.PushBack(t)
			t = t.left
		}
		if l.Len() != 0 {
			v := l.Back()
			t = v.Value.(*Node)
			fmt.Print(t.data)
			fmt.Print(" ")
			t = t.right
			l.Remove(v)
		}
	}
}
//后序遍历 先左后右最后根
//递归
func afterorder(root *Node) {
	if root != nil {
		afterorder(root.left)
		afterorder(root.right)
		fmt.Print(root.data)
		fmt.Print(" ")
	}
}
//非递归
func afterorder2(root *Node) {
	t := root
	l := list.New()
	var previsited *Node //访问过的左右节点
	for t != nil || l.Len() != 0 {
		for t != nil {
			l.PushBack(t)
			t = t.left
		}
		v := l.Back()
		top := v.Value.(*Node)

		if (top.left == nil && top.right == nil) || (top.right == nil && top.left == previsited) || top.right == previsited {
			fmt.Print(top.data)
			fmt.Print(" ")
			previsited = top
			l.Remove(v)
		} else {
			t = top.right
		}
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
	t2 := initBSTree(nums2)
	preorder(t2.root)
	fmt.Println()
	inorder(t2.root)
	fmt.Println()
	nums3 := []interface{}{3, 2, 15, nil, nil, 7, 20}
	t3 := initTree(nums3)
	inorder(t3.root)
	fmt.Println()
	fmt.Println(isValidBST(t3.root))
	levelOrder(t3.root)

	nums4 := []interface{}{0, 1, 2, 3, 4, 5, 6}
	t4 := initTree(nums4)
	preorder(t4.root)
	fmt.Println()
	preorder2(t4.root)
	fmt.Println()
	inorder(t4.root)
	fmt.Println()
	inorder2(t4.root)
	fmt.Println()
	afterorder(t4.root)
	fmt.Println()
	afterorder2(t4.root)
}
