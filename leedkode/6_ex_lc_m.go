package main

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

import (
// "fmt"
)

func convert(s string, numRows int) string { // Done

	if numRows == 1 {
		return s
	}

	tmpStr := make([]string, numRows)
	var res string
	whileStr := numRows + (numRows - 1)

	for i := 0; i < len(s); {
		j := 0
		for ; i < len(s) && j < numRows && i%(whileStr-1) < numRows; i++ {
			tmpStr[j] += string(s[i])
			j++
		}
		for n := 2; i < len(s) && i%(whileStr-1) >= numRows; i++ {
			tmpStr[numRows-n] += string(s[i])
			j++
			n++
		}
	}

	for _, str := range tmpStr {
		res += str
	}

	return res
}

// func main() {
// 	// 	fmt.Println(convert("PREDPRINIMATELU", 4) == "PIERRNTLEPIAUDM")
// 	// 	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
// 	// 	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
// 	// 	fmt.Println(convert("A", 1) == "A")
// 	// 	fmt.Println(convert("AB", 1) == "AB")
// 	// fmt.Println(convert("ABC", 2), "ACB")
// 	// fmt.Println(convert("", 1) == "")
// }
