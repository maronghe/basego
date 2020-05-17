package utils_test

import (
	"basego/utils"
	"fmt"
	"testing"
)

func TestRanndom(t *testing.T) {
	fmt.Printf("%s\n", utils.RandomString(4))
}

func BenchmarkRandom(b *testing.B) {
	for i := 1; i <= 5; i++ {
		b.Run(fmt.Sprintf("i-%d", i), func(b *testing.B) {
			// b.ResetTimer()
			fmt.Printf(" %d %s \n", i, utils.RandomString(4))
		})
	}
}

func TestBigLittleEndian(t *testing.T) {
	var i int = 6
	bByte := utils.IntToBytesBigEndian(i)
	ni := utils.BytesToIntBigEndian(bByte)
	fmt.Printf("%d %b \n", ni, bByte) //	6 [0 0 0 110] 大端

	var j int = 6
	lByte := utils.IntToBytesLittleEndian(j)
	nj := utils.BytesToIntLittleEndian(lByte)
	fmt.Printf("%d %b \n", nj, lByte) // 6 [110 0 0 0] 小端

}
