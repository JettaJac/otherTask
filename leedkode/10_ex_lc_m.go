// Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

// '.' Matches any single character.​​​​
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).

package main

import "fmt"

// func isMatch(s string, p string) bool {
// 	var tmp byte
// 	for i, j := 0, 0; i < len(s) && j < len(p); {
// 		fmt.Println(string(p[j]), string(s[i]), i, j, string(tmp))
// 		if p[j] == '*' && (tmp == s[i] || p[j-1] == '.') {
// 			fmt.Println("lll")
// 			for i < len(s) && p[j] == '*' && tmp == s[i] {

// 				i++
// 			}
// 			i++
// 			j++
// 			continue
// 		}

// 		fmt.Println(string(p[j]), string(s[i]), i, j)
// 		tmp = s[i]
// 		if p[j] == '.' {
// 			i++
// 			j++
// 			continue
// 		}

// 		if s[i] != p[j] {
// 			return false
// 		}
// 		i++
// 		j++
// 	}

//		return true
//	}
func isMatch(s string, p string) bool {
	var res string
	var tmp byte
	for i, j := 0, 0; i < len(s) && j < len(p); j++ {
		// fmt.Println(i, j)
		if s[i] == p[j] || p[j] == '.' || p[j] == '*' {

			// continue

			// fmt.Println(string(tmp), string(s[i]), string(p[j]))
			if j != 0 && p[j] == '*' && p[j-1] == '.' {
				// fmt.Println("string(tmp), string(s[i]), string(p[j-1])")
				tmp = s[i]
				// fmt.Println("string(tmp), string(s[i]), string(p[j-1])")
			}
			res += string(s[i])
			tmp = s[i]
			i++

			for i < len(s) && j < len(p) && p[j] == '*' && (tmp == s[i] /* || p[j-1] == '.'*/) {

				res += string(s[i])
				tmp = s[i]
				// if p[j-1] == '.' {
				// fmt.Println("i, j")
				// 	j++
				// }
				i++

			}

		} else if s[i] != p[j] && p[j+1] == '*' {
			j++
			// break
		}

		if i == len(s) && j < len(p)-1 && (p[j] == '*' && s[i-1] != p[j+1] || p[j] != '*' && s[i-1] == p[j+1]) {
			// fmt.Println("i, len(s)")
			res += string(p[j+1])
		}
	}
	// if p[len(p)-2] != '*' && s[len(s)-1] == p[len(p)-1] {
	// 	res += string(p[len(p)-1])
	// }
	fmt.Println(string(res), string(s))
	return s == res
}

func main() {
	// fmt.Println("aa", isMatch("aa", "a"), "false")
	// fmt.Println("aa", isMatch("aa", "a*"), "true")
	// fmt.Println("aaaap", isMatch("aaaap", "a*p"), "true")
	// fmt.Println("ajp", isMatch("ajp", "a.p"), "true")
	// fmt.Println("ab", isMatch("ab", ".*"), "true")
	// fmt.Println("abd", isMatch("abb", ".*"), "true")
	// fmt.Println("abd", isMatch("add", ".*"), "true")
	// fmt.Println("abd", isMatch("adb", ".*"), "false")
	// fmt.Println("ab", isMatch("ab", ".*c"), "false")
	// fmt.Println("ab", isMatch("ac", ".*c"), "true")
	// fmt.Println("ab", isMatch("acc", ".*c"), "true")
	// fmt.Println("ab", isMatch("acc", ".c"), "false")
	// fmt.Println("ab", isMatch("acc", "*c"), "false")
	// fmt.Println("abc", isMatch("abc", ".*"), "false")
	// fmt.Println("aaaa", isMatch("aaaa", "aaa"), "false")
	// fmt.Println("aaa", isMatch("aaa", "aaaa"), "false")
	// fmt.Println("aaa", isMatch("aaa", "ab*a"), "false")
	// fmt.Println("aab", isMatch("aab", "c*a*b"), "true")
	// fmt.Println("aab", isMatch("aab", "a*b*"), "true")
	// fmt.Println("aab", isMatch("ab", "a*b*"), "true")
	fmt.Println("aaa", isMatch("aaa", "ab*a*c*a"), "true")
}
