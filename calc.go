package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error. Invalid argument amount. One argument expected.")
		return
	}
	fmt.Println(os.Args[1])
}
