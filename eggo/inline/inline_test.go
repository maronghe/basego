package frame_test

import (
	"testing"
)

var Result int

// 叶子内联：调用栈的底部直接内联到调用方的行为。
func BenchmarkYezi(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, i)
	}
	Result = r

	//var r int
	//for i := 0; i < b.N; i++ {
	//	if -1 > i {
	//		r = -1
	//	} else {
	//		r = i
	//	}
	//}
	//Result = r

}

// go:noinline
// 消除函数调用本身的开销
// 允许编译器更有效的应用其他优化策略
// 内联限制在40个表达式以内，并不能有
// loop select switch... 等复杂操作
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
