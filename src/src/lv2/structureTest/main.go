package main

import (
	"dataStructure"
)

func main() {
	list := dataStructure.LinkedList{}
	for i := 0; i < 10; i++ {
		node := dataStructure.Node{Val: i}
		list.AddNode(&node)
	}

	list.Print()

	list.RemoveNode(list.GetNode(5))
	list.RemoveNode(list.GetNode(1))
	list.RemoveNode(list.GetNode(0))

	list.Print()

	dll := dataStructure.DoubleLinkedList{}
	for i := 0; i < 10; i++ {
		node := dataStructure.Node{Val: i}
		dll.AddNode(&node)
	}

	dll.Print()

	dll.RemoveNode(dll.GetNode(5))
	dll.RemoveNode(dll.GetNode(1))
	dll.RemoveNode(dll.GetNode(0))

	dll.Print()
}
