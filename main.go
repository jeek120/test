package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

func main() {
	buf := bytes.Buffer{}

	binary.Write(&buf, binary.BigEndian, uint64(0xfffffffffffffffe))

	fmt.Println(buf.Bytes())
	/*
	d := filepath.Dir("D:\\soft\\gopath\\src\\github.com\\jeek120\\farm\\server\\agent\\agent.go")
	println(d)
	return
	// max4word := "456975"
	var max4word2 int64 = 456975
	max64big := "18446744073709551615"
	result := BigIntAdd(max64big, max4word2)
	println(result)

	n := wordtoint("hlhxczmxsyumqq")
	println(n)*/
}

func wordtoint(word string) uint64 {
	base := 26
	var n uint64
	n *= uint64(base)
	var d byte
	for _,r := range []byte(word) {
		d = r - 'a'
		n = n*uint64(base) + uint64(d)
	}
	return n
}

func lower(b byte) byte {
	return b | ' '
}

func BigIntAdd(numstr string, num int64) string {
	n, _ := new(big.Int).SetString(numstr, 10)
	m := new(big.Int)
	m.SetInt64(num)
	m.Add(n, m)
	return m.String()
}

func BigIntReduce(numstr string, num int64) string {
	n, _ := new(big.Int).SetString(numstr, 10)
	m := new(big.Int)
	m.SetInt64(-num)
	m.Add(n, m)
	return m.String()
}

func BigIntMul(numstr string, num int64)string{
	n, _ := new(big.Int).SetString(numstr, 10)
	m := new(big.Int)
	m.SetInt64(num)
	m.Mul(n, m)
	return m.String()

}

func BigIntDiv(numstr string, num int64)string{

	n, _ := new(big.Int).SetString(numstr, 10)
	m := new(big.Int)
	m.SetInt64(num)
	m.Div(n, m)
	return m.String()

}