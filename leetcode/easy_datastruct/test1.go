package main

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	data int
	next *Node
}

func (ll *LinkedList) add(num int) {
	node := new(Node)
	node.data = num
	currentNode := ll.head
	if currentNode == nil {
		ll.head = node
	} else {
		for currentNode.next != nil { //找出最后一个节点
			currentNode = currentNode.next
		}
		currentNode.next = node
	}
	ll.length++
}
// 获取节点的值，需要遍历节点
func (ll *LinkedList) get(index int) (num int, err error) {
	if ll.length == 0 {
		err = errors.New("empty list")
		return
	}
	if index < 0 || index >= ll.length {
		err = errors.New("out of range")
		return
	}
	currentNode := ll.head
	for i := 0; i <= index; i++ {
		if i == index {
			break
		}
		currentNode = currentNode.next
	}
	return currentNode.data, nil
}
// 删除链表中的节点，传入节点的值
func (ll *LinkedList) remove(num int) (err error) {
	if ll.length == 0 {
		err = errors.New("empty list")
		return
	}
	var preNode *Node
	currentNode := ll.head
	for i := 0; i < ll.length; i++ {
		if currentNode.data == num {
			if i == 0 {
				ll.head = currentNode.next
			} else {
				preNode.next = currentNode.next
			}
			ll.length--
			break
		}
		preNode = currentNode
		currentNode = currentNode.next
	}
	return nil
}

func (ll *LinkedList) size() int {
	return ll.length
}
// 遍历链表的节点
func (ll *LinkedList) forNode() {
	if ll.head == nil {
		fmt.Println("链表为空")
		return
	}
	node := ll.head
	for node != nil {
		fmt.Print(node.data)
		node = node.next
	}
	fmt.Println()
}
// 反转链表
func (ll *LinkedList) reverse() (err error) {
	if ll.length == 0 {
		err = errors.New("empty list")
		return
	}
	if ll.head.next == nil {
		return
	}
	var cur, node, pre *Node
	for i := 0; i < ll.length; i++ {
		if i == 0 {
			cur = ll.head
			node = cur.next
			cur.next = nil
			pre = cur
		} else {
			cur = node
			node = node.next
			cur.next = pre
			pre = cur
		}
	}
	ll.head = cur
	return nil
}

func merge(head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}
	var head *Node
	if head1.data <= head2.data {
		head = head1
		head.next = merge(head1.next, head2)
	} else {
		head = head2
		head.next = merge(head1, head2.next)
	}
	return head
}

func (node *Node) showNode() {
	if node == nil {
		return
	}
	fmt.Print(node.data)
	node.next.showNode()
	return
}

// 1 2 2 1
// 将前一半链表逆置（反转），再与后一半链表比较
func (ll *LinkedList) isPalindrome() bool {
	if ll.head == nil || ll.head.next == nil {
		return true
	}
	var cur, node, pre *Node
	for i := 0; i < (ll.length / 2); i++ {
		if i == 0 {
			cur = ll.head
			node = cur.next
			cur.next = nil
			pre = cur
		} else {
			cur = node
			node = cur.next
			cur.next = pre
			pre = cur
		}
	}
	var head1, head2 *Node
	head1 = cur
	head2 = node
	if ll.length % 2 == 1 {
		head2 = node.next
	}
	for head1 != nil && head2 != nil {
		if head1.data != head2.data {
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}
	return true
}

func (ll *LinkedList) isCycleLinkedList() bool {
	head := ll.head
	if head == nil || head.next == nil {
		return false
	}
	n1 := head
	n2 := head
	n1 = n1.next
	n2 = n2.next.next
	for n1 != n2 && n2 != nil && n2.next != nil {
		n1 = n1.next
		n2 = n2.next.next
	}
	if n1 == n2 {
		return true
	}
	return false
}

func removeReplicationNode(head *Node) {
	node := head
	preNode := head
	m := make(map[int]*Node)
	for node != nil {
		if _, ok := m[node.data]; ok {
			node = node.next
		} else {
			m[node.data] = node
			if node == head {
				preNode = node
			} else {
				preNode.next = node
				preNode = node
				node = node.next
			}
		}
	}
	preNode.next = nil
}

func main() {
	var err error
	// 生成链表
	ll := new(LinkedList)
	arr := []int{1, 2, 3, 4}
	for _, num := range arr {
		ll.add(num)
	}
	ll.forNode()

	// 删除链表中的节点
	err = ll.remove(2)
	if err != nil {
		panic(err)
	}
	ll.forNode()

	// 反转链表
	err = ll.reverse()
	if err != nil {
		panic(err)
	}
	ll.forNode()

	// 合并两个有序链表
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := new(LinkedList)
	l2 := new(LinkedList)
	for _, num := range arr1 {
		l1.add(num)
	}
	for _, num := range arr2 {
		l2.add(num)
	}
	node := merge(l1.head, l2.head)
	node.showNode()
	fmt.Print("\n")
	// 请判断一个链表是否为回文链表
	l3 := new(LinkedList)
	arr3 := []int{1, 2, 3, 2, 1}
	for _, num := range arr3 {
		l3.add(num)
	}
	fmt.Println(l3.isPalindrome())
	// 判断环形链表
	fmt.Println(ll.isCycleLinkedList())
	// 生成链表
	l4 := new(LinkedList)
	nums := []int{1, 3, 1, 5, 8, 3}
	for _, num := range nums {
		l4.add(num)
	}
	l4.forNode()

	// 删除链表重复元素
	removeReplicationNode(l4.head)
	l4.forNode()
}
