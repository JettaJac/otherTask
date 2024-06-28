// Given an integer x, return true if x is a palindrome, and false otherwise.
package main

// import "fmt"

func isPalindrome(x int) bool {
	original := x
	copyInt := 0
	// tmp := []int{}
	// var n int
	if x < 0 {
		return false
	}
	for x > 0 {
		copyInt = copyInt*10 + x%10
		x /= 10
	}

	// for i, j := 0, len(tmp)-1; i < len(tmp)/2 && j >= len(tmp)/2; i++ {
	// 	if tmp[i] != tmp[j] {
	// 		return false
	// 	}
	// 	j--
	// }

	return original == copyInt
}

// func main() {
// 	fmt.Println("121", isPalindrome(121), "true")
// 	fmt.Println("1221", isPalindrome(1221), "true")
// 	fmt.Println("12321", isPalindrome(12321), "true")
// 	fmt.Println("-12321", isPalindrome(-12321), "false")
// 	fmt.Println("13321", isPalindrome(13321), "false")
// 	fmt.Println("10", isPalindrome(10), "false")
// }
