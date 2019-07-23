package main

import "fmt"

/*  类型
	Go语言内置以下这些基础类型：
	 布尔类型： bool。
	 整型： int8、 byte、 int16、 int、 uint、 uintptr等。
	 浮点类型： float32、 float64。
	 复数类型： complex64、 complex128。
	 字符串： string。
	 字符类型： rune。
	 错误类型： error。
	此外， Go语言也支持以下这些复合类型：
	 指针（ pointer）
	 数组（ array）
	 切片（ slice）
	 字典（ map）
	 通道（ chan）
	 结构体（ struct）
	 接口（ interface）
 */

/*  布尔类型
	var v1 bool
	v1 = true
	v2 := (1 == 2) // v2也会被推导为bool类型
	布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换。
 */

/*  整型
	int和int32在Go语言里被认为是两种不同的类型
	Go语言支持下面的常规整数运算： +、 、 *、 /和%。
	Go语言支持以下的几种比较运算符： >、 <、 ==、 >=、 <=和!=。

	两个不同类型的整型数不能直接比较，比如int8类型的数和int类型的数不能直接比较，但
	各种类型的整型变量都可以直接与字面常量（ literal）进行比较，比如：
	var i int32
	var j int64
	i, j = 1, 2
	if i == j { 			// 编译错误
		fmt.Println("i and j are equal.")
	}
	if i == 1 || j == 2 {	 // 编译通过
		fmt.Println("i and j are equal.")
	}

	位运算符
	x<<y左移y位  x>>y右移y位  x^y异或  x&y与  x|y或  ^x取反
 */

/*  浮点型
	Go语言定义了两个类型float32和float64，其中float32等价于C语言的float类型，float64等价于C语言的double类型。
	var fvalue1 float32
	fvalue1 = 12
	fvalue2 := 12.0 // 如果不加小数点， fvalue2会被推导为整型而不是浮点型；推导出的浮点型是float64!!!!!
	fvalue1 = float32(fvalue2)	//必须强制转换类型才能进行赋值

	浮点数比较
	因为浮点数不是一种精确的表达方式，所以像整型那样直接用==来判断两个浮点数是否相等
	是不可行的，这可能会导致不稳定的结果。
	下面是一种推荐的替代方案：
	import "math"
	// p为用户自定义的比较精度，比如0.00001
	func IsEqual(f1, f2, p float64) bool {
	return math.Fdim(f1, f2) < p
	}
 */

/*  复数
	复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（ real），一个表示虚部（ imag）
	var value1 complex64 // 由2个float32构成的复数类型
	value1 = 3.2 + 12i
	value2 := 3.2 + 12i // value2是complex128类型
	value3 := complex(3.2, 12) // value3结果同 value2

	对于一个复数z = complex(x, y)，就可以通过Go语言内置函数real(z)获得该复数的实部，也就是x，通过imag(z)获得该复数的虚部，也就是y。
	更多关于复数的函数，请查阅math/cmplx标准库的文档
 */

/*  字符串
	在Go语言中，字符串也是一种基本类型。相比之下， C/C++语言中并不存在原生的字符串类型，通常使用字符数组来表示，并以字符指针来传递。
	Go语言中字符串的声明和初始化非常简单，举例如下：
	var str string // 声明一个字符串变量
	str = "Hello world" // 字符串赋值
	ch := str[0] // 取字符串的第一个字符
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
	输出结果为：
	The length of "Hello world" is 11
	The first character of "Hello world" is H.
	字符串的内容可以用类似于数组下标的方式获取，但与数组不同，字符串的内容不能在初始化后被修改，比如以下的例子：
	str := "Hello world" // 字符串也支持声明时进行初始化的做法
	str[0] = 'X' // 编译错误

	字符串操作： x+y连接   len(s)长度  s[i]取字符			更多的字符串操作，请参考标准库strings包

	字符串遍历：1、以字节数组byteArray方式遍历,编码为UTF-8；2、以Unicode字符遍历
	1、	str := "Hello,世界"
		n := len(str)
		for i := 0; i < n; i++ {
			ch := str[i] // 依据下标取字符串中的字符，类型为byte
			fmt.Println(i, ch)
		}
	2、 str := "Hello,世界"
		for i, ch := range str {
			fmt.Println(i, ch)//ch的类型为rune
		}
	以Unicode字符方式遍历时，每个字符的类型是rune（早期的Go语言用int类型表示Unicode字符），而不是byte
 */

/*  字符类型
	在Go语言中支持两个字符类型，一个是byte（实际上是uint8的别名），代表UTF-8字符串的单个字节的值；
	另一个是rune，代表单个Unicode字符。
	关于rune相关的操作，可查阅Go标准库的unicode包。另外unicode/utf8包也提供了UTF8和Unicode之间的转换。
	出于简化语言的考虑， Go语言的多数API都假设字符串为UTF-8编码。尽管Unicode字符在标准库中有支持，但实际上较少使用
 */

/*  数组
	以下为一些常规的数组声明方法：
	[32]byte // 长度为32的数组，每个元素为一个字节
	[2*N] struct { x, y int32 } // 复杂类型数组
	[1000]*float64 // 指针数组
	[3][5]int // 二维数组
	[2][2][2]float64 // 等同于[2]([2]([2]float64))

	在Go语言中，数组长度在定义后就不可更改，在声明时长度可以为一个常量或者一个常量
	表达式（常量表达式是指在编译期即可计算结果的表达式）。数组的长度是该数组类型的一个内
	置常量，可以用Go语言的内置函数len()来获取

	元素访问
	1、可以使用数组下标来访问数组中的元素
	2、Go语言还提供了一个关键字range，用于便捷地遍历容器中的元素。当然，数组也是range的支持范围。上面的遍历过程可以简化为如下的写法：
	for i, v := range array {
		fmt.Println("Array element[", i, "]=", v)
	}

	值类型
	需要特别注意的是，在Go语言中数组是一个值类型（ value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。
	如果将数组作为函数的参数类型，则在函数调用时该参数将发生数据复制。
	因此，在函数体中无法修改传入的数组的内容，因为函数内操作的只是所传入数组的一个副本
	下面用例子来说明这一特点：
	package main
	import "fmt"
	func modify(array [10]int) {
		array[0] = 10 // 试图修改数组的第一个元素
		fmt.Println("In modify(), array values:", array)
	}
	func main() {
		array := [5]int{1,2,3,4,5} // 定义并初始化一个数组
		modify(array) // 传递给一个函数，并试图在函数体内修改这个数组内容
		fmt.Println("In main(), array values:", array)
	}
	该程序的执行结果为
	In modify(), array values: [10 2 3 4 5]
	In main(), array values: [1 2 3 4 5]
	从执行结果可以看出， 函数modify()内操作的那个数组跟main()中传入的数组是两个不同的实例。
	那么，如何才能在函数内操作外部的数据结构呢？我们将用数组切片功能来达成这个目标。
 */

/*  数组切片

 */

/*  map(字典)

 */

func main() {
	var myStr = "hello, world"
	fmt.Println(myStr)

	var myArray = [5]int{3,10,5,4,8}
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	for _,v := range myArray {
		fmt.Println(v," ")
	}

	var z = complex(3, 5)
	fmt.Println(z)
	fmt.Println(real(z),"+",imag(z),"i；")
}
