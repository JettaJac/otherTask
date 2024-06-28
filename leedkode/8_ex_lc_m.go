// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
package main

import (
	// "fmt"
	"math"
	// "strconv"
)

func myAtoi(s string) int {

	res := 0.0
	var minus bool

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && (i == 0 || s[i-1] == ' ') {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			res = res*10 + float64(s[i]-'0')
		} else if s[i] == '-' && (i == 0 || s[i-1] == ' ') {
			minus = true
		} else if s[i] == '+' && (i == 0 || s[i-1] == ' ') {
			continue
		} else {
			break
		}
	}
	if minus {
		res = -res
	}

	if res < math.Pow(-2, 31) {
		res = math.Pow(-2, 31)
	}

	if res > math.Pow(2, 31)-1 {
		res = math.Pow(2, 31) - 1
	}
	return int(res)
}

// func main() {
// 	// fmt.Println(myAtoi("5 987"), 5)
// 	// fmt.Println(myAtoi(" -042"), -42)
// 	// fmt.Println(myAtoi(" +042"), 42)
// 	// fmt.Println(myAtoi("1337c0d3"), 1337)
// 	// fmt.Println(myAtoi("0-1"), 0)
// 	// fmt.Println(myAtoi("0+1"), 0)
// 	// fmt.Println(myAtoi("+15"), 15)
// 	// fmt.Println(myAtoi("-15"), -15)
// 	// fmt.Println(myAtoi("9+15"), 9)
// 	// fmt.Println(myAtoi("words and 987"), 0)
// 	// fmt.Println(myAtoi("   -0 123"), 0)
// 	// fmt.Println(myAtoi("   +0 123"), 0)
// 	// fmt.Println(myAtoi("   0 123"), 0)
// 	fmt.Println(myAtoi("9223372036854775808"), 2147483647)
// 	fmt.Println(myAtoi("-91283472332"), -2147483648)

// 	// tmp, _ := strconv.Atoi("   -0 123")
// 	// fmt.Println(tmp)

// 	// fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
// 	// fmt.Println(convert("A", 1) == "A")
// }
