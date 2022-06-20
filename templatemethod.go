package main

import (
	"fmt"
)

type PrintTemplate interface {
	Print(name string)
}

type template struct {
	isTemplate PrintTemplate
	name       string
}

func (t *template) Print() {
	t.isTemplate.Print(t.name)
}

type A struct {
}

func (a *A) Print(name string) {
	fmt.Println("a:" + name)
	// 业务代码……
}

type B struct{}

func (b *B) Print(name string) {
	fmt.Println("b:" + name)
	// 业务代码……
}

func ExampleTemplateMethod() {
	fmt.Println("----------模板方法模式 TemplateMethod Start----------")
	templateA := &A{}
	template := &template{
		isTemplate: templateA,
		name:       "hello world",
	}
	template.Print()

	templateB := &B{}
	template.isTemplate = templateB
	template.Print()

	fmt.Println("----------模板方法模式 TemplateMethod End----------")
}
