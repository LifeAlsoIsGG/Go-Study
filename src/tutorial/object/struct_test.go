package object

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// 自定义类型和类型别名
type MyInt int // 自定义类型：是一个新的类型，具有 int 的特性

type AliasInt = int // 类型别名：本质上是同一个类型，例如 type byte = uint8

func TestType(t *testing.T) {
	var my_int MyInt
	var alias_int AliasInt

	fmt.Printf("type of my_int:%T\n", my_int)       //type of my_int:main.NewInt
	fmt.Printf("type of alias_int:%T\n", alias_int) //type of alias_int:int
}

// 定义结构体
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile // 指针
	father *Profile // 指针
}

type company struct {
	companyName string
	companyAddr string
}

// 将company的属性组合进来
type staff struct {
	company
	name     string
	age      int
	gender   string
	position string
}

func TestDeclare(t *testing.T) {
	//第一种
	xm := Profile{
		name:   "小明",
		age:    18,
		gender: "male",
	}
	xm = Profile{name: "小红"}
	fmt.Println(xm.name)

	//第二种
	xm_2 := new(Profile)
	// 等价于: var xm *Profile = new(Profile)
	fmt.Println(xm_2)
	// output: &{ 0 }

	//第三种
	var xm_3 *Profile = &Profile{}
	fmt.Println(xm_3)
	// output: &{ 0 }

	//选择器.能直接解引用
	xm_3.name = "iswbm"  // 或者 (*xm).name = "iswbm"
	xm_3.age = 18        //  或者 (*xm).age = 18
	xm_3.gender = "male" // 或者 (*xm).gender = "male"
	fmt.Println(xm_3)
	//output: &{iswbm 18 male}
}

//当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
//当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。

// 结构体定义方法：以值做为方法接收者
func (person Profile) fmtProfile() {
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("年龄：%d\n", person.age)
	fmt.Printf("性别：%s\n", person.gender)
}

// 结构体定义方法：修改实例要用指针接收，建议统一使用这个
func (person *Profile) changeProfile() {
	person.name = "小红"
	person.age = 20
	person.gender = "female"
}

// 不管接收器是指针或非指针，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。

func TestMethodAllocate(t *testing.T) {
	// 实例化
	myself := Profile{name: "小明", age: 24, gender: "male"}
	myself.changeProfile()
	myself.fmtProfile()
}

// 测试结构体继承
func TestExtend(t *testing.T) {
	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company:  myCom,
	}

	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName) // 等同于下面
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)
}

func TestPrintStructByFMT(t *testing.T) {
	xm := Profile{
		name:   "小明",
		age:    18,
		gender: "male",
	}

	//%v	值的默认格式表示
	fmt.Printf("%v\n", xm) // {小明 18 male <nil> <nil>}

	//%+v	类似 %v，但输出结构体时会添加字段名
	fmt.Printf("%+v\n", xm) // {name:小明 age:18 gender:male mother:<nil> father:<nil>}

	//%#v	值的 Go 语法表示
	fmt.Printf("%#v\n", xm) //object.Profile{name:"小明", age:18, gender:"male", mother:(*object.Profile)(nil), father:(*object.Profile)(nil)}
}

// --------------------------------------标签tag------------------------------------------
// 字段上还可以额外再加一个属性，用反引号（Esc键下面的那个键）包含的字符串，称之为 Tag，也就是标签。
type Person struct {
	Name string `json:"name" label:"Name is: "`
	//只要发现被omitempty标记的属性 为 false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串中的一种，就会被忽略。
	Age  int    `json:"age,omitempty" label:"Age is: "`
	Addr string `json:"addr,omitempty" label:"Gender is: " default:"unknown"`
}

func TestTag(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  0,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr且age为0，就不会打印了Addr和age属性
	fmt.Printf("%s\n", data1)

	// ================

	p2 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)
}

func TestGetTagByReflect(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	// 三种获取 field
	field_1, _ := reflect.TypeOf(p1).FieldByName("Age")
	//field_2 := reflect.ValueOf(p1).Type().Field(0)         // i 表示第几个字段
	//field_3 := reflect.ValueOf(&p1).Elem().Type().Field(1) // i 表示第几个字段
	tag := field_1.Tag
	fmt.Println(tag)

	// 打印所有字段的 tag
	fmt.Println("打印所有字段的 tag")
	ts := reflect.TypeOf(Person{})
	for i := 0; i < ts.NumField(); i++ {
		f := ts.Field(i)
		fmt.Println(f.Tag)
	}

	// 获取键值对
	//Get 是对 Lookup 的封装
	labelValue := tag.Get("json")
	labelValue_1, _ := tag.Lookup("json")

	fmt.Println(labelValue)
	fmt.Println(labelValue_1)

}

func Print(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}

// 通过tag操作实现打印：
// Name is: MING
// Age is: 29
// Gender is: unknown
func TestPrint(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
		Addr: "",
	}
	Print(p1)
}

// ========================================通过嵌入结构体扩展方法========================================
type Point struct{ X, Y int }

type ColoredPoint struct {
	Point
}

func (p *Point) PointMethod(x int) int {
	return p.X + p.Y + x
}

func TestInvokePointMethod(t *testing.T) {
	cp := &ColoredPoint{
		Point: Point{1, 2},
	}
	// 这里类型为 ColoredPoint 的变量 cp 能直接调用内嵌类型 Point 的方法
	fmt.Println(cp.PointMethod(3))
}

// 方法值：方法也可以作为值，类似函数值
func TestMethodValue(t *testing.T) {
	p := &Point{1, 2}
	point_method := p.PointMethod
	// 这里传入的参数 p 是接收器
	fmt.Println(point_method(3))
}

// 方法表达式：将选择器作为参数来调用，可以在决定调用哪个函数时使用
func TestMethodExpression(t *testing.T) {
	point_method := (*Point).PointMethod
	p := &Point{1, 2}
	// 这里传入的参数 p 是接收器
	fmt.Println(point_method(p, 3))
}
