package main

import (
	"fmt"
)

func main() {
	_, a := doSomething()
	fmt.Println(a)
	arrays()
}

func doSomething() (one, other bool) {
	one = true
	other = true
	return one, other
}

func arrays() {
	x := make([]int, 0, 60)
	fmt.Println(x)
	x = append(x, 1)
	fmt.Println(x)
	y := []int{1, 2, 3, 4}
	z := append(x, append([]int{5}, y...)...)
	z[0] = 30
	fmt.Println(x, z)
	a := z[2:4]
	b := z[1:]
	fmt.Println(a, b)

	c := make(map[int]string)
	c[4] = "boem"
	c[15] = "asdasd"
	c[10] = "asd"
	fmt.Println(c)
	fmt.Println(c)
	for elem := range c {
		fmt.Println(elem)
	}
	for elem := range c {
		fmt.Println(elem)
	}

	coordinates := make(map[coordinate]string)
	coordinates[coordinate{1, 2}] = "Tekst"
	coordinates[coordinate{1, 2}] = "Some other text"
	fmt.Println(coordinates)
}

type coordinate struct {
	x int
	y int
}
