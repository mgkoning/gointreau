package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func check(err error) {
	if err == nil {
		return
	}
	panic(err)
}

func main() {
	fmt.Println("Hello, world!")
	multipleReturnValues()
	slicesAndArrays()
	interfaceImplementation()
	readAFile()
	goRoutinesAndChannels()
}

func multipleReturnValues() {
	div, mod := divWithMod(5, 2)
	fmt.Printf("Div %v, mod %v\n", div, mod)
}

func divWithMod(a, b int) (div, mod int) {
	div = a / b
	mod = a % b
	return
}

func slicesAndArrays() {
	// array := [3]int{1, 2, 3}
	// array2 := [4]int{1, 2, 3, 4}
	// array = array2

	slice := make([]int, 0)
	slice = append(slice, 14)
	slice2 := slice
	slice2[0] = 13
	fmt.Println("Slice:", slice)
}

func interfaceImplementation() {
	persons := []person{
		{95, "Windows"},
		{34, "Michiel"},
		{10, "Cartman"},
	}
	fmt.Println(persons)
	var data sort.Interface = byAge(persons)
	sort.Sort(data)
	fmt.Println(persons)
}

type person struct {
	age  int
	name string
}

type byAge []person

func (b byAge) Len() int {
	return len(b)
}

func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byAge) Less(i, j int) bool {
	return b[i].age < b[j].age
}

func readAFile() {
	f, err := os.Open("some.txt")
	check(err)
	defer f.Close()

	bytes := make([]byte, 5)
	_, err = io.ReadFull(f, bytes)
	check(err)

	text := string(bytes)
	fmt.Println("Konnichiwa", text)
}

func goRoutinesAndChannels() {
	done := make(chan bool)
	go waitAndPrint(done)
	now := time.Now()
	fmt.Println("Fired!")
	//_ = <-done
	for {
		select {
		case <-done:
			fmt.Println("Received after", time.Since(now))
			return
		default:
			fmt.Println("Nothing after", time.Since(now))
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func waitAndPrint(done chan bool) {
	time.Sleep(time.Second * 15)
	fmt.Println("Woke up!")
	done <- true
}
