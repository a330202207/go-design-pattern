package main

import "fmt"

type Pay interface {
	Pay(string) int
}

type PayReq struct {
	OrderId string
}

type APayReq struct {
	PayReq
}

func (p *APayReq) Pay() string {
	fmt.Println("订单号:", p.OrderId)
	return "APay支付成功"
}

type BPayReq struct {
	PayReq
	Uid int64
}

func (p *BPayReq) Pay() string {
	fmt.Println("订单号:", p.OrderId)
	fmt.Println("Uid:", p.Uid)
	return "BPay支付成功"
}

func exampleAPay() {
	aPay := APayReq{PayReq{OrderId: "123456789"}}
	aPay.Pay()
}

func exampleBPay() {
	PayReq := PayReq{OrderId: "987654321"}
	bPay := BPayReq{Uid: 1, PayReq: PayReq}
	bPay.Pay()
}

func ExampleFactory() {
	fmt.Println("----------工厂方法模式 Factory Method Start----------")
	exampleAPay()
	exampleBPay()
	fmt.Println("----------工厂方法模式 Factory Method End----------")
}
