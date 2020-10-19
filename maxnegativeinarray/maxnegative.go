package main

import "fmt"

func main() {
	intArr := []int{-1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5, 4, -500}
	maxnegative := -1000000
	for _, x := range intArr{
		if x < 0{
			if maxnegative<x {
				maxnegative = x
			}
		}
	}
	fmt.Println(maxnegative)
}

