package main

import "fmt"

func main() {
	var a [2]string
	var b [3]int
	
	a[0] = "Hello"
	a[1] = "World"
	
	b[0] = 123
	b[1] = 345
	b[2] = 567
	
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	
	fmt.Println(b[0], b[1])
	fmt.Println(b)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
