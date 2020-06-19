package main

import (
	"fmt"
	"time"

	"go.uber.org/ratelimit"
	rl "github.com/juju/ratelimit"
)

func main() {
	tokenBucket()
}

func test1() {

	//  基于漏桶算法实现
	rl := ratelimit.New(1) // per second

    prev := time.Now()
    for i := 0; i < 10; i++ {
        now := rl.Take() // will block util : time.Second / time.Duration(rate)
        fmt.Println(i, now.Sub(prev))
        prev = now
    }

}

func tokenBucket(){
	b := rl.NewBucket(time.Second, 10) // 内部维护一个容量总数，定时向其中添加数据
	fmt.Println(b.Available())
	go func(){
		for {
			b.Take(2)
			fmt.Println("take 2 " , b.Available())
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(b.Available())
}