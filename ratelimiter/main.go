package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	//  基于漏桶算法实现
	rl := ratelimit.New(1) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take() // will block util : time.Second / time.Duration(rate)
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

    // Output:
    // 0 0
    // 1 10ms
    // 2 10ms
    // 3 10ms
    // 4 10ms
    // 5 10ms
    // 6 10ms
    // 7 10ms
    // 8 10ms
    // 9 10ms
}