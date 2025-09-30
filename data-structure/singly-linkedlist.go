package main

import "fmt"

type linkedNode[T any] struct {
	next *linkedNode[T]
	data T
}

type singlyLinkedlist[T any] struct {
	head *linkedNode[T]
	size int
}

func (l *singlyLinkedlist[T]) show() {
	cur := l.head.next
	i :=1
	for cur != nil {
		fmt.Println("第",i,"个节点值为",cur.data)
		i++
		cur =cur.next
	}

}
//插入节点
func (l *singlyLinkedlist[T]) insert(index int, data T) (ok bool) {
    if index <= 0 || index > l.size+1 {
        fmt.Println("索引值必须大于0且不大于链表长度+1")
        return false
    }

    // 遍历找到前一个节点
    prev := l.head
    for i := 1; i < index; i++ {
        prev = prev.next
    }

    // 创建新节点并插入
    newNode := &linkedNode[T]{data: data, next: prev.next}
    prev.next = newNode
    l.size++
    return true
}
//批量插入
func (l *singlyLinkedlist[T]) BatchInsert(index int, datas []T) bool {
    if index <= 0 || index > l.size+1 {
        fmt.Println("索引值必须大于0且不大于链表长度+1")
        return false
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
        node := &linkedNode[T]{data: d}
        prev.next = node
        prev = node
        l.size++
    }

    // 将原来的后续节点接回
    prev.next = next

    return true
}

//更新节点
func (l *singlyLinkedlist[T]) update(index int, data T) (ok bool) {
	if(index <=0|| index>l.size){
		  fmt.Println("索引值必须大于0且不大于链表长度")
        return false
	}
	target := l.head.next
	// 遍历找到目标节点
	for i:=1;i<index;i++{
		target=target.next
	}
	target.data =data
	return true
}
//删除节点
func (l *singlyLinkedlist[T]) delete(index int) bool {
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
    target.next = nil //断开目标节点与链表的连接

    l.size--
    return true
}
// 此方法可获得节点指针用于操作
func (l *singlyLinkedlist[T]) find(index int)(node *linkedNode[T],ok bool){
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
// 此方法可获得节点数据
func (l *singlyLinkedlist[T]) findData(index int)(data T,ok bool){
	if(index <=0|| index>l.size){
		  fmt.Println("索引值必须大于0且不大于链表长度")
        return *new(T), false
	}
	target := l.head.next
	// 遍历找到目标节点
	for i:=1;i<index;i++{
		target=target.next
	}
	
	return target.data,true
} 
func main() {
	l :=&singlyLinkedlist[int]{head:&linkedNode[int]{}}
	fmt.Println(l.BatchInsert(1,[]int{10,20,30}))
	l.show()
	fmt.Println(l.BatchInsert(2,[]int{100,200}))
	l.show()
}