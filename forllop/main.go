package main

import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5}
	sum := 0
	for _, value := range intArr {
		sum +=value
		if value <= sum/len(intArr) {
			fmt.Println(value)
		}
	}
}

//func sort() {
//	intArr := []int{1,4,2,3,9,5,7,1,0,3,3,5}

//for i := 0; i < len(intArr); i++ {
//	for j := len(intArr) - 1; j > i; j-- {
//		if intArr[j-1] > intArr[j] {
//			tmp := intArr[j-1]
//			intArr[j-1] = intArr[j]
//			intArr[j] = tmp
//
//		}
//	}
//	//}
//
//
//	for i := 0; i < len(intArr) - 1; i++ {
//		for j := i+1; j < len(intArr); j++ {
//			if intArr[j] < intArr[i] {
//				tmp := intArr[i]
//				intArr[i] = intArr[j]
//				intArr[j] = tmp
//			}
//		}
//	}
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
