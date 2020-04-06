/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */



package clone_test

import (
	"basego/clone"
	"math/rand"
	"testing"
)


//→  go test -run="" -bench=. -benchmem -v -count=10                                                                                     master ✖︎ [12:07:28]
//goos: darwin
//goarch: amd64
//pkg: basego/clone
//BenchmarkIntSlice
//BenchmarkIntSlice-12                1630            764160 ns/op         8031313 B/op          1 allocs/op
//BenchmarkIntSlice-12                1722            654995 ns/op         8029830 B/op          1 allocs/op
//BenchmarkIntSlice-12                1857            640737 ns/op         8027922 B/op          1 allocs/op
//BenchmarkIntSlice-12                1858            710186 ns/op         8027910 B/op          1 allocs/op
//BenchmarkIntSlice-12                1629            663210 ns/op         8031328 B/op          1 allocs/op
//BenchmarkIntSlice-12                1743            684228 ns/op         8029515 B/op          1 allocs/op
//BenchmarkIntSlice-12                1592            766925 ns/op         8031973 B/op          1 allocs/op
//BenchmarkIntSlice-12                1508            704923 ns/op         8033554 B/op          1 allocs/op
//BenchmarkIntSlice-12                1722            760476 ns/op         8029830 B/op          1 allocs/op
//BenchmarkIntSlice-12                1574            809989 ns/op         8032298 B/op          1 allocs/op
//BenchmarkIntSlice2
//BenchmarkIntSlice2-12                897           1267922 ns/op         8053967 B/op          1 allocs/op
//BenchmarkIntSlice2-12                937           1265572 ns/op         8051816 B/op          1 allocs/op
//BenchmarkIntSlice2-12                943           1370291 ns/op         8051508 B/op          1 allocs/op
//BenchmarkIntSlice2-12                813           1357253 ns/op         8059172 B/op          1 allocs/op
//BenchmarkIntSlice2-12                916           1280988 ns/op         8052919 B/op          1 allocs/op
//BenchmarkIntSlice2-12                939           1279828 ns/op         8051713 B/op          1 allocs/op
//BenchmarkIntSlice2-12                936           1271710 ns/op         8051867 B/op          1 allocs/op
//BenchmarkIntSlice2-12                940           1396431 ns/op         8051662 B/op          1 allocs/op
//BenchmarkIntSlice2-12                882           1354317 ns/op         8054823 B/op          1 allocs/op
//BenchmarkIntSlice2-12                885           1353601 ns/op         8054649 B/op          1 allocs/op
//PASS
//ok      basego/clone    42.586s

// 结果： 性能总体提升0.5倍

func BenchmarkIntSlice(b *testing.B) {

	var s []int
	for i := 0 ; i < 1e6; i ++{
		s = append(s, rand.Int())
	}

	for i := 0 ; i < b.N ; i++{
		clone.IntSlice(s)
	}
}



func BenchmarkIntSlice2(b *testing.B) {

	var s []int
	for i := 0 ; i < 1e6; i ++{
		s = append(s, rand.Int())
	}

	for i := 0 ; i < b.N ; i++{
		clone.IntSlice2(s)
	}
}

