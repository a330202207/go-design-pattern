package main

import (
	"fmt"
)

type Shape interface {
	accept(visitor)
}

type square struct {
}

func (s *square) accept(v visitor) {
	v.visitorForSquare(s)
}

type circle struct {
}

func (c *circle) accept(v visitor) {
	v.visitorForCircle(c)
}

type visitor interface {
	visitorForSquare(*square)
	visitorForCircle(*circle)
}

type sideCalculator struct {
}

func (sc *sideCalculator) visitorForSquare(s *square) {
	fmt.Println("square side")
}

func (sc *sideCalculator) visitorForCircle(c *circle) {
	fmt.Println("circle side")
}

type radiusCalculator struct {
}

func (r *radiusCalculator) visitorForSquare(s *square) {
	fmt.Println("square radius")
}

func (r *radiusCalculator) visitorForCircle(s *circle) {
	fmt.Println("circle radius")
}

func ExampleVisitor() {
	fmt.Println("----------访问者模式 Visitor Start----------")
	square := &square{}
	circle := &circle{}

	side := &sideCalculator{}

	square.accept(side)
	circle.accept(side)

	radius := &radiusCalculator{}
	square.accept(radius)
	square.accept(radius)

	fmt.Println("----------访问者模式 Visitor End----------")
}
