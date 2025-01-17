package src

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func Add(a int, b int) int {
	return a + b
}

func Mul(a int, b int) int {
	return a * b
}

type Weekday uint

const (
	Sunday Weekday = 2 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// 子测试
func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

func Test2(t *testing.T) {
	var a interface{} = 1
	v, ok := a.(string)
	t.Log(v, ok)
}

var cwd string = "xxx"

func Test3(t *testing.T) {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}

func Test4(t *testing.T) {
	var myfloat float32 = 10000018
	fmt.Println("myfloat: ", myfloat)
	fmt.Println(myfloat + 100)

	var myfloat1 float32 = 5.2
	fmt.Println("myfloat1: ", myfloat1+1)
	fmt.Println(myfloat1 + 2)

	s := `hello, world`
	fmt.Println(len(s))             // "12"
	fmt.Printf("%c %c", s[0], s[2]) // "104 119" ('h' and 'w')

	strconv.FormatInt(int64(10), 2)
	symbol := [...]string{0: "$", 1: "€", 2: "￡", 3: "￥"}
	fmt.Println(symbol)

	a := []int{}
	fmt.Println(a == nil)

}

func Test5(t *testing.T) {
	var a int8 = -60
	var b int8 = -128 + a

	println(b)

}

func Test6(t *testing.T) {
	a := 0
	a1 := &a
	b := 1
	b1 := &b
	*a1, *b1 = *b1, *a1
	fmt.Println(*a1, *b1)
}

type a1 struct {
	c int
}

type b struct {
	a1
	a int
	b int
}

func (a *a1) test1() {
	fmt.Println()
}

//func TestName(t *testing.T) {
//	a := &b{}
//	a.test1()
//
//	sync.Map{}
//}

func TestTrim(t *testing.T) {
	s := "  aa  sdva dgd "
	bytes := []byte(s)
	i := 0
	j := len(bytes) - 1
	space_count := 0
	for i <= j {
		if bytes[i] == ' ' {
			bytes[i] = bytes[j]
			bytes[j] = '\x00'
			j--
			space_count++
			continue
		}
		i++
	}
	fmt.Println(string(bytes))
	fmt.Println(space_count)
}

func TestTrim2(t *testing.T) {
	s := "  aa  sdva dgd "
	bytes := []byte(s)
	i := 0
	j := 0
	lens := len(bytes)
	count := 0
	for j < lens {
		// 如果不是空格
		if bytes[j] != ' ' {
			bytes[i] = bytes[j]
			i++
		} else {
			count++
		}
		j++
	}
	fmt.Println(string(bytes[:lens-count]))
	fmt.Println(count)
}

func findSumPair(nums []int, target int) [][]int {
	result := [][]int{}
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[target-num]; ok {
			result = append(result, []int{num, target - num})
		}
		m[num]++
	}
	return result
}

func TestFindSumPair(t *testing.T) {
	arr := []int{1, 4, 2, 5, 9, 2, 3}
	fmt.Println(findSumPair(arr, 6))
}

// 定义一个函数，接受一个整数数组和一个目标值作为参数，数组中可能存在重复元素
func findPairs(arr []int, target int) [][]int {
	result := [][]int{}
	// 定义一个哈希表，用来存储数组中的元素和它们的索引
	hash := make(map[int]int)
	for i, num := range arr {
		// 计算目标值减去当前元素的差值
		diff := target - num
		// 检查哈希表中是否存在这个差值
		if _, ok := hash[diff]; ok {
			// 如果存在，就将这两个元素组成一对，并添加到结果中
			result = append(result, []int{num, diff})
			// 从哈希表中删除这两个元素，以避免重复
			delete(hash, num)
			delete(hash, diff)
		} else {
			// 如果不存在，就将当前元素和它的索引添加到哈希表中
			hash[num] = i
		}
	}
	// 返回结果
	return result
}

func TestFindPairs(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 2, 4}
	target := 6
	fmt.Println(findPairs(arr, target))
}

// 定义一个函数，接受一个整数数组作为参数
func maxCoins(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return arr[0]
	}
	dp := make([]int, n)
	dp[0] = arr[0]
	dp[1] = max(arr[0], arr[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxCoins(t *testing.T) {
	arr1 := []int{10, 11, 12}
	arr2 := []int{10, 100, 91, 1, 1}
	fmt.Println(maxCoins(arr1))
	fmt.Println(maxCoins(arr2))
}

func TestURLQueryEscape(t *testing.T) {
	q := url.QueryEscape(strings.Join([]string{"a", "&", "c"}, " "))
	println(q)
}
