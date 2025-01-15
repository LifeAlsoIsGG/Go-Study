package object

import (
	"fmt"
	"strconv"
	"testing"
)

// 接口定义
type Phone interface {
	call()
}

// 定义结构体
type Nokia struct {
	name string
}

// 结构体实现接口方法，隐式，自动实现
// 一个类型如果定义了接口的所有方法，那它就隐式地实现了该接口。
func (phone *Nokia) call() {
	fmt.Println("我是Nokia打电话")
}

//--------------------------------------多态------------------------------------------

// 只要实现了这两个方法"就是一个商品",必须都实现
type Good interface {
	settleAccount() int
	orderInfo() string
}

type PhoneStruct struct {
	name     string
	quantity int
	price    int
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

// Phone
func (phone PhoneStruct) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone PhoneStruct) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

// FreeGift
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

func TestMultiType(t *testing.T) {
	iPhone := PhoneStruct{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	//这里体现了多态
	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
