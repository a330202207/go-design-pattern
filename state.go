package main

import (
	"fmt"
)

type state interface {
	open(*door)
	close(*door)
}

type door struct {
	opened  state
	closed  state
	damaged state

	currentState state
}

func (d *door) open() {
	d.currentState.open(d)
}

func (d *door) close() {
	d.currentState.close(d)
}

func (d *door) setState(s state) {
	d.currentState = s
}

type opened struct {
}

func (o *opened) open(d *door) {
	fmt.Println("门已开启")
}

func (o *opened) close(d *door) {
	fmt.Println("关闭成功")
}

type closed struct {
}

func (c *closed) open(d *door) {
	fmt.Println("开启成功")
}

func (c *closed) close(d *door) {
	fmt.Println("门已关闭")
}

type damaged struct {
}

func (de *damaged) open(d *door) {
	fmt.Println("门已损坏，无法开启")
}

func (de *damaged) close(d *door) {
	fmt.Println("门已损坏，无法关闭")
}

func ExampleState() {
	fmt.Println("----------状态模式 State Start----------")
	door := &door{}

	opened := &opened{}
	door.setState(opened)
	door.open()
	door.close()

	closed := &closed{}
	door.setState(closed)
	door.open()
	door.close()

	damaged := &damaged{}
	door.setState(damaged)
	door.open()
	door.close()

	fmt.Println("----------状态模式 State End----------")
}
