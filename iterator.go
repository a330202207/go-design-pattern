package main

import "fmt"

type collection interface {
	createIterator() iterator
}

type part struct {
	title  string
	number int
}

type partCollection struct {
	part
	parts []*part
}

func (p *partCollection) createIterator() iterator {
	return &partIterator{
		parts: p.parts,
	}
}

type iterator interface {
	hasNext() bool
	getNext() *part
}

type partIterator struct {
	index int
	parts []*part
}

func (p *partIterator) hasNext() bool {
	if p.index < len(p.parts) {
		return true
	}

	return false
}

func (p *partIterator) getNext() *part {
	if p.hasNext() {
		part := p.parts[p.index]
		p.index++
		return part
	}

	return nil
}

func ExampleIterator() {
	fmt.Println("----------迭代器模式 Iterator Start----------")
	part1 := &part{
		title:  "part1",
		number: 10,
	}

	part2 := &part{
		title:  "part2",
		number: 20,
	}

	part3 := &part{
		title:  "part3",
		number: 30,
	}

	partCollection := &partCollection{
		parts: []*part{part1, part2, part3},
	}

	iterator := partCollection.createIterator()

	for iterator.hasNext() {
		part := iterator.getNext()
		fmt.Println(part)
	}

	fmt.Println("----------迭代器模式 Iterator End----------")
}
