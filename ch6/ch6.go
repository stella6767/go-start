package main

import "fmt"

/**
array, slice
*/

func main() {

	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	c := [5]int{2: 3, 4: 30}
	fmt.Println(c)

	s := make([]int, 3)
	s = append(s, 11)
	s = append(s, 12, 13)
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))

	d := make([]int, len(s))
	copy(d, s)
	fmt.Println("cpy:", d)

	l := s[2:5]
	fmt.Println("sl1:", l)

	nums := []int{2, 3, 4}

	for _, num := range nums {
		fmt.Println("num:", num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
