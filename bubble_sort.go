package main

import (
	"fmt"
)

func main() {
    arr := []int{1, 10, 5, 2}
    len := len(arr)
    fmt.Printf("length: %d\n",len)
    for i := 0; i < len-1; i++ {
        for j := 0; j < len-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
	fmt.Println(arr)
        }
    }
    //fmt.Println(arr)
}


//length: 4
//[1 10 5 2]
//[1 5 10 2]
//[1 5 2 10]
//[1 5 2 10]
//[1 2 5 10]
//[1 2 5 10]
