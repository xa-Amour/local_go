package main

import (
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

/**
设计模式

创建模式: 简单工厂,抽象工厂,单例...
行为模式: 观察者模式...
结构模式: 适配器模式...
同步模式: 信号量模式
 */

func main() {
	defer func() {
		if e := recover(); e != nil{
			println(e)
		}
	}()
	fun := map[string]func(){
		"m1":m1,
		"m2":m2,
		"m3":m3,
		"m4":m4,
		"m5":m5,
	}
	fun[os.Args[1]]()
}

//1.简单工厂
type Product interface {
	Create()
}
type ProductA struct {}
func (ProductA)Create()  {
	println("产品A")
}
type ProductB struct {}
func (ProductB)Create()  {
	println("产品B")
}
func GetProduct(productType string) Product {
	print("这里是总工厂的")
	switch productType {
	case "A":
		return ProductA{}
	case "B":
		return ProductB{}
	default:
		return ProductA{}
	}
}
func m1()  {
	ProA := GetProduct("A")
	ProA.Create()

	ProB := GetProduct("B")
	ProB.Create()
}

//2.抽象工厂(工厂方法)
type ProductsFactory interface {
	GetProduct(productT string) Product
}
type ShenZhenProduct struct {}
func (ShenZhenProduct)GetProduct(pt string) Product {
	print("这里是深圳工厂的")
	switch pt {
	case "A":
		return ProductA{}
	case "B":
		return ProductB{}
	default:
		return ProductA{}
	}
}
type BeijingProduct struct {}
func (BeijingProduct)GetProduct(pt string) Product {
	print("这里是北京工厂的")
	switch pt {
	case "A":
		return ProductA{}
	case "B":
		return ProductB{}
	default:
		return ProductA{}
	}
}
func MakeFactory(where string) ProductsFactory {
	switch where {
	case "SZ":
		return ShenZhenProduct{}
	case "BJ":
		return BeijingProduct{}
	default:
		return ShenZhenProduct{}
	}
}
func m2()  {
	sz := MakeFactory("SZ")
	sz.GetProduct("A").Create()
	sz.GetProduct("B").Create()

	zf := MakeFactory("BJ")
	zf.GetProduct("A").Create()
	zf.GetProduct("B").Create()
}

//3.单例模式
func m3()  {
	//1.使用锁+全局变量
	i1,i2 := m3_1(),m3_1()
	println(i1.ran,i2.ran)

	//2.使用锁+atomic+全局变量+int
	i3,i4 := m3_2(),m3_2()
	println(i3.ran,i4.ran)

	//3.使用sync.Once+全局变量 (原理=atomic)
	i5,i6 := m3_3(),m3_3()
	println(i5.ran,i6.ran)
}
type in struct {
	ran int
}
var instance *in
var mu sync.Mutex
func m3_1() *in {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		instance = &in{rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)}
	}
	return instance
}

var instance1 *in
var i uint32
func m3_2() *in {
	if atomic.LoadUint32(&i) == 1 {
		return instance1
	}
	mu.Lock()
	defer mu.Unlock()
	if i == 0 {
		instance1 = &in{rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)}
		atomic.StoreUint32(&i,1)
	}
	return instance1
}

var instance2 *in
var once sync.Once
func m3_3() *in {
	once.Do(func() {
		instance2 = &in{rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)}
	})
	return instance2
}

//4.消费者-生产者模式
func m4()  {
	ch := make(chan int)
	go produser(ch)
	consumer(ch)
}
func consumer(ch <-chan int)  {
	for c := range ch {
		println(c)
	}
}
func produser(ch chan<-int)  {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

//5.观察者模式
func m5()  {
	oba := new(ObserverA)
	obb := new(ObserverB)

	com := new(Company)
	com.Add(oba)
	com.Add(obb)
	com.change()
}
type Observer interface {
	Receive()
}
type ObserverA struct {}
func (ObserverA)Receive()  {
	println("observerA received")
}
type ObserverB struct {}
func (ObserverB)Receive()  {
	println("observerB received")
}
type Company struct {
	obs []Observer
}
func (c *Company)Add(observer Observer)  {
	c.obs = append(c.obs, observer)
}
func (c *Company)change()  {
	for _,o := range c.obs {
		o.Receive()
	}
}