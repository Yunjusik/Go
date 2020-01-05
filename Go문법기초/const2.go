package main
import "fmt"
func main() {
  const a, b int = 10,99
  const c,d,e = true, 0.84, "test"
  const (
    x,y int16=59,89
    i,k      = "Data", 78000
  )
fmt.Println(a,b,c,d,e)
fmt.Println(x,y,i,k)
}
