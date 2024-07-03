/*You are given a string s and an array of strings words. All the strings of words are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of words concatenated.

For example, if words = ["ab","cd","ef"], then "abcdef", "abefcd", "cdabef", "cdefab", "efabcd", and "efcdab" are all concatenated strings. "acdbef" is not a concatenated string because it is not the concatenation of any permutation of words.
Return an array of the starting indices of all the concatenated substrings in s. You can return the answer in any order.
*/

package main

import (
	"fmt"
	"strings"
)

func findSubstring(s string, words []string) []int {
	result := []int{}

	size, lenWords := 0, 0

	if len(words) >= 0 {
		size = len(words[0])
		lenWords = len(words) * size
	}
	// fmt.Println(size, lenWords)

	for i := 0; i < len(s)-lenWords+1; i += size {
		tmpStr := s[i : i+lenWords]
		// tmp := 0
		testMap := make(map[int]int)
		fmt.Println(tmpStr)
		for j := 0; j < len(words); j++ {
			// fmt.Println(tmpStr, words[j])
			id := strings.Index(tmpStr, words[j])
			// fmt.Println(tmpStr, words[j])
			fmt.Println(id, size, id%len(words))
			if id%size != 0 || id == -1 {
				fmt.Println("delete")
				break
			}
			testMap[j] = 1
			// tmp = id
			fmt.Println(j, len(words)-1, len(testMap), len(words))
			fmt.Println(testMap)
			if j == len(words)-1 && len(testMap) == len(words) {
				result = append(result, i)
			}
		}
		fmt.Println("____")
	}

	return result
}

// func main() {
// 	// fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}), "[0,9]")
// 	// fmt.Println(findSubstring("barsa", []string{"ar", "ba"}), "[]")
// 	// fmt.Println(findSubstring("babababababa", []string{"ar", "ba"}), "[]")
// 	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}), "[]")
// 	// fmt.Println(findSubstring("basarkdj", []string{"sa", "ba"}), "[0]")
// 	// fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}), "[6,9,12]")
// 	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}), "[8]")
// }
