package main

import "fmt"

/**
함수와 대입
*/

func main() {

	a, b := 3, 4
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)

	total := add(a, b)
	fmt.Println(total)

	i, i2, i3 := multiReturn()

	fmt.Println(i, i2, i3)

	variadicExample("hello", "it is s golang 가변인자")

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	// 람다
	calc := func(a int, b int) int {
		return a + b
	}

	fmt.Println(calc)       //output : 0x467fa0
	fmt.Println(calc(1, 2)) //output : 3
}

func add(a int, b int) int {
	return a + b
}

func multiReturn() (int, int, int) {
	return 3, 4, 5
}
func variadicExample(s ...string) {
	fmt.Println(s)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
