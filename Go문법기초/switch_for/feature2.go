package main
import "fmt"
func main() {
//go특징2
//문장 끝 세미콜론 주의(;)
//자동으로 끝에 세미콜론 삽임
// 두 문장을 한문장에 표현할 경우 명시적으로 세미콜론 사용 but 권장 X
//반복문 및 제어문(조건문)(if,for)사용할 경우에 주의해야함
// 예제1
for i := 0; i <= 10; i++{
  fmt.Print("ex:");fmt.Println(i) //이런식으로도 표현 가능하지만 가독성 X
}

//예제2
for j := 10; j >= 0 ; j-- {
  fmt.Println("ex2:", j)
}


}
