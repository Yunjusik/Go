package main

import (
	"math"
)

func main() {
	//ex1(오버플로우)
	var n1 uint8 = math.MaxUint8 + 1 // 넘쳐 흘러  , constant 256 overflow MaxUint8
	var n1 uint8 = math.MaxUint16 + 1
	var n1 uint8 = math.MaxUint32 + 1
	var n1 uint8 = math.MaxUint64 + 1

	//ex2 (overflow error2, 범위 미만)

	var n1 uint8 = -1 //constant -1 overflows uint8
}
