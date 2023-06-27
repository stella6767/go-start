package main

import "fmt"

/**
Generics

대괄호 사용
*/

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// 타입 제한자
//직접 만들필요없이 Go Contraints 에서 import

type Integer interface {
	int | int8 | int16 | int32 | int64
}

func min[T Integer | float32 | float64](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func main() {

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

	//min(3, "1")
	min(3, 1.2)

}
