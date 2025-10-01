package dataStructure

import "fmt"

type stackNode[T any] struct {
	data T
	next *stackNode[T]
}
type Stack[T any] struct {
	front *stackNode[T]
	len   int
}

func (s *Stack[T]) push(data T) {
	newNode := &stackNode[T]{data: data}
	newNode.next = s.front
	s.front = newNode
	s.len++
}
func (s *Stack[T]) pop() (popNode *stackNode[T], ok bool) {
	if s.front == nil {
		return nil, false
	}
	popNode = s.front
	s.front = s.front.next
	popNode.next = nil
	s.len--
	return popNode, true
}
func (s *Stack[T]) show() {
	cur := s.front
	for cur != nil {
		fmt.Print(cur.data,",")
		cur =cur.next
	}
	fmt.Println()
}
func StackTest() {
	s := &Stack[int]{}

	fmt.Println("=== 1. 入栈 push [10,20,30] ===")
	for _, v := range []int{10, 20, 30} {
		fmt.Println("入栈:", v)
		s.push(v)
	}
	fmt.Print("当前栈内容: ")
	s.show() // 输出: 30,20,10,

	fmt.Println("=== 2. 再入栈 40 ===")
	s.push(40)
	fmt.Print("当前栈内容: ")
	s.show() // 输出: 40,30,20,10,

	fmt.Println("=== 3. 出栈 pop ===")
	if node, ok := s.pop(); ok {
		fmt.Println("出栈:", node.data)
	}
	fmt.Print("当前栈内容: ")
	s.show() // 输出: 30,20,10,

	fmt.Println("=== 4. 出栈直到空 ===")
	for {
		node, ok := s.pop()
		if !ok {
			break
		}
		fmt.Println("出栈:", node.data)
	}
	fmt.Print("当前栈内容: ")
	s.show() // 栈为空

	fmt.Println("=== 5. 出栈空栈测试 ===")
	if _, ok := s.pop(); !ok {
		fmt.Println("栈为空，无法出栈")
	}
}
