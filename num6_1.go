package main

import "fmt"

type Processor struct{}

func (p *Processor) Process(callback func()) {
	fmt.Println("Выполняется обработка...")
	callback()
	fmt.Println("Обработка завершена.")
}

func myCallback() {
	fmt.Println("Вызвана функция обратного вызова!")
}

func main() {
	processor := &Processor{}
	processor.Process(myCallback)
}
