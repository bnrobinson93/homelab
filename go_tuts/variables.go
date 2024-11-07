package main
import (
  "fmt"
  "math"
)

const s string = "constant"

func main() {
  var a = "initial"
  fmt.Println(a)

  var b, c int = 1,2
  fmt.Println(b,c)

  var d = true
  fmt.Println(d)

  var e int // zero value: 0, false, ""
  fmt.Println(e)

  f := true
  fmt.Println(f)


  // const
  fmt.Println(s)

  const n = 500000000
  const d = 3e20 / n
  fmt.Println(d)
  fmt.Println(int64(d))
  fmt.Println(math.Sin(n))
}
