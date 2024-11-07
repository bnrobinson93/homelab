package main

import (
  "fmt"
  "maps"
)

func main() {
  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1:", v1)

  v3:= m["k3"] // this doesn't exist so is a zero value
  fmt.Println("v3:", v3)

  fmt.Println("len:", len(m))

  // delete one
  delete(m, "k2")
  fmt.Println("map:", m)

  // delete all
  clear(m)
  fmt.Println("map:", m)

  // optional return indicates presence; ignore acutal value
  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)

  n2 := map[string]int{"foo": 1, "bar": 2}
  if maps.Equal(n, n2) {
    fmt.Println("n == n2")
  }
}
