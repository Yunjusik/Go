package main
import "fmt"
func main(){
  const (
_ = iota
A
B
_
C
  ) // 내가 생략하고싶을때 _밑줄 사용해서 넘김. A는 1부터 들어감
  const(
    _= iota + 0.7*3
    DEFAULT //1+ 2.1
    SILVER //2+2.1
    _
    GOLD    //4+2.1
    PLATINUM//5+2.1
  )
  fmt.Println("D:",DEFAULT)
  fmt.Println("S:",SILVER)
  fmt.Println("G:",GOLD)
  fmt.Println("P:",PLATINUM)



}
