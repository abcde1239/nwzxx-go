package dataStructure

import "fmt"

type linkedNode[T any] struct {
	next *linkedNode[T]
	data T
}

type singlyLinkedList[T any] struct {
	head *linkedNode[T]
	size int
}

func (l *singlyLinkedList[T]) show() {
	cur := l.head.next
	for cur != nil {
		fmt.Print(cur.data,",")
		cur =cur.next
	}
     fmt.Println() //结束时换行

}
func (l *singlyLinkedList[T])add(data T){
    cur:=l.head.next
    for cur.next !=nil{
        cur =cur.next

    }
    newNode :=&linkedNode[T]{data:data}
    cur.next =newNode
}
//插入节点
func (l *singlyLinkedList[T]) insert(index int, data T) (ok bool) {
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
func (l *singlyLinkedList[T]) BatchInsert(index int, datas []T) bool {
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
func (l *singlyLinkedList[T]) update(index int, data T) (ok bool) {
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
func (l *singlyLinkedList[T]) delete(index int) bool {
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
func (l *singlyLinkedList[T]) find(index int)(node *linkedNode[T],ok bool){
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
func (l *singlyLinkedList[T]) findData(index int)(data T,ok bool){
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
//此为迭代法反转
func (l *singlyLinkedList[T]) reverse() {
    var prev *linkedNode[T] = nil
    cur := l.head

    for cur != nil {
        next := cur.next   
        cur.next = prev    
        prev = cur         
        cur = next         
    }

    l.head = prev 
}
//此为递归法反转，仅供理解
/* func (l *singlyLinkedList[T]) reverseRecursive() {
    l.head = reverseRecursiveHelper(l.head)
}

// 辅助函数：递归反转链表，从 node 开始
func reverseRecursiveHelper[T any](node *linkedNode[T]) *linkedNode[T] {
    if node == nil || node.next == nil {
        return node 
    }

    newHead := reverseRecursiveHelper(node.next) 
    node.next.next = node 
    node.next = nil      
    return newHead
} */

func SinglyLinkedListTest() {
    l := &singlyLinkedList[int]{head: &linkedNode[int]{}}

    fmt.Println("=== 1. 批量插入 [10,20,30] 到位置 1 ===")
    fmt.Println("结果:", l.BatchInsert(1, []int{10, 20, 30}))
    l.show() // 10,20,30

    fmt.Println("=== 2. 在位置 2 插入 3 ===")
    fmt.Println("结果:", l.insert(2, 3))
    l.show() // 10,3,20,30

    fmt.Println("=== 3. 尝试非法插入到位置 0 ===")
    fmt.Println("结果:", l.insert(0, 114514)) // 应该失败
    l.show()

    fmt.Println("=== 4. 删除位置 2 的节点 ===")
    fmt.Println("结果:", l.delete(2))
    l.show()

    fmt.Println("=== 5. 查找第 1 个节点和数据 ===")
    if temp, ok := l.find(1); ok {
        data, _ := l.findData(1)
        fmt.Println("find:", temp, "findData:", data)
    }

    fmt.Println("=== 6. 更新节点数据 ===")
    fmt.Println("非法更新(位置0):", l.update(0, 34)) // 应该失败
    l.show()
    fmt.Println("合法更新(位置1):", l.update(1, 34))
    l.show()

    fmt.Println("=== 7. 反转链表 ===")
    l.reverse()
    l.show()

    fmt.Println("=== 8. 测试尾插 add(666) ===")
    l.add(666)
    l.show()

    fmt.Println("=== 链表长度 ===")
    fmt.Println(l.size)
}
