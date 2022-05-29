package main

import (
	"fmt"
)

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Search(string)
}

const (
	LeafNode = iota
	ComponentNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case ComponentNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (l *Leaf) Search(pre string) {
	fmt.Printf("leaf %s-%s\n", pre, l.Name())
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{childs: make([]Component, 0)}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(child.Parent())
	c.childs = append(c.childs, child)
}

func (c *Composite) Search(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Search(pre)
	}
}

func ExampleObjectTree() {
	fmt.Println("----------对象树模式 Object Tree Start----------")
	root := NewComponent(ComponentNode, "root")
	c1 := NewComponent(ComponentNode, "c1")
	c2 := NewComponent(ComponentNode, "c2")
	c3 := NewComponent(ComponentNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	root.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c3.AddChild(l3)

	root.Search("")

	fmt.Println("----------对象树模式 Object Tree End----------")
}
