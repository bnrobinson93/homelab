package main

import (
  "fmt"
)

type base struct {
  num int
}

func (b base) describe() string {
  return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
  base
  str string
}

func main() {
  co := container{
    base: base{
      num: 1,
    },
    str: "some name",
  }

  // co.num pulls from base
  fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

  // Or you can do it explicitly
  fmt.Println("also num:", co.base.num)


  type describer interface {
    describe() string
  }

  var d describer = co
  fmt.Println("describer:", d.describe())
}
