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

// printing the list
func (l linkedList) printList() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

// Delete a specific node
// Handle corner cases such as:
// 1. If user give a value which is not present it will give null pointer dereference
// 2. If we have a empty linkedList it will give the error
// 3. If user wants to delete the head
func (l *linkedList) deleteWithValue(value int) {
	// handling the empty list case
	if l.length == 0 {
		return
	}

	// handling delete at head
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	// we are using the previous node because this is a singly linkedList
	// and we don't need pass through otherwise we are not able to access
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		// update previousToDelete
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
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
	mylist.printList()
	mylist.deleteWithValue(1)
	mylist.printList()
}
