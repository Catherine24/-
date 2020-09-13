package main

import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5}
	sum := 0
	negative := 0
	for i := 0; i < len(intArr); i++ {
		tmp := intArr[i]
		if tmp > 0 && negative >=0{
			continue // пропускаем текущее
		} else if tmp < 0 {
			negative = tmp
		}

		if tmp <= 0 {
			tmp= tmp * -1
		}

		sum += tmp
	}
	fmt.Println(sum)
}
