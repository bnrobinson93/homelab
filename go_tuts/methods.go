package main

import "fmt"

type rect struct {
  width, height int
}

func (r *rect) area() int { // receiver type = *rect
  return r.width * r.height
}

func (r rect) perim() int { // receiver value
  return 2*r.width + 2*r.height
}

func main() {
  r := rect{width: 10, height: 5}
  fmt.Println("Area: ", r.area())
  fmt.Println("Perimeter: ", r.perim())

  rp := &r
  fmt.Println("Area: ", rp.area())
  fmt.Println("Perimeter: ", rp.perim())
}
