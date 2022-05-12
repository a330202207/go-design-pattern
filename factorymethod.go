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
	// todo
	fmt.Println(p.OrderId)
	return "APay支付成功"
}

type BPayReq struct {
	PayReq
	Uid int64
}

func (p *BPayReq) Pay() string {
	// todo
	fmt.Println(p.OrderId)
	fmt.Println(p.Uid)
	return "BPay支付成功"
}
