package main

import (
	"fmt"
)

func Add(x, y int) int {
   return x + y
}

func main() {
	y := Add(2, 3)
	fmt.Println(y)
}
