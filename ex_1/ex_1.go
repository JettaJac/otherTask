package main

import (
	"fmt"
)

func main () {
	fmt.Println("Start")
	array1 := [5]int{1, 2, 3, 0, 0}
	m := 3
	array2 := [2]int{2, 5}
	n := 2

	slice := array1[0:m]
	fmt.Println(slice , m)
	for count := 0; count < m + n; count++ {
		for _, b := range array2 {
			if a > b {
				array1[count] = b
			}
		}
		fmt.Println(a)
	}
	fmt.Println(array1)
}