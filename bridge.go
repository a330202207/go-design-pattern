package main

import (
	"fmt"
)

// SendMessage 定义发送信息
type SendMessage interface {
	send(text, to string)
}

type sms struct{}

func NewSms() SendMessage {
	return &sms{}
}

func (*sms) send(text, to string) {
	fmt.Println(fmt.Sprintf("send %s to %s sms", text, to))
}

type email struct{}

func NewEmail() SendMessage {
	return &email{}
}

func (*email) send(text, to string) {
	fmt.Println(fmt.Sprintf("send %s to %s email", text, to))
}

type systemA struct {
	method SendMessage
}

func NewSystemA(method SendMessage) *systemA {
	return &systemA{
		method: method,
	}
}

func (s *systemA) SendMessage(text, to string) {
	s.method.send(fmt.Sprintf("[System A] %s", text), to)
}

type systemB struct {
	method SendMessage
}

func NewSystemB(method SendMessage) *systemB {
	return &systemB{
		method: method,
	}
}

func (s *systemB) SendMessage(text, to string) {
	s.method.send(fmt.Sprintf("[System B] %s", text), to)
}

func ExampleSystemA() {
	NewSystemA(NewSms()).SendMessage("hello", "world")
	NewSystemA(NewEmail()).SendMessage("hello", "world")
}

func ExampleSystemB() {
	NewSystemB(NewSms()).SendMessage("hello", "world")
	NewSystemB(NewEmail()).SendMessage("hello", "world")
}

func ExampleBridge() {
	fmt.Println("----------桥接模式 Bridge Start----------")
	ExampleSystemA()
	ExampleSystemB()
	fmt.Println("----------桥接模式 Bridge End----------")
}
