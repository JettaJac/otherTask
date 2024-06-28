// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {

	var resStr []byte

	if x < 0 {
		resStr = append(resStr, '-')
		x *= -1
	}

	tmpStr := strconv.Itoa(x)
	fmt.Println(len(tmpStr))

	for i := len(tmpStr) - 1; i >= 0; i-- {
		resStr = append(resStr, tmpStr[i])
	}

	res, _ := strconv.Atoi(string(resStr))

	if float64(res) < math.Pow(-2, 31) || float64(res) > math.Pow(2, 31)-1 {
		res = 0
	}

	return int(res)
}

// func main() {
// 	fmt.Println(reverse(120), 321)
// 	// fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
// 	// fmt.Println(convert("A", 1) == "A")
// }
