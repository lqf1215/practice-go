package base

import (
	"fmt"
	"testing"
)

/**
genericity 泛型
减少重复代码的编写
比如说 在比较两个数大小的时候，没有泛型时 只是传入类型不一样，我们就要写一份一模一样的函数（只是传入类型不一样） 如果有了泛型就可以减少重复代码了。
*/

func TestGenericity(t *testing.T) {
	a := int(1)
	b := int(2)
	num := GetMaxNum1(a, b)
	fmt.Println(num)

	//num1:=GetMaxNum2(a,b)
	//fmt.Println(num1)

}

//使用泛型
func GetMaxNum1[T int | int8](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 自定义泛型类型
// 如类型太多了 这时候我们可以使用自定义泛型类型

//想声明接口一样声明自定义泛型类型
type GetMaxNumType interface {
	int | int8 | int16 | int32 | int64
}

//func GetMaxNum2(T  GetMaxNumType) (a, b T) T  {
//	if a>b{
//		return a
//	}
//	return b
//}
