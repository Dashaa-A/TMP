package main

import "fmt"

type Subject interface {
	Request()
}

type RealSubject struct{}

func (rs *RealSubject) Request() {
	fmt.Println("RealSubject: Обработка запроса.")
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Request() {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	fmt.Println("Proxy: Дополнительная логика перед выполнением запроса.")
	p.realSubject.Request()
	fmt.Println("Proxy: Дополнительная логика после выполнения запроса.")
}

func main() {
	proxy := &Proxy{}
	proxy.Request()
}
