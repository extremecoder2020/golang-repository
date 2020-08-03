package main

import (
	"fmt"
)

func main() {
  var response float64
  fmt.Scanln(&response)
  forDecimal(response)
}   
    
func forDecimalN(x float64){
  fmt.Printf("%d", int16(x))
};
