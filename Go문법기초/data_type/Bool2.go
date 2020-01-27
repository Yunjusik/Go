package main

import "fmt"

func main() {
	//ex1 (논리연산자)
	fmt.Println("ex1 :", true && true)
	fmt.Println("ex2 :", true && false)
	fmt.Println("ex3 :", false && false)
	fmt.Println("ex4 :", true || true)
	fmt.Println("ex5 :", true || false)
	fmt.Println("ex6 :", false || false)
	fmt.Println("ex7 :", !true)
	fmt.Println("ex8 :", !false)

	//ex2 (비교연산자)
	num1 := 15
	num2 := 37

	fmt.Println("ex2-1 :", num1 < num2)
	fmt.Println("ex2-1 :", num1 > num2)
	fmt.Println("ex2-1 :", num1 >= num2)
	fmt.Println("ex2-1 :", num1 <= num2)
	fmt.Println("ex2-1 :", num1 != num2)
	fmt.Println("ex2-1 :", num1 == num2)
}
