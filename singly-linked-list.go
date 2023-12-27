package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	head *Node[T]
}

type Node[T any] struct {
	next *Node[T]
	val T
}

func (l *List[T]) Append(val T) {
	newNode := &Node[T]{val: val}
	if l.head == nil	{
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}	
	curr.next = newNode
}

//Prints the elements of the list
func (l *List[T]) PrintList() {	
	curr := l.head
	for curr != nil {
		fmt.Printf("%v -> ", curr.val)
		curr = curr.next
	}
	
	fmt.Println("nil")
}

func main() {
	l := &List[string]{}
	l.Append("hello")
	l.Append("world")
	l.PrintList()
	
	l2 := &List[string]{}
	l2.PrintList()
}