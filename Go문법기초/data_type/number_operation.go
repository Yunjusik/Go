//데이터 타입 : numeric 연산(1)

package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자연산 (산술, 비교)
	//타입이 같아야 가능. 정수, 실수 등등 같은 도메인에서 해야함
	// 다른 타입끼리는 반드시 형 변환 후 연산을 해주자
	//형 변환이 없을 경우, 예외 에러 가 발생함
	// +, -, *, %, / , <<, >> , &, ^

	//예제1
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64 //해당 데이터 범위에서 나올수있는 최댓값 반환

	fmt.Println("ex1:", n1)
	fmt.Println("ex2:", n2)
	fmt.Println("ex3:", n3)
	fmt.Println("ex4:", n4)
	fmt.Println("ex4:", math.MaxInt8) // int가 uint보다 가용범위가 더 적음. 대략 절반정도
	fmt.Println("ex4:", math.MaxInt16)
	fmt.Println("ex4:", math.MaxInt32)
	fmt.Println("ex4:", math.MaxFloat32)

	n5 := 100000 // int
	n6 := int16(10000)
	n7 := uint8(100) //uint 255까지, 100말고 300값넣으면 overflow 나옴
	//fmt.Println("ex2:", n5+n6)  //n5, n6끼리는 형태가 다르므로, 에러나옴. int 와 int16은 다르다!
	fmt.Println("ex2:", n5+int(n6))
	//fmt.Println("ex2:", n6 + n7)
	fmt.Println("ex2:", n6+int16(n7))
	fmt.Println("ex2: ", n6 > int16(n7)) //비교 연산자도 타입이 같아야 비교함.
	fmt.Println("ex2:", n6-int16(n7) > 5000)
	fmt.Println(math.MaxInt16)
}
