package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Printf("before sorted : %+v\n", arr)
	arr = BubbleSort(arr)
	fmt.Printf("after sorted : %+v\n", arr)

}

func BenchmarkBubbleSort(b *testing.B) {
	for size := 0; size < 100; size += 10 {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			var arr []int
			for i := 0; i < size; i++ {
				arr = append(arr, rand.Intn(100))
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BubbleSort(arr)
			}
		})
	}

}
