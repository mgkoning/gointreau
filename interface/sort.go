package main

import (
	"fmt"
	"sort"
)

func main() {
	persons := []person{
		person{95, "Windows", 0},
		person{34, "Michiel", 185},
		person{13, "Cartman", 120},
		person{53, "Rutger", 182},
	}

	somePersons := persons[1:3]
	fmt.Println(somePersons)

	sort.Sort(byAge(persons))

	fmt.Println(persons)
	fmt.Println(somePersons)

	doPrint(persons[0])
}

type printer interface {
	Print() string
}

func doPrint(subject printer) {
	fmt.Println(subject.Print())
}

func (p person) Print() string {
	return fmt.Sprintf(
		"Person %v has age %v years and length %v cm",
		p.name, p.age, p.length)
}

type person struct {
	age    int
	name   string
	length int
}

type byAge []person

func (b byAge) Len() int {
	return len(b)
}

func (b byAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
