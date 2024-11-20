package internal

import "fmt"

func Type() {
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// List 可以保存任何類型的值
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) insert(value T) *List[T] {
	return &List[T]{val: value, next: l}
}

func (l *List[T]) print() {
	for current := l; current != nil; current = current.next {
		fmt.Println(current.val)
	}
}

func ListTest() {
	var l *List[int]
	l = l.insert(1)
	l = l.insert(2)
	l = l.insert(3)

	l.print()

	var strList *List[string]
	strList = strList.insert("A")
	strList = strList.insert("B")
	strList = strList.insert("C")

	strList.print()
}
