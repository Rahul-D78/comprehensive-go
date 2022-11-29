package dsa

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

func LinkedList() {
	mylist := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	fmt.Println(mylist)
}
