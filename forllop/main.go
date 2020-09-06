package main

import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5}

	for i := 0; i < len(intArr); i++ {
		for j := len(intArr) - 1; j > i; j-- {
			if intArr[j-1] > intArr[j] {
				tmp := intArr[j-1]
				intArr[j-1] = intArr[j]
				intArr[j] = tmp
			}
		}
	}
	fmt.Println(intArr[0], intArr[1])
}

//
//

//
//
//	for i := 0; i < len(intArr); i++ {
//		fmt.Println(i, intArr[i])
//	}
//
//	for _, value := range intArr {
//		fmt.Println(value)
//	}
//
//}
