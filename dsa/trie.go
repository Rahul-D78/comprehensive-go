package dsa

import "fmt"

// alphabetSize is the number of possible charecters in a trie
const alphabetSize = 26

// Node represents each node in trie
type nodeTrie struct {
	children [alphabetSize]*nodeTrie
	isEnd    bool
}

// Trie represents trie and has a pointer pointing to the root
type trie struct {
	root *nodeTrie
}

// initTrie will create a new trie
func initTrie() *trie {
	result := &trie{root: &nodeTrie{}}
	return result
}

// Insert will take a word as argument and insert it into the trie
func (t *trie) insert(w string) {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// create a node get the pointer and insert into the index
			currentNode.children[charIndex] = &nodeTrie{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search method will return true if a given word is present in the trie
// This function is a method of struct trie
func (t *trie) search(w string) bool {
	wordLen := len(w)
	currentNode := t.root
	for i := 0; i < wordLen; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func Trie() {
	myTrie := initTrie()
	fmt.Println("--")
	myTrie.insert("hello")
	fmt.Println(myTrie.search("world"))
	fmt.Println("--")

}
