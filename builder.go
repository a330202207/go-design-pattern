package main

import (
	"fmt"
)

// Builder 建造者接口
type Builder interface {
	Part1()
	Part2()
	Part3()
}

// Director 管理类
type Director struct {
	builder Builder
}

// NewDirector 构造函数
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 建造
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Building struct{}

func (b *Building) Part1() {
	fmt.Println("part1")
}

func (b *Building) Part2() {
	fmt.Println("part2")
}

func (b *Building) Part3() {
	fmt.Println("part3")
}

func ExampleBuilder() {
	fmt.Println("----------建造者模式 Builder----------")
	builder := &Building{}
	director := NewDirector(builder)
	director.Construct()
	fmt.Println("----------建造者模式 Builder----------")
}
