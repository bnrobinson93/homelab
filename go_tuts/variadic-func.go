package main

import "fmt"

// accept any number of ints as arguments
func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0

  // ignore the index
  for _, num := range nums { // type of nums = []int
    total += num
  }
  fmt.Println(total)
}

func main() {
  sum(1, 2)
  sum(1, 2, 3)
  nums := []int{1,2,3,4} // a slice
  sum(nums...) // spreads the slice
}
