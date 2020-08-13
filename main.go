package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var numbers int
	fmt.Println("Введите любое число: ")
	_, err := fmt.Fscan(os.Stdin, &numbers)

	if err != nil {
		fmt.Println("ты ввел не число")
		os.Exit(1)
	}

	res := ""
	var del = numbers % 100
	var del2 = numbers % 10
	if del >= 11 && del <= 14 {
		res = "тов"
	} else if del2 == 1 {
		res = "т"
	} else if del2 >= 2 && del2 <= 4 {
		res = "та"
	} else {
		res = "тов"
	}

	resString := fmt.Sprintf("%d програмис%s", numbers, res)
	fmt.Println(resString)
}

func test1() (int, error) {
	err := errors.New("new error")
	return 0, err
}
