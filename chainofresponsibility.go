package main

import "fmt"

type department interface {
	execute(*Do)
	setNext(department)
}

type aPart struct {
	next department
}

func (a *aPart) execute(p *Do) {
	if p.aPartDone {
		fmt.Println("aPart Done")
		a.next.execute(p)
		return
	}

	fmt.Println("aPart")
	p.aPartDone = true
	a.next.execute(p)
}

func (a *aPart) setNext(next department) {
	a.next = next
}

type bPart struct {
	next department
}

func (b *bPart) execute(p *Do) {
	if p.bPartDone {
		fmt.Println("bPart Done")
		b.next.execute(p)
		return
	}

	fmt.Println("bPart")
	p.bPartDone = true
	b.next.execute(p)
}

func (b *bPart) setNext(next department) {
	b.next = next
}

type endPart struct {
	next department
}

func (e *endPart) execute(p *Do) {
	if p.bPartDone {
		fmt.Println("endPart Done")
		return
	}

	fmt.Println("endPart")
	p.bPartDone = true
	e.next.execute(p)
}

func (e *endPart) setNext(next department) {
	e.next = next
}

type Do struct {
	aPartDone   bool
	bPartDone   bool
	endPartDone bool
}

func ExampleChain() {
	fmt.Println("----------责任链模式 Chain of Responsibility Start----------")
	startPart := &endPart{}

	aPart := &aPart{}
	aPart.setNext(startPart)

	bPart := &bPart{}
	bPart.setNext(aPart)

	do := &Do{}
	bPart.execute(do)

	fmt.Println("----------责任链模式 Chain of Responsibility End----------")
}
