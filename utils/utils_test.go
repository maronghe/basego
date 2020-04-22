package utils_test

import (
	"fmt"
	"testing"

	"basego/utils"
)

func TestRandom(t *testing.T) {

	fmt.Println(utils.RandomString(10))

}
