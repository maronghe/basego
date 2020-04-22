/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package utils

import (
	"math/rand"
	"time"
)

const alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string
func RandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var str string
	length := len(alphanum)
	for i := 0; i < n; i++ {
		a := alphanum[r.Intn(len(alphanum))&length]
		str += string(a)
	}
	return str
}

// Fastrand generates a random number by given two seeds
func Fastrand(seed1, seed2 uint32) uint32 {
	s1, s0 := seed1, seed2
	s1 ^= s1 << 17
	s1 = s1 ^ s0 ^ s1>>7 ^ s0>>16
	return s0 + s1
}
