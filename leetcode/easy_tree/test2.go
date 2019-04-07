package main

import (
	"fmt"
	"strings"
)

//二叉树
//性质1：二叉树第i层上的结点数目最多为 2{i-1} (i≥1)。
//性质2：深度为k的二叉树至多有2{k}-1个结点(k≥1)。
//性质3：包含n个结点的二叉树的高度至少为log2 (n+1)，根据"性质2"计算
//性质4：在任意一棵二叉树中，若终端结点的个数为n0，度为2的结点数为n2，则n0=n2+1。
//性质5：若对含 n 个结点的完全二叉树从上到下且从左至右进行 1 至 n 的编号，则对完全二叉树中任意一个编号为 i 的结点：
//(1) 若 i=1，则该结点是二叉树的根，无双亲, 否则，编号为 [i/2] 的结点为其双亲结点;
//(2) 若 2i>n，则该结点无左孩子，否则，编号为 2i 的结点为其左孩子结点;
//(3) 若 2i+1>n，则该结点无右孩子结点，否则，编号为2i+1 的结点为其右孩子结点。

type Tree2 struct {
	root *Node2
}

type Node2 struct {
	Data  string //rune类型默认值为0
	Left  *Node2
	Right *Node2
}
// 这里用到性质5
func initTree2(data []string) *Tree2 {
	t := new(Tree2)
	node := make([]*Node2, len(data))
	for i, v := range data {
		n := &Node2{
			Data: v,
		}
		node[i] = n
	}
	t.root = node[0]
	for i := 0; i < len(node)/2; i++ {
		if node[i*2+1].Data != "" {
			node[i].Left = node[i*2+1]
		}
		if i*2+2 < len(node) {
			if node[i*2+2].Data != "" {
				node[i].Right = node[i*2+2]
			}
		}
	}
	return t
}

// 在一个二叉树中找到所有节点的内容包含某子串的节点
// 并使用快速排序方法找到顺序为第n位的节点值。
// 排序规则如下：子串出现的次数，如果次数一样则按字符数排序，如果字符数一样则按ascii排序
func Inorder(node *Node2, substr string, retNode *[]*Node2) {
	if node != nil {
		Inorder(node.Left, substr, retNode)
		fmt.Print(node.Data + " ")
		if strings.Contains(node.Data, substr) {
			*retNode = append(*retNode, node)
		}
		Inorder(node.Right, substr, retNode)
	}
}

func QuickSort(nodes []*Node2, substr string, left int, right int) []*Node2 {
	value := nodes[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && !isSmall(value, nodes[j], substr) {
			j--
		}
		if j >= p {
			nodes[p] = nodes[j]
			p = j
		}
		for i <= p && isSmall(value, nodes[i], substr) {
			i++
		}
		if i <= p {
			nodes[p] = nodes[i]
			p = i
		}
	}
	nodes[p] = value
	if p-left > 1 {
		QuickSort(nodes, substr, left, p-1)
	}
	if right-p > 1 {
		QuickSort(nodes, substr, p+1, right)
	}
	return nodes
}

func isSmall(p *Node2, n *Node2, substr string) bool {
	pc := strings.Count(p.Data, substr)
	nc := strings.Count(n.Data, substr)
	if pc > nc {
		return true
	} else if pc < nc {
		return false
	} else {
		pstr := p.Data
		nstr := n.Data
		if len(pstr) > len(nstr) {
			return true
		} else if len(pstr) < len(nstr) {
			return false
		} else {
			pr := []rune(pstr)
			nr := []rune(nstr)
			for i, _ := range pr {
				pascii := pr[i]
				nascii := nr[i]
				if pascii > nascii {
					return true
				} else {
					return false
				}
			}
		}
	}
	return false
}

func main() {
	c := []string{"apple", "boy", "car", "driver", "else", "fun", "good", "helloworld"}
	t := initTree2(c)
	retNode := make([]*Node2, 0)
	Inorder(t.root, "o", &retNode)
	fmt.Println()
	for _, v := range retNode { //helloworld boy good
		fmt.Print(v.Data + " ")
	}
	fmt.Println()
	nodes := QuickSort(retNode, "o", 0, len(retNode)-1)
	for _, v := range nodes { //boy good helloworld
		fmt.Print(v.Data + " ")
	}
	fmt.Println()

}