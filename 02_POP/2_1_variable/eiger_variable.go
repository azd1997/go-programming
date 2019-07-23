package main

import "fmt"

/*  变量声明
	var v1 int
	var v2 string
	var v3 [10]int // 数组
	var v4 []int // 数组切片
	var v5 struct {
		f int
	}
	var v6 *int // 指针
	var v7 map[string]int // map， key为string类型， value为int类型
	var v8 func(a int) int

	var (
		v1 int
		v2 string
	)
*/

/*  变量初始化
	var v1 int = 10 // 正确的使用方式1
	var v2 = 10 // 正确的使用方式2，编译器可以自动推导出v2的类型
	v3 := 10 // 正确的使用方式3，编译器可以自动推导出v3的类型
*/

/*  变量赋值
	var v10 int
	v10 = 123	// 普通赋值方式

	i, j = j, i	// 多重赋值方式，可以用来变量的交换
*/

/*  匿名变量
	func GetName() (firstName, lastName, nickName string) {
		return "May", "Chan", "Chibi Maruko"
	}

	_, _, nickName := GetName()
*/

func main(){
	fmt.Println("变量声明：")
	var a int = 9
	var b int = 1
	a, b = b, a
	println("a = ", a, ";", "b = ", b)

}
