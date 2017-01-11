package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer timeSince(time.Now())
	defer fmt.Println("asdasd")
	fmt.Println("asdd")
	fmt.Println("bbbbb")

	_, err := os.Open("sadasd.txt")
	defer panic("heeeeelp")
	if err != nil {
		panic(err)
	}

}

func timeSince(t time.Time) {
	fmt.Printf("%v\n", time.Since(t))
}
