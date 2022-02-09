package main

import (
	"fmt"
	"os"

	"github.com/info700700/calc-go/interpreter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Ошибка. Неверное количество аргументов. Требуется 1 аргумент.")
		return
	}

	str := os.Args[1]

	result, err := interpreter.Exec(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
