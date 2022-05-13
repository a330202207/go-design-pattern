package main

import (
	"fmt"
	"sync"
)

// singleton 单例实例
type singleton struct {
	Value int
}

type Singleton interface {
	getValue() int
}

func (s singleton) getValue() int {
	return s.Value
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 构造方法，用于获取单例模式对象
func GetInstance(v int) Singleton {
	once.Do(func() {
		fmt.Println("Value:", v)
		instance = &singleton{Value: v}
	})

	return instance
}

func ExampleSingleton() {
	fmt.Println("----------单例模式 Singleton----------")
	ins1 := GetInstance(1)
	ins2 := GetInstance(2)
	if ins1 == ins2 {
		fmt.Println("ins1 的地址是:", &ins1)
		fmt.Println("ins2 的地址是:", &ins2)
		fmt.Println("ins1 和 ins1 实例相等")
	}

	fmt.Println("----------单例模式 Singleton----------")
}
