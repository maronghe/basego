/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package topkf

import "sort"

// Top K Frequent Element
// nums is [1,1,2,2,4,5]
func TopKFrequent(nums []int, k int) []int {
	magic := make(map[int]int)

	for _, v := range nums {
		magic[v]++
	}

	pairs := make([][2]int, 0, len(magic))
	for v, f := range magic {
		pairs = append(pairs, [2]int{v, f})
	}
	sort.Sort(sort.Reverse(Pair(pairs)))

	//possible get paris is [[1,2] , [2,2] , [4,1] , [5,1]]

	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ret = append(ret, pairs[i][0])
	}

	return ret
}

// 数组长度为2的int的slice
type Pair [][2]int
// 实现Interface interface
func (p Pair) Less(i, j int) bool {return p[i][1] < p[j][1]}

func (p Pair) Swap(i, j int) {p[i], p[j] = p[j], p[i]}

func (p Pair) Len() int {return len(p)}

