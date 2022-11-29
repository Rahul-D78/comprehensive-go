package dsa

import "fmt"

type stack struct {
	items []int
}

// push will add a value at the end
func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

// pop will remove a value from the top
func (s *stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func Stack() {
	mystack := stack{}
	fmt.Println(mystack)
	mystack.push(100)
	mystack.push(200)
	mystack.push(300)
	fmt.Println(mystack)
	mystack.pop()
	fmt.Println(mystack)
}
