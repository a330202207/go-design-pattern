package main

import (
	"fmt"
)

type Memento interface {
}

type Text struct {
	context string
}

type textMemento struct {
	context string
}

func (t *Text) Write(content string) {
	t.context = content
}

func (t *Text) Save() Memento {
	return &textMemento{context: t.context}
}

func (t *Text) Load(m Memento) {
	tm := m.(*textMemento)
	t.context = tm.context
}

func (t *Text) Show() {
	fmt.Println("content:", t.context)
}

func ExampleMemento() {
	fmt.Println("----------备忘录模式 Memento Start----------")

	text := &Text{context: "Hello World"}

	text.Show()
	progress := text.Save()

	text.Write("How Are You")
	text.Show()

	text.Load(progress)
	text.Show()

	fmt.Println("----------备忘录模式 Memento End----------")
}
