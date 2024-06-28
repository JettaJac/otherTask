package main

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

import (
	"fmt"
)

func convert(s string, numRows int) string {
	res1 := make([]string, len(s))
	res := ""
	var c int
	n := 0

	for i := 0; i < len(s); i++ {
		if i/numRows > n {
			n++
		}
		fmt.Println("N", n)
		if n == numRows+1 {
			n = 0
			c++
		}
		count := numRows*n + i%(numRows+1) + c
		if count < len(s) {

			fmt.Println("count", count)

			res1[count] = string(s[i])
		}
		n++
	}
	fmt.Println(res1)

	return res
}

// func main() {
// 	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
// 	// fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
// 	// fmt.Println(convert("A", 1) == "A")
// }
