package object

import (
	"fmt"
	"testing"
)

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
