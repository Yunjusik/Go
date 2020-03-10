// 데이터 타입  :숫자1

package main

import "fmt"

func main() {
	//데이터타입이 숫자형일 경우
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned(양수)
	// 정수: 8진수(0), 16진수(0x), 10진수

	var num1 int32 = 17
	var num2 int = -68
	var num3 int = 0631       // 0으로 시작해서 8진수로 인식함
	var num4 int = 0x32fa2c75 //0x로 시작해서 16진수 인식

	fmt.Println("ex1", num1)
	fmt.Println("ex1", num2)
	fmt.Println("ex1", num3)
	fmt.Println("ex1", num4)
}
