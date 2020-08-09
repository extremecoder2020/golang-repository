package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Please, type your name:")
	scannerName := bufio.NewScanner(os.Stdin)
	scannerName.Scan()
	userName := scannerName.Text()

	fmt.Println("Now, type your address:")
	scannerAddr := bufio.NewScanner(os.Stdin)
	scannerAddr.Scan()
	userAddr := scannerAddr.Text()

	var userInfo map[string]string
	userInfo = make(map[string]string)
	userInfo["name"] = userName
	userInfo["addr"] = userAddr

	jsonOut, _ := json.Marshal(userInfo)
	os.Stdout.Write(jsonOut)
}
