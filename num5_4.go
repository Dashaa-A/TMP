package main

import "fmt"

type Mediator interface {
	SendMessage(message string, colleague Colleague)
}

type Colleague interface {
	ReceiveMessage(message string)
	SendMessage(message string)
	SetMediator(mediator Mediator)
}

type SomeColleague struct {
	name     string
	mediator Mediator
}

func NewSomeColleague(name string) *SomeColleague {
	return &SomeColleague{name: name}
}

func (sc *SomeColleague) SetMediator(mediator Mediator) {
	sc.mediator = mediator
}

func (sc *SomeColleague) SendMessage(message string) {
	if sc.mediator != nil {
		sc.mediator.SendMessage(message, sc)
	}
}

func (sc *SomeColleague) ReceiveMessage(message string) {
	fmt.Printf("%s получил: %s\n", sc.name, message)
}

type SomeMediator struct {
	colleagues []Colleague
}

func NewSomeMediator() *SomeMediator {
	return &SomeMediator{colleagues: []Colleague{}}
}

func (sm *SomeMediator) AddColleague(colleague Colleague) {
	sm.colleagues = append(sm.colleagues, colleague)
	colleague.SetMediator(sm)
}

func (sm *SomeMediator) SendMessage(message string, sender Colleague) {
	for _, col := range sm.colleagues {
		if col != sender {
			col.ReceiveMessage(message)
		}
	}
}

func main() {
	mediator := NewSomeMediator()
	colleague1 := NewSomeColleague("Коллега 1")
	colleague2 := NewSomeColleague("Коллега 2")

	mediator.AddColleague(colleague1)
	mediator.AddColleague(colleague2)

	colleague1.SendMessage("Привет от Коллеги 1")
	colleague2.SendMessage("Привет от Коллеги 2")
}
