//Go의 특징

package main
import "fmt"
func main() {
  //Go는 모호하거나, 애매한 문법을 금지
  //후치연산자 허용 (i++), 전치 연산자 비허용 (++i 안됨)
  //증감연산 반환 값 없음. sum := i++ 와 같은 표현 불가
  //포인터(Pointer 허용, 연산 비허용)
  //주석 (//, /* 등)


  //example 1
  sum, i := 0,0
  for i<= 100 {
    //sum += i++
    sum += i // IDEA:
    i++
  }
 fmt.Println("ex1 :", sum)
}
