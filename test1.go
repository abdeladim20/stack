package main

import "fmt"

type Stack struct {
	result []int
}

func (s *Stack) Push(i int) {
	s.result = append(s.result, i)
}

func (s *Stack) Peek() (int, bool) {
	if len(s.result) == 0 {
		return 0, false
	}
	return s.result[len(s.result)-1], true
}

func (s *Stack) Pop() (int, bool) {
	if len(s.result) == 0 {
		return 0, false
	}
	lastindex := len(s.result) - 1
	result := s.result[lastindex]
	s.result = s.result[:lastindex]
	return result, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.result) == 0
}

func search(stack []int, element int) int {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == element {
			return len(stack) - i
		}
	}
	return -1
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack)
	element, _ := stack.Pop()
	fmt.Println(element)
	fmt.Println()
	
	// searching an element
	fmt.Println(search(stack.result, 9))

		// Peek element
	if top, ok := stack.Peek(); ok {
		fmt.Println("Peeked Element : ", top)
	}

		// Pop elements from the stack
	for !stack.IsEmpty() {
		if result, ok := stack.Pop(); ok {
			fmt.Println("Pooped element : ", result)
		}
	}

		// Try to pop from an empty stack
	if _, ok := stack.Pop(); !ok {
		fmt.Println("Stack is empty")
	}

}
