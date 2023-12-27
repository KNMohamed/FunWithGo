package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	head *Node[T]
}

type Node[T comparable] struct {
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

func (l *List[T]) Remove(val T) {
	//Case when list is empty
	if l.head == nil {
		return
	}
	//Case when found at beginning of list
	if l.head.val == val {
		l.head = l.head.next
		return
	}	
	curr := l.head
	for curr.next != nil {
		if curr.next.val == val {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}

func main() {
	l := &List[string]{}
	l.Append("hello")
	l.Append("world")
	l.PrintList()
	l.Remove("hello")
	
	l.Append("this")
	l.Append("is")
	l.Append("a")
	l.Append("test")
	l.Remove("a")
	l.PrintList()
	
	l2 := &List[string]{}
	l2.PrintList()
}