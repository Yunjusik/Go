package main
import "fmt"
func main(){
//열거형 - 상수를 사용하는, 일정 규칙에 따라 숫자를 계산 및 증가시키는 묶음

  const (
  JAN = 1
  Feb = 2
  Mar = 3
  Apr = 4
  May = 5
  Jun = 6
  ) // 위와같은 경우 상수가 1씩 증가되는 규칙성이 있음.

  fmt.Println(JAN)
  fmt.Println(Feb)
  fmt.Println(Mar)
  fmt.Println(Apr)
  fmt.Println(May)
  fmt.Println(Jun)



}
