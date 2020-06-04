package main

import (
	"fmt"
	"sync"
)

// 1 -> 2 -> 3
// r w r w r
func main() {

	wg := sync.WaitGroup{}
	wg.Add(5)
	rw := sync.RWMutex{}
	i := 1

	go func() {
		rw.RLock()
		fmt.Println(i)
		rw.RUnlock()
		wg.Done()
	}()

	go func() {
		rw.Lock()
		i = 2
		rw.Unlock()
		wg.Done()
	}()

	go func() {
		rw.RLock()
		fmt.Println(i)
		rw.RUnlock()
		wg.Done()
	}()

	go func() {
		rw.Lock()
		i = 3
		rw.Unlock()
		wg.Done()
	}()

	go func() {
		rw.RLock()
		fmt.Println(i)
		rw.RUnlock()
		wg.Done()
	}()

	wg.Wait()

}
