package dataStructure

import "fmt"

type binaryTreeNode[T comparable] struct {
	data      T
	leftNode  *binaryTreeNode[T]
	rightNode *binaryTreeNode[T]
}
type binaryTree[T comparable] struct {
	root *binaryTreeNode[T]
	size int
}
//使用递归，真实场景仍应使用迭代实现DFS
func (t *binaryTree[T]) preOrder(root *binaryTreeNode[T]) {
	if root == nil {
		return
	}
	fmt.Print(root.data,",")
	t.preOrder(root.leftNode)
	t.preOrder(root.rightNode)
}
func (t *binaryTree[T]) inOrder(root *binaryTreeNode[T]) {
		if root == nil {
		return
	}
	t.preOrder(root.leftNode)
	fmt.Print(root.data,",")
	t.preOrder(root.rightNode)
}

func (t *binaryTree[T]) postOrder(root *binaryTreeNode[T]) {
		if root == nil {
		return
	}
	t.preOrder(root.leftNode)
	t.preOrder(root.rightNode)
	fmt.Print(root.data,",")
}
//BFS
func (t *binaryTree[T]) levelOrder(root *binaryTreeNode[T]) {
		if(root ==nil){
			return
		}
		queue :=[]*binaryTreeNode[T]{root}
		for len(queue)>0{
			node :=queue[0]
			queue =queue[1:]
			fmt.Print(node.data,",")
			if(node.leftNode !=nil){
				queue=append(queue, node.leftNode)
			}
			if(node.rightNode !=nil){
				queue =append(queue, node.rightNode)
			}

		}
}
//此为通过栈迭代实现的DFS,在空间复杂度上优于递归
func (t *binaryTreeNode[T]) preOrderByStack(root *binaryTreeNode[T]){
		if(root == nil){
			return
		}
		stack := []*binaryTreeNode[T]{root}
		for len(stack)>0{
			node :=stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			fmt.Print(node.data,",")
			if(node.rightNode!=nil){
				stack =append(stack, node.rightNode)
			}
			if(node.leftNode !=nil){
				stack=append(stack, node.leftNode)
			}

		}


	}
func (t *binaryTree[T]) insert(data T)(root *binaryTreeNode[T]) {
		queue :=[]*binaryTreeNode[T]{t.root}
		newNode :=&binaryTreeNode[T]{data:data}
		for len(queue)>0{
			node :=queue[0]
			queue =queue[1:]
			
			if(node.leftNode ==nil){
				node.leftNode = newNode
				return t.root
			}else{
				queue =append(queue, node.leftNode)
			}
			if(node.rightNode ==nil){
				node.rightNode =newNode
				return t.root
			}else{
				queue =append(queue, node.rightNode)
			}

		}
		return t.root
}
//此处为依据数据删除，并利用尾节点补充
func(t*binaryTree[T]) deleteNode( data T) *binaryTreeNode[T] {
	if t.root == nil {
		return nil
	}
	var target *binaryTreeNode[T]
	var parent*binaryTreeNode[T]
	queue := []*binaryTreeNode[T]{t.root}
	var last *binaryTreeNode[T]
	for len(queue) > 0 {
		last = queue[0]
		queue = queue[1:]
		if last.data == data {
			target = last
		}
		if last.leftNode != nil {
			parent = last
			queue = append(queue, last.leftNode)
		}
		if last.rightNode != nil {
			parent = last
			queue = append(queue, last.rightNode)
		}
	}
	// 替换目标节点
	if target != nil && last != nil {
		target.data = last.data
		// 删除最后一个节点
		if parent.rightNode == last {
			parent.rightNode = nil
		} else {
			parent.leftNode = nil
		}
	}
	return t.root
}
func (t *binaryTree[T]) getHeight(root *binaryTreeNode[T]) int {
	if root == nil {
		return 0
	}

	h := 0
	queue := []*binaryTreeNode[T]{root}

	for len(queue) > 0 {
		levelSize := len(queue) // 当前层的节点数量
		h++           

		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			if node.leftNode != nil {
				queue = append(queue, node.leftNode)
			}
			if node.rightNode != nil {
				queue = append(queue, node.rightNode)
			}
		}
	}
	return h
}

// 普通树实在没什么可写的了，下一个更新二叉搜索会舒服一点

func BinaryTreeTest() {
	// 初始化二叉树
	tree := &binaryTree[int]{root: &binaryTreeNode[int]{data: 1}}

	// 层序插入
	fmt.Println("=== 1. 批量插入 [2,3,4,5,6] ===")
	for _, v := range []int{2, 3, 4, 5, 6} {
		tree.insert(v)
	}
	fmt.Print("层序遍历: ")
	tree.levelOrder(tree.root)
	fmt.Println()

	// 前序遍历
	fmt.Print("前序遍历: ")
	tree.preOrder(tree.root)
	fmt.Println()

	// 中序遍历
	fmt.Print("中序遍历: ")
	tree.inOrder(tree.root)
	fmt.Println()

	// 后序遍历
	fmt.Print("后序遍历: ")
	tree.postOrder(tree.root)
	fmt.Println()

	// 获取高度
	fmt.Println("树高度:", tree.getHeight(tree.root))

	// 删除节点
	fmt.Println("=== 2. 删除节点 3 ===")
	tree.deleteNode(3)
	fmt.Print("层序遍历删除后: ")
	tree.levelOrder(tree.root)
	fmt.Println()
	fmt.Println("树高度删除后:", tree.getHeight(tree.root))

	// 再删除根节点
	fmt.Println("=== 3. 删除根节点 1 ===")
	tree.deleteNode(1)
	fmt.Print("层序遍历删除根后: ")
	tree.levelOrder(tree.root)
	fmt.Println()
	fmt.Println("树高度删除根后:", tree.getHeight(tree.root))
	 println()
}
