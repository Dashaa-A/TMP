package main

import "fmt"

type Visitor interface {
	VisitElementA(element ElementA)
	VisitElementB(element ElementB)
}

type Element interface {
	Accept(visitor Visitor)
}

type ElementA struct{}

func (e ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

func (e ElementA) OperationA() {
	fmt.Println("Операция с элементом A")
}

type ElementB struct{}

func (e ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

func (e ElementB) OperationB() {
	fmt.Println("Операция с элементом B")
}

type SomeVisitor struct{}

func (v SomeVisitor) VisitElementA(element ElementA) {
	element.OperationA()
}

func (v SomeVisitor) VisitElementB(element ElementB) {
	element.OperationB()
}

func main() {
	elementA := ElementA{}
	elementB := ElementB{}
	visitor := SomeVisitor{}

	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
