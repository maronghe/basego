/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package _sync_pool_test

import (
	"basego/_sync_pool"
	"fmt"
	"runtime/debug"
	"testing"
)

//
func TestMyPool(t *testing.T) {

	defer debug.SetGCPercent(debug.SetGCPercent(-1))


	p :=_sync_pool.Pool.Get().(*_sync_pool.Person)
	fmt.Println(p)

	p.Name = "hello"

	_sync_pool.Pool.Put(p)
	fmt.Println(p)

	fmt.Println(_sync_pool.Pool.Get().(*_sync_pool.Person))
	fmt.Println(_sync_pool.Pool.Get().(*_sync_pool.Person))

	fmt.Printf("%s",_sync_pool.Pool.Get().(*_sync_pool.Person))

}

