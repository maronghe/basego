/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package ds_test

import (
	"basego/ds"
	"fmt"
	"math/rand"
	"testing"
)

func TestBSTFeatures(t *testing.T) {

	/** Let's create flowing BST
						50(2)
					/    \
				10      70(2)
			    \    /  \
					30 	60  80
				 /  \
				20  40         */

	bst := ds.NewBinarySearchTree(50)
	bst.Insert(10).
		Insert(70).
		Insert(30).
		Insert(20).
		Insert(60).
		Insert(40).
		Insert(80).
		Insert(50).
		Insert(70)

	//ds.InOrder(bst)   // 10 20 30 40 50(2) 60 70(2) 80
	//fmt.Println()
	//ds.PreOrder(bst)  // 50(2) 10 30 20 40 70(2) 60 80
	//fmt.Println()
	//ds.PostOrder(bst) // 20 40 30 20 60 80 70(2) 50(2)
	//fmt.Println()

	//fmt.Println(bst.Search(70))
	//fmt.Println(bst.Search(100))

	fmt.Println(bst.Search(70))
	bst.Delete(70)
	fmt.Println(bst.Search(70))
	bst.Delete(70)
	fmt.Println(bst.Search(70))
	fmt.Println(bst)

}

//./binary_search_tree_test.go:60:10: constant 111111111111111111111 overflows int
func TestBSTBoundary(t *testing.T) {

	bst := ds.NewBinarySearchTree(0)
	bst.Insert(-1).
		Insert(111111111111111111111).
		Insert(30)

	fmt.Println(bst)

}




//
//BenchmarkBST-12         93076572                23.4 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         65985549                15.6 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         56842039                20.3 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         85269330                12.5 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         56428350                20.0 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         91609491                14.8 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         86930814                17.6 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         101735659               20.1 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         88029139                18.0 ns/op             0 B/op          0 allocs/op
//BenchmarkBST-12         202366458               23.6 ns/op             0 B/op          0 allocs/op
//
// ### Unstable!
//
func BenchmarkBST(b *testing.B) {

	root := ds.NewBinarySearchTree(50)
	for i := 0 ; i < 1e6 ; i++ {
		root.Insert(rand.Intn(100))
	}

	for i := 0 ; i < b.N ; i++ {
		root.Search(40)
	}

	//fmt.Println(root)
}
