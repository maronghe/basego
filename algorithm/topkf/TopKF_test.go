/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package topkf_test

import (
	"basego/algorithm/topkf"
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1,1,2,2,4,5}
	k := 2
	fmt.Println(topkf.TopKFrequent(nums,k))
}

func BenchmarkTopKFrequent(b *testing.B) {
	nums := []int{1,1,2,2,4,5}
	k := 2

	for i := 0 ; i < b.N ; i ++{
		topkf.TopKFrequent(nums,k)
	}

}
