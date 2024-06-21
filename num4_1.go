package main

import "fmt"

type Iterator struct {
	container []int
	index     int
}

func NewIterator(container []int) *Iterator {
	return &Iterator{container: container, index: 0}
}

func (it *Iterator) Next() {
	it.index++
}

func (it *Iterator) Value() int {
	return it.container[it.index]
}

func (it *Iterator) HasMore() bool {
	return it.index < len(it.container)
}

func main() {
	vec := []int{17, 22, 34, 91, 65}

	it := NewIterator(vec)

	for it.HasMore() {
		fmt.Print(it.Value(), " ")
		it.Next()
	}
}
