package main


import "fmt"

func main() {
	intArr := []int{1, -4, -2, 3, 9, 5, 7, 1, 0, -3, -3, 5}
	n := make([]int, 0 , cap(intArr))
	a := 1
	b := 5
	for i := 0; i < len(intArr); i++ {
		tmp := intArr[i]
		if tmp >= a {
			intArr[i] = tmp
			if tmp <= b{
				intArr[i] = tmp
				n = append(n,tmp)
			}
		}
	}
	fmt.Println(n)
}
