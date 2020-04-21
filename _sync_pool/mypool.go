/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package _sync_pool

import (
	"fmt"
	"sync"
)

var Pool *sync.Pool

type Person struct {
	Name string
}

func init() {
	Pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("create object.")
			return &Person{}
		},
	}
}
