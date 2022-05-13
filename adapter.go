package main

import (
	"fmt"
)

// Cm 厘米
type Cm interface {
	getLength(float64) float64
}

// M 米
type M interface {
	getLength(float64) float64
}

func NewM() M {
	return &getLengthM{}
}

type getLengthM struct{}

func (*getLengthM) getLength(cm float64) float64 {
	return cm / 10
}

func NewCm() Cm {
	return &getLengthCm{}
}

type getLengthCm struct{}

func (a *getLengthCm) getLength(m float64) float64 {
	return m * 10
}

// LengthAdapter 适配器
type LengthAdapter interface {
	getLength(string, float64) float64
}

// NewLengthAdapter 实例适配器
func NewLengthAdapter() LengthAdapter {
	return &getLengthAdapter{}
}

type getLengthAdapter struct{}

// getLength 获取长度
func (*getLengthAdapter) getLength(isType string, into float64) float64 {
	if isType == "m" {
		return NewM().getLength(into)
	}
	return NewCm().getLength(into)
}

func ExampleAdapter() {
	fmt.Println("----------适配器模式 Adapter----------")

	into := 10.5
	getLengthM := NewLengthAdapter().getLength("m", into)
	getLengthCm := NewLengthAdapter().getLength("cm", into)
	fmt.Println("getLengthM:", getLengthM)
	fmt.Println("getLengthCm:", getLengthCm)

	fmt.Println("----------适配器模式 Adapter----------")
}
