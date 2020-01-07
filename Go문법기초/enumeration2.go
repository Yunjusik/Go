package main
import "fmt"
func main(){
  const (
    A = iota * 10
    B
    C
  )
  fmt.Println(A,B,C)
  //결과 0 , 1, 2 나옴
  //10곱하면 각각 0,10,20 나옴
  //iota쓰면 처음 실행 이후 1씩 increment해서 다음변수에 자동 상속
  //기본적으로 iota사용하면 0부터 시작함.
  //
}
