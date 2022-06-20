package main

import (
	"fmt"
)

type Travel struct {
	name     string
	strategy Strategy
}

type Strategy interface {
	traffic(*Travel)
}

func (t *Travel) traffic() {
	t.strategy.traffic(t)
}

func NewTravel(name string, strategy Strategy) *Travel {
	return &Travel{name: name, strategy: strategy}
}

type Walk struct {
}

func (w *Walk) traffic(t *Travel) {
	fmt.Println(t.name + " walk")
}

type Ride struct {
}

func (r *Ride) traffic(t *Travel) {
	fmt.Println(t.name + " ride")
}

type Drive struct {
}

func (d *Drive) traffic(t *Travel) {
	fmt.Println(t.name + " drive")
}

func ExampleStrategy() {
	fmt.Println("----------策略模式 Strategy Start----------")

	walk := &Walk{}
	Travel1 := NewTravel("张三", walk)
	Travel1.traffic()

	ride := &Ride{}
	Travel2 := NewTravel("李四", ride)
	Travel2.traffic()

	drive := &Drive{}
	Travel3 := NewTravel("王五", drive)
	Travel3.traffic()

	fmt.Println("----------策略模式 Strategy End----------")
}
