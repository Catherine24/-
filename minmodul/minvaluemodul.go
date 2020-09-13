package main

import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5}
	min := 100
	for _, v := range intArr{
		if v < 0{
			v = v * -1
		}
		if v < min{
			min = v
		}

	}
	fmt.Println(min)
}
