package main
import "fmt"
func main(){
  //예제1
  sum1 := 0
  for i :=0; i<=100; i++{
    sum1 += i
  }
  fmt.Println("ex1:",sum1)
  //예제2
  sum2, i:=0,0
  for i<=100 {
    sum2 += i //
    i++ // *참고* 후치연산은 go에서 반환 값이 없음. 그래서 후치연산값을 다른 변수로 대입할 수 없음
  }
    fmt.Println("ex2:",sum2)
//예제3
sum3, i:= 0,0
for {
  if i>100 {
    break
  }
  sum3 += i
  i++
}
      fmt.Println("ex3:",sum3)

// 예제4
for i, j := 0, 0; i <= 10; i,j = i+1, j+10 {
  fmt.Println("ex4:",i,j) //i는 1씩, j는 10씩 커져감.
//에러발생 패턴
//for i, j := 0, 0; i <= 10; i++,j  += 10 { //i++ ->후치연산은 반환값이 없어 에러
//}





}
