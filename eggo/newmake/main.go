package main

import (
	"fmt"
)

type P struct {
	n   int
	arr []int
}

func main() {

	p := new(P)
	fmt.Printf("%p\n", p)
	fmt.Println(p)

	p2 := &P{}
	fmt.Printf("%p\n", p2)
	fmt.Println(p2)

	ps := new([]P)
	fmt.Printf("%p", ps)

}
