/*
 * Copyright (c) 2020 RongHe Ma.
 * Recording from daily work flow.
 */

package utils

import (
	"bytes"
	"encoding/binary"
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

// IntToBytesBigEndian int to Big-Endian []byte
func IntToBytesBigEndian(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// IntToBytesLittleEndian int to Little-Endian []byte
func IntToBytesLittleEndian(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToIntBigEndian Big-Endian []byte to int
func BytesToIntBigEndian(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

// BytesToIntLittleEndian Little-Endian []byte to int
func BytesToIntLittleEndian(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int(x)
}
