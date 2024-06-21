package main

import "fmt"

type IStrategy interface {
	Execute()
}

type Strategy1 struct{}

func (s Strategy1) Execute() {
	fmt.Println("Выполнение стратегии 1")
}

type Strategy2 struct{}

func (s Strategy2) Execute() {
	fmt.Println("Выполнение стратегии 2")
}

type Context struct {
	strategy IStrategy
}

func (c *Context) SetStrategy(newStrategy IStrategy) {
	c.strategy = newStrategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func main() {
	strategy1 := Strategy1{}
	strategy2 := Strategy2{}

	context := Context{strategy: strategy1}
	context.ExecuteStrategy()

	context.SetStrategy(strategy2)
	context.ExecuteStrategy()
}
