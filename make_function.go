package main

import "fmt"

func main() {
  var sli = make([]int, 0, 10)
  sli = append(sli, 10, 100, 1000)
	fmt.Println(sli)
  fmt.Println(sli[2])
}

// [10 100 1000]
// 1000
