package main

// Given a string s, return the longest palindromic substring in s.

// import "fmt"

func longestPalindrome(s string) string {
	res := ""
	if len(s) > 0 {
		c := len(s) / 2
		memory := make([]int, 2)

		memory[0] = c
		memory[1] = 1

		for j := c; j <= len(s) && len(s)-j > memory[1]/2; j++ {
			count(j, s, memory)
		}

		for i := c - 1; i >= 0 && i >= memory[1]/2; i-- {
			count(i, s, memory)
		}

		size := memory[0] + memory[1]
		res = s[memory[0]:size]
	}
	return res
}

func count(c int, s string, memory []int) {
	tmpMemory := make([]int, 2)
	tmpMemory[0] = c
	tmpMemory[1] = 1

	l := c - 1
	r := c + 1
	for l >= 0 && s[c] == s[l] {
		tmpMemory[0] = l
		tmpMemory[1] += 1
		l--
	}

	for r < len(s) && s[c] == s[r] {
		tmpMemory[1] += 1
		r++
	}

	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			if tmpMemory[1] > memory[1] {
				memory[0] = tmpMemory[0]
				memory[1] = tmpMemory[1]
			}
			break
		}
		tmpMemory[0] = l
		tmpMemory[1] = r - l + 1
		l--
		r++
	}
	if tmpMemory[1] > memory[1] {
		memory[0] = tmpMemory[0]
		memory[1] = tmpMemory[1]
	}
}

// func main() {
// 	// fmt.Println(longestPalindrome("cbdfg"), "d")
// 	fmt.Println(longestPalindrome(""), "- ничего")
// 	fmt.Println(longestPalindrome("ac"), "c")
// 	fmt.Println(longestPalindrome("a"), "a")
// 	fmt.Println(longestPalindrome("add"), "dd")
// 	fmt.Println(longestPalindrome("dd"), "dd")
// 	fmt.Println(longestPalindrome("ddd"), "ddd")
// 	fmt.Println(longestPalindrome("babad"), "aba")
// 	fmt.Println(longestPalindrome("cbbd"), "bb")
// 	fmt.Println(longestPalindrome("krdeedrk"), "krdeedrk")
// 	fmt.Println(longestPalindrome("krdedrk"), "krdedrk")
// 	fmt.Println(longestPalindrome("krdedrkdd"), "krdedrk")
// 	fmt.Println(longestPalindrome("rdeeedr"), "rdeeedr")
// 	fmt.Println(longestPalindrome("srdeeedr"), "rdeeedr")
// 	fmt.Println(longestPalindrome("jrdeeedrjs"), "jrdeeedrj")
// 	fmt.Println(longestPalindrome("scbbrdeedr"), "rdeedr")
// 	fmt.Println(longestPalindrome("scdabbr"), "bb")
// 	fmt.Println(longestPalindrome("dddd"), "dddd")
// 	fmt.Println(longestPalindrome("ddddd"), "ddddd")
// 	fmt.Println(len(longestPalindrome("dddddd")) == len("dddddd"))
// 	fmt.Println(longestPalindrome("dddddddd"), "dddddddd")
// 	fmt.Println(longestPalindrome("ddfghjk"), "dd")
// 	fmt.Println(longestPalindrome("ddfghhjk"), "dd")
// 	fmt.Println(longestPalindrome("dfghjkk"), "kk")
// 	fmt.Println(longestPalindrome("dfglhjkk"), "kk")
// 	fmt.Println(longestPalindrome("ababababa"), "ababababa")
// 	fmt.Println(longestPalindrome("abadd"), "aba")
// 	fmt.Println(longestPalindrome("banana"), "anana")
// 	fmt.Println(longestPalindrome("aaaabbbbbbccccccbbbbbbaaaa"), "aaaabbbbbbccccccbbbbbbaaaa")
// 	// fmt.Println(longestPalindrome("aaaabbbbbbbbbbccccccccbbbbbbbbbbaaaa"), "aaaabbbbbbbbbbccccccccbbbbbbbbbbaaaa")
// 	// fmt.Println(longestPalindrome("aaaabbbbbbbbbbcccccccbbbbbbbbbbaaaa"), "aaaabbbbbbbbbbcccccccbbbbbbbbbbaaaa")
// 	fmt.Println(len(longestPalindrome("aaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaa")) == 996)
// 	// fmt.Println(len("aaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaa"))
// }
