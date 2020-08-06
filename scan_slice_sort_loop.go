package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var groupInt []int
	for {
		fmt.Println("Put int number or X to exit:")
		var userResponse string
		fmt.Scan(&userResponse)

		checkResponse := userResponse

		if checkResponse == "x" {
			fmt.Println("Please, type a integer number!")
		} else if checkResponse == "X" {
			break
		} else {
			inputToInt, _ := strconv.Atoi(checkResponse)
			groupInt = append(groupInt, inputToInt)
			sort.Ints(groupInt)
			fmt.Println(groupInt)
		}
	}
}
