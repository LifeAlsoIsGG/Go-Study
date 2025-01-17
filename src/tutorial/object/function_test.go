package object

import (
	"fmt"
	"testing"
	"time"
)

// refer:https://www.liwenzhou.com/posts/Go/function/

// panic & recover
func set_data(x int) {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var arr [10]int
	arr[x] = 1
}
func TestPanic(t *testing.T) {
	set_data(11)

	// 如果能执行到这句，说明panic被捕获了
	// 后续的程序能继续运行
	fmt.Println("everything is ok")
}

// panic&recover 无法跨协程
func TestPanicGroutine(t *testing.T) {
	// 这个 defer 并不会执行
	defer fmt.Println("in main")

	go func() {
		defer println("in goroutine")
		panic("")
	}()

	time.Sleep(2 * time.Second)
}

// 多个 defer调用，按照栈先进后出
// java
// python
// go
func TestMultiDefer(t *testing.T) {
	name := "go"
	defer fmt.Println(name) // 输出: go

	name = "python"
	defer fmt.Println(name) // 输出: python

	name = "java"
	fmt.Println(name)
}

// defer 与 return 顺序
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

// 5 6 5 5
func TestDefer_1(t *testing.T) {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

// 其它 defer 调用例子
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4
func TestDefer_2(t *testing.T) {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
