package algorithm

import "sync"

// 数组栈 后进先出
type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

type LinkStack struct {
	root *LinkStackNode
	size int
	lock sync.Mutex
}

type LinkStackNode struct {
	Value string
	Next  *LinkStackNode
}

func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.root == nil {
		stack.root = new(LinkStackNode)
		stack.root.Value = v
	} else {
		preNode := stack.root
		newNode := new(LinkStackNode)
		newNode.Value = v

		newNode.Next = preNode
		// 将新节点放在头部
		stack.root = newNode
	}
	stack.size = stack.size + 1

}
