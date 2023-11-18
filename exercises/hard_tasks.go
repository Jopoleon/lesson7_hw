package main

import (
	"fmt"
)

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func printLinkedList(head *Node) {
	for head != nil {
		fmt.Print(head.Value, " ")
		head = head.Next
	}
	fmt.Println()
}

// Вопрос: Как будет выглядеть связный список после выполнения функции reverseLinkedList?
func linkedList() {
	head := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next:  nil,
			},
		},
	}
	fmt.Println("Original List:")
	printLinkedList(head)

	head = reverseLinkedList(head)
	fmt.Println("Reversed List:")
	printLinkedList(head)
}

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func deepCopy(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newNode := &TreeNode{Value: root.Value}
	newNode.Left = deepCopy(root.Left)
	newNode.Right = deepCopy(root.Right)
	return newNode
}

func printTreePreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	printTreePreOrder(node.Left)
	printTreePreOrder(node.Right)
}

// Вопрос: Что произойдет после выполнения deepCopy?
func tree() {
	root := &TreeNode{Value: 1, Left: &TreeNode{Value: 2}, Right: &TreeNode{Value: 3}}
	copy := deepCopy(root)

	fmt.Println("Original Tree:")
	printTreePreOrder(root)
	fmt.Println("\nCopied Tree:")
	printTreePreOrder(copy)
}

func modifySlice(slice []*int) {
	temp := 5
	slice[0] = &temp
	slice = append(slice, &temp)
}

// Вопрос: Каков будет результат выполнения программы?
func slice() {
	a, b := 1, 2
	slice := []*int{&a, &b}
	modifySlice(slice)
	fmt.Println(*slice[0], *slice[1])
}

func hardTasks() {
	linkedList()

	// tree()
	//
	// slice()
}
