package main

import (
	"fmt"
	"testing"
)

func create() {
	// 第一种方法
	var _ [3]int = [3]int{1, 2, 3}

	// 第二种方法
	_ = [3]int{1, 2, 3}

	//第三种方式，不指定大小
	_ = [...]int{1, 2, 3}
}

func create2() {

	/*如果你觉得每次写 [3]int 有点麻烦，你可以为 [3]int 定义一个类型字面量，也就是别名类型。
	使用 type 关键字可以定义一个类型字面量，后面只要你想定义一个容器大小为3，元素类型为int的数组 ，都可以使用这个别名类型。
	*/
	type arr3 [3]int

	myarr := arr3{2: 3}
	fmt.Printf("%d 的类型是: %T\n", myarr, myarr)
	fmt.Println(myarr[0:3])
}

// 切片构造方式
func TestSliceCreation(t *testing.T) {

	//切片表达式：从字符串、数组、指向数组或切片的指针（即从切片再构造）中构造子字符串或切片
	myarr := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("myarr:%v len(s):%v cap(s):%v\n", myarr, len(myarr), cap(myarr))
	// 指定 low 和 high，不指定 cap 时切片终止索引会一直到原数组的最后一个数，也就是 len - low
	mysli1 := myarr[1:3]
	fmt.Printf("mysli1:%v len(s):%v cap(s):%v\n", mysli1, len(mysli1), cap(mysli1)) // mysli1:[2 3] len(s):2 cap(s):4
	// 指定 low、high 和 cap
	mysli2 := myarr[1:3:5]
	fmt.Printf("mysli2 的长度为：%d，容量为：%d\n", len(mysli2), cap(mysli2))
	fmt.Println(mysli2)

	// 从切片再构造切片，或字符串构造切片
	fmt.Println(mysli2[:], "xxx"[:])

	//第二种方式：声明 var name []T
	var _ []string
	var _ []int
	var _ = []int{}
	var _ = []int{1, 2}

	//第三种方式：使用 make
	//make 函数的格式：make( []Type, size, cap )，same as make([]T, cap)[:len]，没指定cap则默认为len值
	mysli3 := make([]int, 2)
	mysli4 := make([]int, 2, 10)
	fmt.Println(mysli3, mysli4)
}

func TestSliceCompare(t *testing.T) {
	// 切片比较，只能和 nil 比较
	var sli1 []int
	fmt.Println(sli1 == nil) // true

	// 判断空切片应该使用 len(s) == 0 来判断，一个nil值的切片的长度和容量都是0。
	// 如果用空 slice 与 nil比较，结果为 false，会误以为这个 slice
	var (
		sli2 []int
		sli3 = []int{}
	)
	println(sli2 == nil, sli3 == nil) // true false
}

func TestSliceMethod(t *testing.T) {
	myarr := []int{100}
	fmt.Println(myarr)
	// 追加一个元素
	myarr = append(myarr, 2)
	// 追加多个元素
	myarr = append(myarr, 3, 4)
	// 追加一个切片, ... 表示解包，不能省略
	myarr = append(myarr, []int{7, 8}...)
	// 在第一个位置插入元素
	myarr = append([]int{0}, myarr...)
	// 在中间插入一个切片(两个元素)
	myarr = append(myarr[:5], append([]int{5, 6}, myarr[5:]...)...)
	fmt.Println(myarr)
}

func TestSlice_2(t *testing.T) {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Println(myslice)
	fmt.Println(cap(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Println(myslice)
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}

// append 在不扩容情况下是如何修改底层数组
func TestSliceAppend(t *testing.T) {
	array := [...]int8{1, 2, 3, 4, 5}
	sli1 := array[:2]
	sli2 := array[1:]
	fmt.Println(sli1) // [1 2]
	fmt.Println(sli2) // [2 3 4 5]
	sli1 = append(sli1, 6)
	fmt.Println(sli1) // [1 2 6]
	fmt.Println(sli2) // [2 6 4 5]
}

func TestSliceGrowing118(t *testing.T) {
	var s []int
	for i := 0; i < 4098; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
