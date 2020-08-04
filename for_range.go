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

/*
output:
index: 0 | value: 1 
index: 1 | value: 2 
index: 2 | value: 3
*/
