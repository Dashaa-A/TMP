package main

import "fmt"

type Target interface {
	Request()
}

type Adaptee struct{}

func (a Adaptee) SpecificRequest() {
	fmt.Println("Запрос от Adapter'a")
}

type Adapter struct {
	adaptee Adaptee
}

func NewAdapter(adaptee Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() {
	a.adaptee.SpecificRequest()
}

func main() {
	adaptee := Adaptee{}
	adapter := NewAdapter(adaptee)

	adapter.Request()
}
