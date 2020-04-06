/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package clone

// cle from int slice
func IntSlice(os []int) []int {
	return append(os[:0:0], os...)
}

func IntSlice2(os []int) []int {
	ns := make([]int,len(os))
	copy(os,ns)
	return ns
}
