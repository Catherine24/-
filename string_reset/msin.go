package main

import (
	"fmt"
	"strconv"
)

func main()  {
	str := "Heeello woooorld!"
	res := ""
	for i := 0; i < len(str); i++ {
		count := 1
		ch1 := string(str[i])
		for j := i+1; j<len(str); j++ {
			ch2 := string(str[j])
			if ch1 != ch2 {
				break
			}
			count++
			i++
		}

		if count >= 3 {
			res += ch1 + strconv.Itoa(count)
		} else if count == 2 {
			res += ch1 + ch1
		} else {
			res += ch1
		}
	}

	fmt.Println(res)

	// He3llo world!
}