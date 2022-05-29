package main

import (
	"fmt"
)

// ColorFlyweightFactory 颜色工厂
type ColorFlyweightFactory struct {
	maps map[string]*ColorFlyweight
}

var colorFactory *ColorFlyweightFactory

func GetColorFlyweightFactory() *ColorFlyweightFactory {
	if colorFactory == nil {
		colorFactory = &ColorFlyweightFactory{maps: make(map[string]*ColorFlyweight)}
	}

	return colorFactory
}

func (f *ColorFlyweightFactory) Get(filename string) *ColorFlyweight {
	color := f.maps[filename]
	if color == nil {
		color = NewColorFlyweight(filename)
		f.maps[filename] = color
	}

	return color
}

type ColorFlyweight struct {
	data string
}

func NewColorFlyweight(filename string) *ColorFlyweight {
	data := fmt.Sprintf("color data %s", filename)
	return &ColorFlyweight{data: data}
}

type ColorViewer struct {
	*ColorFlyweight
}

func NewColorViewer(name string) *ColorViewer {
	color := GetColorFlyweightFactory().Get(name)
	return &ColorViewer{ColorFlyweight: color}
}

func ExampleFlyweight() {
	fmt.Println("----------享元模式 Flyweight Start----------")

	viewer1 := NewColorViewer("blue")
	fmt.Println("viewer1:", viewer1)
	viewer2 := NewColorViewer("blue")
	fmt.Println("viewer2:", viewer2)

	fmt.Println("----------享元模式 Flyweight End----------")
}
