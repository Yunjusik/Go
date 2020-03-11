package main

import (
	"fmt"
)

func main() {
	//예제1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("Ex1:", n1+n2)
	fmt.Println("Ex1:", n1-n2)
	fmt.Println("Ex1:", n1*n2)
	fmt.Println("Ex1:", n1/n2)
	fmt.Println("Ex1:", n1%n2)
	fmt.Println("Ex1:", n1<<2) //2bit 이동, 왼 ,우
	fmt.Println("Ex1:", n1>>2)
	fmt.Println("Ex1:", ^n1) //보수

	var n3 int = 12
	var n4 float32 = 8.2
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	// fmt.Println("ex3:", n3+n4) int , float32는 연산 아안됨
	fmt.Println("ex3:", float32(n3)+n4)
	fmt.Println("ex3:", n3+int(n4))    //
	fmt.Println("ex3:", n5+uint16(n6)) //원래 1024+ 120000으로 121024가 나와야하는데 55488이 나옴.
	//이는 uint16의 범위제한때문에 n6값이 하향되서 더해지는것.
	//그래서 데이터형의 가용범위를 아는것이 중요합니다.

}
