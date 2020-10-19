package main

import "fmt"

func main() {
	intArr := []int{5, 4,  3, 4, -3, 5, 4, 10}
	min := 1000000
	max := -1000000
	indexmin := 0
	indexmax := 0

	for i,x := range intArr{
		if min > x{
			min = x
			indexmin = i
		}
		if max < x{
			max = x
			indexmax = i
		}
	}
	intArr[indexmin] = max
	intArr[indexmax] = min
	fmt.Println(intArr)
}
