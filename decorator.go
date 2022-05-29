package main

import (
	"fmt"
)

type pizza interface {
	getPrice() int
}

type base struct {
}

func (b *base) getPrice() int {
	return 15
}

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 10
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 20
}

func ExampleDecorator() {
	fmt.Println("----------装饰模式 Decorator Start----------")

	pizza := &base{}

	pizzaWithTomato := &tomatoTopping{pizza: pizza}

	fmt.Printf("tomatoPizza price is %d\n", pizzaWithTomato.getPrice())

	pizzaWithCheese := &cheeseTopping{pizza: pizza}
	fmt.Printf("cheesePizza price is %d\n", pizzaWithCheese.getPrice())

	fmt.Println("----------装饰模式 Decorator End----------")
}
