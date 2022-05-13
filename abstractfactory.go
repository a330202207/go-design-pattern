package main

import (
	"fmt"
)

// SaveArticle 抽象模式工厂接口
type SaveArticle interface {
	CreateProse() Prose                 // 创建散文
	CreateAncientPoetry() AncientPoetry // 创建古诗
}

// SaveRedis Redis保存
type SaveRedis struct{}

// SaveMysql Mysql保存
type SaveMysql struct{}

// RedisProse Redis工厂
type RedisProse struct{}

// MysqlProse Mysql工厂
type MysqlProse struct{}

// Prose 散文
type Prose interface {
	SaveProse()
}

// AncientPoetry 古诗
type AncientPoetry interface {
	SaveAncientPoetry()
}

// CreateProse 创建Redis散文
func (*SaveRedis) CreateProse() Prose {
	return &RedisProse{}
}

// CreateAncientPoetry 创建Redis古诗
func (*SaveRedis) CreateAncientPoetry() AncientPoetry {
	return &RedisProse{}
}

// CreateProse 创建Mysql散文
func (*SaveMysql) CreateProse() Prose {
	return &MysqlProse{}
}

// CreateAncientPoetry 创建Mysql古诗
func (*SaveMysql) CreateAncientPoetry() AncientPoetry {
	return &MysqlProse{}
}

func (*RedisProse) SaveProse() {
	fmt.Println("Redis Save Prose")
}

func (*RedisProse) SaveAncientPoetry() {
	fmt.Println("Redis Save Ancient Poetry")
}

func (*MysqlProse) SaveProse() {
	fmt.Println("Mysql Save Prose")
}

func (*MysqlProse) SaveAncientPoetry() {
	fmt.Println("Mysql Save Ancient Poetry")
}

func Save(saveArticle SaveArticle) {
	saveArticle.CreateProse().SaveProse()
	saveArticle.CreateAncientPoetry().SaveAncientPoetry()
}

func ExampleSaveRedis() {
	var factory SaveArticle
	factory = &SaveRedis{}
	Save(factory)
}

func ExampleSaveMysql() {
	var factory SaveArticle
	factory = &SaveMysql{}
	Save(factory)
}

// ExampleAbstractFactory 抽象工厂模式
func ExampleAbstractFactory() {
	fmt.Println("----------抽象工厂模式 Abstract Factory----------")
	ExampleSaveRedis()
	ExampleSaveMysql()
	fmt.Println("----------抽象工厂模式 Abstract Factory----------")
}
