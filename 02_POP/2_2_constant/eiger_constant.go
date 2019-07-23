package main

import "fmt"

/*  常量定义
	const Pi float64 = 3.14159265358979323846
	const zero = 0.0 // 无类型浮点常量
	const (
		size int64 = 1024
		eof = -1 // 无类型整型常量
	)
	const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
	const a, b, c = 3, 4, "foo"
	// a = 3, b = 4, c = "foo", 无类型整型和字符串常量

	const mask = 1 << 3		// 常量定义的右值也可以是一个在编译期运算的常量表达式
	// 由于常量的赋值是一个编译期行为，所以右值不能出现任何需要运行期才能得出结果的表达式

*/

/*  预定义常量
	true  false  iota
	iota比较特殊，可以被认为是一个可被编译器修改的常量，在每一个const关键字出现时被
	重置为0，然后在下一个const出现之前，每出现一次iota，其所代表的数字会自动增1

	const ( // iota被重设为0
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)

	如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。因此，上
	面的前两个const语句可简写为：
	const ( 		// iota被重设为0
		c0 = iota 	// c0 == 0
		c1 			// c1 == 1
		c2 			// c2 == 2
	)
*/

/*  GO中的枚举
	枚举指一系列相关的常量
	const (
		Sunday = iota		//注意开头字母大写表示导出为外部可见
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays 		// 这个常量没有导出
	)
 */


func main() {
	// 初始化
	const (
		A = iota
		B
		C
		numofCharacters
	)

	fmt.Println(B)
	fmt.Println(numofCharacters)

	const notice int = 3*iota
	println(notice)
}