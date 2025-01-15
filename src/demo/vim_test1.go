package main

import "fmt"

func main() {
	fmt.Println("abc abc bac")

	s := "你好啊"
	fmt.Println(len(s)) // "12"
	fmt.Println(s[:5])  // "104 119" ('h' and 'w')

	fmt.Println("响铃：\a")
	fmt.Println("退格：hello\bworld")
	fmt.Println("换页：\f")
	fmt.Println("换行：\n新的一行")
	fmt.Println("回车：\r覆盖前面的内容")
	fmt.Println("制表符：hello\tworld")
	fmt.Println("垂直制表符：\v")
	fmt.Println("单引号：'\\''")
	fmt.Println("双引号：\"\\\"\"")
	fmt.Println("反斜杠：\\\\")

}
