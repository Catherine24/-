package main

import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5, 4, 10}
	sum := 0
	q := 0
	middle := 0
	for _, a := range intArr{
		if a > 0 {
			sum += a
			q += 1
		}
		middle = sum/q
	}
	fmt.Println(middle)
}

