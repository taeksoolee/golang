package dataStructure

import "fmt"

type DoubleLinkedList struct {
	root *Node
	tail *Node
}

func (l *DoubleLinkedList) AddNode(node *Node) {
	node.prev = l.tail

	if l.root == nil {
		// Node가 없을 경우
		l.root = node
		l.tail = l.root
	} else {
		l.tail.next = node
		l.tail = l.tail.next
	}
}

func (l *DoubleLinkedList) RemoveNode(node *Node) {
	if l.root == node {
		l.root = node.next
		if l.tail == node {
			l.tail = node.next
		}
		return
	}

	if node.next == nil {
		// 선택 노드가 마지막 요소일 경우
		node.prev.next = nil
		l.tail = node.prev
	} else {
		node.prev.next = node.prev.next.next
	}
}

func (l *DoubleLinkedList) GetNode(index int) *Node {
	node := l.root
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *DoubleLinkedList) Print() {
	if l.tail == nil {
		fmt.Println("not find node")
		return
	}

	temp := l.root
	for temp != nil {
		fmt.Printf("%d", temp.Val)
		temp = temp.next
		if temp != nil {
			fmt.Print(" => ")
		}
	}
	fmt.Println("")
}
