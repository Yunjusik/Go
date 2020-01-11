package main
import "fmt"
func main(){
//예제 1번
Loop1:
for i := 0; i<5; i++ {
  for j := 0; j<5; j++{
    if i == 2 && j == 4 {
      break Loop1   //특정조건일때 loop1을 탈출. 따라서 i,j가 2,4전까지만 출력
      // 만약 break만 있고 loop1을 다 지워버리면, break는 해당 조건(2,4)에만 걸리고, 다시 for문을 돌아 3,1부터는 println이 동작
    }
    fmt.Println("ex :",i,j)
  }
}
//예제2번
for i := 0; i<10; i++ {
  if i % 2 == 0 {
    continue          //i가 2의배수일때 아래로 내려가지않고, 다시 조건문으로 회귀함. 따라서 홀수만 출력됨
  }
  fmt.Println("ex2 :", i)
}


Loop2: //참고 : loop lable 바로 아래에는 소스코드, println등 사용 불가, 바로아래에는 for써야함.
for i := 0; i < 3; i++{
  for j := 0 ; j<3; j++{
    if i == 1 && j == 2{
      continue Loop2 //이 조건에서 아래로 안가고 다시 loop2로 올라감. (1,2)일때 출력 안하고 다시 for문으로돌아감.
    }
    fmt.Println("ex3 : ", i,j)
  }
}


}
