package base

import (
	"testing"
)

/**
 堆栈
计算机种的堆栈包含有两部分：数据结构的堆栈和内存分配中堆栈
内存分配中的堆和栈
堆：一般由程序员分配释放 栈：由操作系统自动分配释放
堆栈的缓存方式：栈使用的是一级缓存，通常都是被调用时处于存储空间中，调用完毕立刻释放。
堆则是存放在二级缓存中，生命周期由虚拟机的垃圾回收算法决定。

*/

func TestName(t *testing.T) {
	m := foo(1)
	println(*m)
}

func foo(m0 int) *int {
	var m1 int = 11
	var m2 int = 12
	var m3 int = 13
	var m4 int = 14
	var m5 int = 15

	println(&m0, &m1, &m2, &m3, &m4, &m5)
	return &m3
}

func Test1(t *testing.T) {

}
