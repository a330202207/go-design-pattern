package main

import (
	"fmt"
)

type APIA interface {
	TestA() string
}

func NewAPIA() APIA {
	return &apiRunA{}
}

type apiRunA struct {
}

func (a *apiRunA) TestA() string {
	return "A api running"
}

type APIB interface {
	TestB() string
}

func NewAPIB() APIB {
	return &apiRunB{}
}

type apiRunB struct {
}

func (a *apiRunB) TestB() string {
	return "B api running"
}

type API interface {
	Test() string
}

func NewAPI() API {
	return &apiRun{
		a: NewAPIA(),
		b: NewAPIB(),
	}
}

type apiRun struct {
	a APIA
	b APIB
}

func (a *apiRun) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func ExampleFacade() {
	fmt.Println("----------外观模式 Facade Start----------")

	api := NewAPI()
	ret := api.Test()
	fmt.Println(ret)

	fmt.Println("----------外观模式 Facade End----------")
}
