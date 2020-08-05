package main

import (
	"fmt"
)

func main() {

	arrayOne := [...]int{1, 2, 3, 4}
	sliceOne := arrayOne[1:3]
	
	fmt.Println(sliceOne)
	fmt.Println(sliceOne[0])
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))
}
