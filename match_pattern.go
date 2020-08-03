package main

import (
	"fmt"
	"strings"
)

func main() {
	
  fmt.Printf("Enter an word: \n")
	var userWord string
	fmt.Scan(&userWord)
	userInput := userWord

	var firstLetter = strings.HasPrefix(userInput, "i")
	var lastLetter = strings.HasSuffix(userInput, "n")
	var randonLetter = strings.ContainsAny(userInput, "a")

	if firstLetter && lastLetter && randonLetter {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}
}
