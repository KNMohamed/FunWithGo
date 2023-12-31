package main

import (
	"fmt"
)

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	val T
}

func (l *List[T]) PushFront(val T) {
	newNode := &Node[T]{val: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}
}

func (l *List[T]) PushBack(val T) {
	newNode := &Node[T]{val: val}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

func (l *List[T]) PrintListForward() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%v -> ", curr.val)
		curr = curr.next
	}
	fmt.Println("nil")
}

func (l *List[T]) PrintListBackwards() {
	curr := l.tail
	for curr != nil {
		fmt.Printf("%v -> ", curr.val)
		curr = curr.prev
	}
	fmt.Println("nil")
}

//Inserts a node after the value T if it exsts in the list
func (l *List[T]) InsertAfter(node *Node[T], data T) {
	curr := l.head

	for curr != nil {
		next := curr.next
		if curr.val == data {
			if curr == l.tail {
				l.tail = node
			}
			curr.next = node
			node.prev = curr
			node.next = next
			if next != nil {
				next.prev = node
			}
			break
		}
		curr = curr.next
	}
}

func (l *List[T]) InsertBefore(node *Node[T], data T) {
	curr := l.tail
	for curr != nil {
		before := curr.prev
		if curr.val == data {
			if curr == l.head {
				l.head = node
			}
			node.next = curr
			curr.prev = node
			node.prev = before
			if before != nil {
				before.next = node
			}
			break
		} 
		curr = before
	}
}

func main() {
	l := &List[string]{}
	l.PushFront("is")
	l.PushBack("a")
	l.PushBack("test")

	newNode := &Node[string]{val: "this"}
	l.InsertBefore(newNode, "is")
	l.PrintListForward()
	
	l.PushBack("to")
	l.PushBack("determine")
	l.PushBack("the")
	l.PushBack("correct")
	l.PushBack("order")
	l.PrintListForward()

	newNode = &Node[string]{val: "print"}
	l.InsertAfter(newNode, "correct")
	l.PrintListForward()
	l.PrintListBackwards()
}