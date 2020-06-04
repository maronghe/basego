package main

import (
	"fmt"
	"sync"
	"time"
)

// print :
//  logan
//  hello world true

func main() {

	a := "hello world"

	ch := make(chan string, 1)

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		ch <- a
		time.Sleep(2 * time.Second)
		fmt.Println(a)
		wg.Done()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		select {
		case v, ok := <-ch:
			fmt.Println(v, ok)
		}
		wg.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		a = "logan"
		wg.Done()
	}()

	wg.Wait()
}
