package main

import (
	"fmt"
)

type p1 struct {
}

func (p *p1) getMessage(data string) {
	fmt.Println("p1 get message:", data)
}

type p2 struct {
}

func (p *p2) getMessage(data string) {
	fmt.Println("p2 get message:", data)
}

type p3 struct {
}

func (p *p3) getMessage(data string) {
	fmt.Println("p3 get message:", data)
}

type Message struct {
	p1 *p1
	p2 *p2
	p3 *p3
}

func (m *Message) sendMessage(i interface{}, data string) {
	switch i.(type) {
	case *p1:
		m.p2.getMessage(data)
	case *p2:
		m.p1.getMessage(data)
	case *p3:
		m.p1.getMessage(data)
		m.p2.getMessage(data)
	}
}

func ExampleMediator() {
	fmt.Println("----------中介者模式 Mediator Start----------")

	message := &Message{}

	p1 := &p1{}
	p2 := &p2{}
	p3 := &p3{}

	message.sendMessage(p1, "p1")
	message.sendMessage(p2, "p2")
	message.sendMessage(p3, "p3")

	fmt.Println("----------中介者模式 Mediator End----------")
}
