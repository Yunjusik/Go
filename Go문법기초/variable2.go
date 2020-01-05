package main

import "fmt"
func main(){
  //여러개 선언하는
  var(
    name string ="machine"
    height int32
    weight float32
    isRunning bool
    )//변수선언의 가독성 향상
   height =250
   weight = 350.56
   isRunning = true

   fmt.Println("name: ", name, "height:", height, "weight:",weight, "isRunning:", isRunning)
  
}
