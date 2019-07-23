package main

import "os" // 用于获得命令行参数os.Args
import "fmt"
import "mycalculator/src/simplemath"
import "strconv"

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n" +
		"\tadd\tAddition of A and B.\n" +
		"\tsub\tSubtract A from B.\n" +
		"\tmulti\tMultiply A by B.\n" +
		"\tdiv\tDivide A by B.\n" +
		"\tsqrt\tSquare root of a non-negative value A.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	fmt.Println("Hello, mycalc!")
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: mycalc add <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])	// Atoi 是将字符串转换为整型变量，是会有小数丢失的
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: mycalc add <integer1> <integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sub":
		if len(args) != 4 {
			fmt.Println("USAGE: mycalc sub <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: mycalc sub <integer1> <integer2>")
		}
		ret := simplemath.Sub(v1, v2)
		fmt.Println("Result: ", ret)
	case "multi":
		if len(args) != 4 {
			fmt.Println("USAGE: mycalc multi <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: mycalc multi <integer1> <integer2>")
		}
		ret := simplemath.Multi(v1, v2)
		fmt.Println("Result: ", ret)
	case "div":
		if len(args) != 4 {
			fmt.Println("USAGE: mycalc div <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: mycalc div <integer1> <integer2>")
		}
		ret := simplemath.Div(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: mycalc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: mycalc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
