package main

import "fmt"

type Product struct {
	partA string
	partB string
	partC string
}

func (p *Product) SetPartA(partA string) {
	p.partA = partA
}

func (p *Product) SetPartB(partB string) {
	p.partB = partB
}

func (p *Product) SetPartC(partC string) {
	p.partC = partC
}

func (p *Product) Show() {
	fmt.Printf("Части продукта: %s, %s, %s\n", p.partA, p.partB, p.partC)
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetProduct() *Product
}

type SomeBuilder struct {
	product *Product
}

func NewSomeBuilder() *SomeBuilder {
	return &SomeBuilder{product: &Product{}}
}

func (b *SomeBuilder) BuildPartA() {
	b.product.SetPartA("Часть A")
}

func (b *SomeBuilder) BuildPartB() {
	b.product.SetPartB("Часть B")
}

func (b *SomeBuilder) BuildPartC() {
	b.product.SetPartC("Часть C")
}

func (b *SomeBuilder) GetProduct() *Product {
	return b.product
}

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	builder := NewSomeBuilder()
	director := &Director{}
	director.SetBuilder(builder)

	director.Construct()
	product := builder.GetProduct()
	product.Show()
}
