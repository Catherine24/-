package main

import "fmt"

func main() {
	intArr := []int{1, 4,  3, 4, -3, 5, 4, 10}
	max := -1000000
	repeat_number := 0
	for i, x:= range intArr{
		repeat:= 0
		for j:=i+1; j < len(intArr); j++{
			tmp:=intArr[j]
			if tmp == x {
				repeat +=1
			}
		}
		if repeat > max{
			max = repeat
			repeat_number = x
		}
	}
	fmt.Println(repeat_number)
}
