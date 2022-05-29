package main

import (
	"fmt"
)

type Type struct {
	name string
}

func (t *Type) Clone() *Type {
	tc := *t
	return &tc
}

func ExamplePrototype() {
	fmt.Println("----------原型模式 Prototype Start----------")
	t1 := &Type{
		name: "type1",
	}

	t2 := t1.Clone()
	if t1 == t2 {
		fmt.Println("error! get clone not working")
	}
	fmt.Println("t1地址:", &t1)
	fmt.Println("t2地址:", &t2)

	fmt.Println("----------原型模式 Prototype End----------")
}
