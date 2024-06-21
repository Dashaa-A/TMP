package main

import "fmt"

type AbstractProductA interface {
	MethodA()
}

type ProductA1 struct{}

func (p ProductA1) MethodA() {
	fmt.Println("Продукт A1 метод A")
}

type ProductA2 struct{}

func (p ProductA2) MethodA() {
	fmt.Println("Продукт A2 метод A")
}

type AbstractProductB interface {
	MethodB()
}

type ProductB1 struct{}

func (p ProductB1) MethodB() {
	fmt.Println("Продукт B1 метод B")
}

type ProductB2 struct{}

func (p ProductB2) MethodB() {
	fmt.Println("Продукт B2 метод B")
}

type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

type ConcreteFactory1 struct{}

func (f ConcreteFactory1) CreateProductA() AbstractProductA {
	return ProductA1{}
}

func (f ConcreteFactory1) CreateProductB() AbstractProductB {
	return ProductB1{}
}

type ConcreteFactory2 struct{}

func (f ConcreteFactory2) CreateProductA() AbstractProductA {
	return ProductA2{}
}

func (f ConcreteFactory2) CreateProductB() AbstractProductB {
	return ProductB2{}
}

func main() {
	factory1 := ConcreteFactory1{}
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()

	productA1.MethodA()
	productB1.MethodB()

	factory2 := ConcreteFactory2{}
	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()

	productA2.MethodA()
	productB2.MethodB()
}
