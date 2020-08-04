package main

import (
	"fmt"
)

func main() {
	x := [3]int{1, 2, 3}
	for i, v:= range x {
	    fmt.Printf("index: %d | value: %d \n",i,v);
	}
}
