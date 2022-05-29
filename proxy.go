package main

import (
	"fmt"
)

type Subject interface {
	Proxy() string
}

type Proxy struct {
	real RealSubject
}

type RealSubject struct{}

func (p Proxy) Proxy() string {
	var res string

	// 在调用真实对象之前，检查缓存，判断权限，等等
	p.real.Pre()

	// 调用真实对象
	p.real.Real()

	// 调用之后的操作，如缓存结果，对结果进行处理，等等
	p.real.After()

	return res
}

func (RealSubject) Real() {
	fmt.Println("real action")
}

func (RealSubject) Pre() {
	fmt.Println("pre action")
}

func (RealSubject) After() {
	fmt.Println("after action")
}

func ExampleProxy() {
	fmt.Println("----------代理模式Proxy Start----------")

	var sub Subject
	sub = &Proxy{}
	sub.Proxy()

	fmt.Println("----------代理模式Proxy End----------")
}
