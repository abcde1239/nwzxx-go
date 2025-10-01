package dataStructure

import "fmt"

type queueNode[T any] struct {
	data T
	next *queueNode[T]
}
type Queue[T any] struct {
	head *queueNode[T]
	tail *queueNode[T]
	len  int
}

func (q *Queue[T]) enqueue(data T) {
	newNode := &queueNode[T]{data: data}
	if q.head == nil {
		// 队列为空
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.len++
}

func (q *Queue[T]) dequeue() (dequeueNode *queueNode[T], ok bool) {
	if q.head == nil {
		return nil, false
	}
	popNode := q.head
	q.head = q.head.next
	if q.head == nil {
		// 队列为空，尾指针置空
		q.tail = nil
	}
	popNode.next = nil
	q.len--
	return popNode, true
}
func (q *Queue[T]) show() {
	cur := q.head
	for cur != nil {
		fmt.Print(cur.data,",")
		cur =cur.next
	}
fmt.Println()
}
func QueueTest() {
	q := &Queue[int]{}

	fmt.Println("=== 1. 批量入队 [10,20,30] ===")
	for _, v := range []int{10, 20, 30} {
		q.enqueue(v)
	}
	q.show() // 输出: 10,20,30,

	fmt.Println("=== 2. 入队 40 ===")
	q.enqueue(40)
	q.show() // 输出: 10,20,30,40,

	fmt.Println("=== 3. 出队一次 ===")
	if node, ok := q.dequeue(); ok {
		fmt.Println("出队:", node.data)
	}
	q.show() // 输出: 20,30,40,

	fmt.Println("=== 4. 出队直到空 ===")
	for {
		node, ok := q.dequeue()
		if !ok {
			break
		}
		fmt.Println("出队:", node.data)
	}
	q.show() // 队列为空

	fmt.Println("=== 5. 出队空队列测试 ===")
	if _, ok := q.dequeue(); !ok {
		fmt.Println("队列为空，无法出队")
	}
	 println()
}
