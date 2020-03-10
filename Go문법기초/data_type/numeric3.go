// 데이터 타입  :숫자2
package main

import "fmt"

func main() {
	//실수 (부동소수점)
	// float32 (7자리) float32(15자리, 소수점)

	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.0202222
	var num4 float32 = 10.0

	//지수 표기법
	var num5 float32 = 14e6
	var num6 float64 = .156875E+3
	var num7 float64 = 5.32521e-10

	fmt.Println("ex1:", num1)
	fmt.Println("ex1:", num2)
	fmt.Println("ex1:", num3)
	fmt.Println("ex1:", num4)
	fmt.Println("ex1:", num4-0.1)
	fmt.Println("ex1:", float32(num4-0.1))
	fmt.Println("ex1:", float64(num4-0.1)) //형이 다르면 값이 논리적으론 같아도, 코드상에선 다를 수 있음
	fmt.Println("ex1:", num5)
	fmt.Println("ex1:", num6)
	fmt.Println("ex1:", num7)
}
