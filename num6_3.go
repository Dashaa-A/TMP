package main

import "fmt"

type Component interface {
	Operation()
}

type Leaf struct{}

func (l *Leaf) Operation() {
	fmt.Println("Leaf: выполнение операции.")
}

type Composite struct {
	children []Component
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Operation() {
	fmt.Println("Composite: выполнение операции для всех дочерних компонентов.")
	for _, child := range c.children {
		child.Operation()
	}
}

func main() {
	composite := &Composite{}
	composite.Add(&Leaf{})
	composite.Add(&Leaf{})
	composite.Operation()
}
