/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package s_test

import (
	"basego/s"
	"fmt"
	"math/rand"
	"testing"
)

var nums []int

func init() {
	nums = make([]int , 1e8)
	for i := 0 ; i < 1e8 ; i++ {
		nums[i] = i
	}
}


func TestBinarySearch(t *testing.T) {
	// nums := []int{-1,0,3,5,9,12}
	target := 9
	fmt.Printf("target %d's index is %d\n", target, s.BinarySearch(nums, target))
}

func BenchmarkBinarySearch(b *testing.B) {

	for i := 0 ; i < b.N ; i++ {
		s.BinarySearch(nums, rand.Intn(1e8))
	}

}


func BenchmarkOSearch(b *testing.B) {
	for i := 0 ; i < b.N ; i++ {
		s.OSearch(nums, rand.Intn(1e8))
	}
}