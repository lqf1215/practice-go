package algorithm

import (
	"fmt"
	"testing"
)

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func TestLinkNode(t *testing.T) {
	node := new(LinkNode)
	node.Data = 2
	// 新的节点
	node1 := new(LinkNode)
	node1.Data = 3
	node.NextNode = node1 // node1 链接到 node 节点上

	// 新的节点
	node2 := new(LinkNode)
	node2.Data = 4
	node1.NextNode = node2 // node2 链接到 node1 节点上

	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			// 打印节点值
			fmt.Println(nowNode.Data)
			// 获取下一个节点
			nowNode = nowNode.NextNode
			continue
		}

		// 如果下一个节点为空，表示链表结束了
		break
	}
}

// 循环链表
type Ring struct {
	next, prev *Ring       // 前和后 节点
	Value      interface{} // 数据

}

// 初始化空的循环链表，前驱和后驱都指向自己，因为是循环的
func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

func TestRing(t *testing.T) {
	r := new(Ring)
	r.init()

	slice := make([]int, 0)
	slice = append(slice, 1, 2)
	slice = append(slice, 2)
	fmt.Println(slice, len(slice), cap(slice))
}

// 创建N个节点的唤循环链表 会连续绑定前驱和后驱节点，时间复杂度为：O(n)。
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

// 获取下一个节点 获取前驱或后驱节点，时间复杂度为：O(1)。
func (r *Ring) Next() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.next
}

// 获取上一个节点
func (r *Ring) Prev() *Ring {
	if r.next == nil {
		return r.init()
	}
	return r.prev
}

//因为链表是循环的，当 n 为负数，表示从前面往前遍历，否则往后面遍历：因为需要遍历 n 次，所以时间复杂度为：O(n)。
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 往节点A，链接一个节点，并且返回之前节点A的后驱节点 添加节点的操作比较复杂，如果节点 s 是一个新的节点。
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev()
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

func linkNewTest() {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

func TestLinkNew(t *testing.T) {
	linkNewTest()
	// 创建一个容量为2的切片
	array := make([]int, 0, 2)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)

	// 虽然 append 但是没有赋予原来的变量 array
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	_ = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)

	fmt.Println("-------")

	// 赋予回原来的变量
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1, 1, 1, 1)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2)
	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
	//array = append(array, 1)
	//fmt.Println("cap", cap(array), "len", len(array), "array:", array)
}
