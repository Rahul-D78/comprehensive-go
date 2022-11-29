package dsa

import "fmt"

// Node represents the components of a binary search tree
type bst struct {
	key   int
	left  *bst
	right *bst
}

// Insert will add a new node to the tree
// The key to add and we need to check that it should not present in tree
func (b *bst) insert(k int) {
	if b.key < k {
		// move right
		// If the child is empty we need to put the new node with the value as k
		if b.right == nil {
			b.right = &bst{key: k}
		} else {
			// we need to call insert method recursively
			// we calling this method on the right child
			b.right.insert(k)
		}
	} else if b.key > k {
		// move left
		if b.left == nil {
			b.left = &bst{key: k}
		} else {
			// use recursion
			b.left.insert(k)
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value exists
func (b *bst) search(k int) bool {
	if b == nil {
		return false
	}
	if b.key < k {
		// move left
		return b.right.search(k)
	} else if b.key > k {
		// move right
		return b.left.search(k)
	}
	return true
}

func Bst() {
	// make a tree use a struct litelal
	tree := &bst{key: 100}
	tree.insert(50)
	tree.insert(300)
	tree.insert(400)
	tree.insert(500)
	fmt.Println(tree)
	fmt.Println(tree.search(40))
}
