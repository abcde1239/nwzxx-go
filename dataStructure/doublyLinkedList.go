package dataStructure

import (
	"fmt"
)

type doublyLinkedNode[T any] struct {
	prev *doublyLinkedNode[T]
	next *doublyLinkedNode[T]
	data T
}

type doublyLinkedList[T any] struct {
	head *doublyLinkedNode[T]
	size int
}

func (l *doublyLinkedList[T]) show() {
	cur := l.head.next
	i :=1
	for cur != nil {
		fmt.Print(cur.data,",")
		i++
		cur =cur.next
	}
     fmt.Println() //结束时换行

}
func (l *doublyLinkedList[T]) add(data T) {
    newNode := &doublyLinkedNode[T]{data: data}

    if l.head.next == nil {
        // 链表为空，直接插入第一个节点
        l.head.next = newNode
        newNode.prev = l.head
    } else {
        // 找到最后一个节点
        cur := l.head.next
        for cur.next != nil {
            cur = cur.next
        }
        cur.next = newNode
        newNode.prev = cur
    }

    l.size++
}
func (l *doublyLinkedList[T]) insert(index int, data T) bool {
    if index <= 0 || index > l.size+1 {
        fmt.Println("索引值必须大于0且不大于链表长度+1")
        return false
    }

    // 找到插入位置的前一个节点
    prev := l.head
    for i := 1; i < index; i++ {
        prev = prev.next
    }

    // 创建新节点
    newNode := &doublyLinkedNode[T]{data: data, prev: prev, next: prev.next}

    // 处理 next 节点的 prev 指针（如果不是尾部插入）
    if prev.next != nil {
        prev.next.prev = newNode
    }

    // 连接 prev 和 newNode
    prev.next = newNode

    l.size++
    return true
}

func (l *doublyLinkedList[T]) BatchInsert(index int, datas []T) bool {
    if index <= 0 || index > l.size+1 {
        fmt.Println("索引值必须大于0且不大于链表长度+1")
        return false
    }

    if len(datas) == 0 {
        return true // 插入空集合，相当于不做事
    }

    // 遍历找到前一个节点
    prev := l.head
    for i := 1; i < index; i++ {
        prev = prev.next
    }

    // 保存原来的后续节点
    next := prev.next

    // 批量插入
    for _, d := range datas {
        node := &doublyLinkedNode[T]{data: d, prev: prev}
        prev.next = node
        prev = node
        l.size++
    }

    // 将原来的后续节点接回
    prev.next = next
    if next != nil {
        next.prev = prev
    }

    return true
}
func (l *doublyLinkedList[T]) delete(index int) bool {
    if index <= 0 || index > l.size {
        fmt.Println("索引值必须大于0且不大于链表长度")
        return false
    }

    // 找到前一个节点
    prev := l.head
    for i := 1; i < index; i++ {
        prev = prev.next
    }

    // 删除目标节点
    target := prev.next
    prev.next = target.next

    if target.next != nil {
        target.next.prev = prev
    }

    // 断开目标节点与链表的连接
    target.next = nil
    target.prev = nil

    l.size--
    return true
}
 func (l *doublyLinkedList[T]) find(index int)(node *doublyLinkedNode[T],ok bool){
	if(index <=0|| index>l.size){
		  fmt.Println("索引值必须大于0且不大于链表长度")
        return nil,false
	}
	target := l.head.next
	// 遍历找到目标节点
	for i:=1;i<index;i++{
		target=target.next
	}
	
	return target,true
} 
func DoublyLinkedListTest(){
	  // 创建带头结点的空链表
    list := &doublyLinkedList[int]{head: &doublyLinkedNode[int]{}}

    fmt.Println("=== 测试 add ===")
    list.add(10)
    list.add(20)
    list.add(30)
    list.show() // 10,20,30,

    fmt.Println("=== 测试 insert ===")
    list.insert(2, 15)
    list.show() // 10,15,20,30,

    fmt.Println("=== 测试 BatchInsert ===")
    list.BatchInsert(3, []int{100, 200, 300})
    list.show() // 10,15,100,200,300,20,30,

    fmt.Println("=== 测试 find ===")
    if node, ok := list.find(4); ok {
        fmt.Println("第4个节点数据:", node.data) // 应该是 200
    }

    fmt.Println("=== 测试 delete ===")
    list.delete(3) // 删除 100
    list.show() // 10,15,200,300,20,30,

    fmt.Println("=== 最终链表长度 ===")
    fmt.Println("size:", list.size) // 应该和元素数对应
}