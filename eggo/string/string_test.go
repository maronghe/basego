package string_test

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	test1()
	//test2()
}

func test1() {
	str := "Hello world"
	// str[1] = '1' // 编译不过

	sliceStr := []byte(str)
	sliceStr[0] = '1'
	fmt.Printf("%s\n", sliceStr)
}

func test2() {
	str := "Hello world"
	sliceStr := *(*[]byte)(unsafe.Pointer(&str))
	sliceStr[0] = '1' // panic
	fmt.Printf("%s\n", sliceStr)
}
