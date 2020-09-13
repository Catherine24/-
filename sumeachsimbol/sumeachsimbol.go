package main

import "fmt"

func main() {
	intArr := []int{12, 104, 81}
	sum := 0

	for _, v := range intArr{
		c := v
		for  {
			b := c % 10
			sum +=b
			c = c/10
				if c<=0{
					break
			}

		}
	}
	fmt.Println(sum)
}


