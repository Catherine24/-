package main

import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5}
	sum := 0
	negative := true
	for _, x := range intArr{

		if negative {
			if x < 0 {
				negative = false
			}
			continue
		}
		if x <= 0 {
			x= x * -1
		}
		sum += x
	}
	fmt.Println(sum)
}