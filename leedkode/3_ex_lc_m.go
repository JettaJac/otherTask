// Done
// Given a string s, find the length of the longest
// substring
//  without repeating characters.

// Example 1:
// Ввод: s = "pwwkew"
//  Вывод: 3
//  Объяснение: Ответ: "wke" длиной 3.
// Обратите внимание, что ответ должен быть подстрокой, "pwke" — это подпоследовательность, а не подстрока.

package main

import (
// "fmt"
)

// func lengthOfLongestSubstring(s string) int { // медленный
// 	myMap := make(map[byte]int)
// 	memory, tmp := 0, 0
// 	for i := 0; i < len(s); i++ {
// 		// fmt.Println(string(s[i]))
// 		if _, ok := myMap[s[i]]; !ok {
// 			tmp++
// 		} else {
// 			i = myMap[s[i]] + 1
// 			myMap = make(map[byte]int)
// 			tmp = 1

// 		}
// 		myMap[s[i]] = i
// 		memory = max(memory, tmp)
// 		fmt.Println(myMap, i, tmp)
// 	}
// 	return max(memory, tmp)
// }

func lengthOfLongestSubstring(s string) int { //можно еще оптимальнее на базе этого, толькл мы не проверяем наличие в мапе элемента, а используем значение по ключу
	myMap := make(map[byte]int)
	memory := 0
	left := 0
	if len(s) > 0 {
		myMap[s[left]] = 1
		memory = 1
	}

	for right := 1; right < len(s); right++ {
		_, ok := myMap[s[right]]
		if ok {
			for ok == true {
				delete(myMap, s[left])
				left++
				_, ok = myMap[s[right]]
			}
		}
		myMap[s[right]] = 1
		memory = max(memory, len(myMap))
	}

	return memory
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("adavdf"), 4)
// 	fmt.Println(lengthOfLongestSubstring(""), 0)
// 	fmt.Println(lengthOfLongestSubstring("D"), 1)
// 	fmt.Println(lengthOfLongestSubstring("DrD"), 2)
// 	fmt.Println(lengthOfLongestSubstring("abcbadcbb"), 4)
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"), 3)
// }
