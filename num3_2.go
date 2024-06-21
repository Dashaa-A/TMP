package main

import "fmt"

type AbstractClass interface {
	TemplateMethod()
	Step1()
	Step2()
	Step3()
}

type TemplateMethodClass struct{}

func (t TemplateMethodClass) TemplateMethod() {
	t.Step1()
	t.Step2()
	t.Step3()
}

func (t TemplateMethodClass) Step1() {
	fmt.Println("Шаг 1 выполнен")
}

func (t TemplateMethodClass) Step2() {
	fmt.Println("Шаг 2 выполнен")
}

func (t TemplateMethodClass) Step3() {
	fmt.Println("Шаг 3 выполнен")
}

func main() {
	templateMethod := TemplateMethodClass{}
	templateMethod.TemplateMethod()
}
