package Unique_Generalized_Abbreviations

import (
	"fmt"
	"math"
)

/*
Given a word, write a function to generate all of its unique generalized abbreviations.

Generalized abbreviation of a word can be generated by replacing each substring of the word by the count of characters in the substring.
Take the example of “ab” which has four substrings: “”, “a”, “b”, and “ab”.
After replacing these substrings in the actual word by the count of characters we get all the generalized abbreviations: “ab”, “1b”, “a1”, and “2”.

Input: "BAT"
Output: "BAT", "BA1", "B1T", "B2", "1AT", "1A1", "2T", "3"

Input: "code"
Output: "code", "cod1", "co1e", "co2", "c1de", "c1d1", "c2e", "c3", "1ode", "1od1", "1o1e", "1o2",
"2de", "2d1", "3e", "4"

ref: https://leetcode-cn.com/problems/generalized-abbreviation/
*/

func generateAbbreviations(word string) []string {
	if word == "" {
		return []string{}
	}
	ans := make([]string, 0, int(math.Pow(2, float64(len(word)))))
	wLen := len(word)
	queue := make([]AbbreviatedWord, 0, int(math.Pow(2, float64(len(word)))))
	queue = append(queue, AbbreviatedWord{"", 0, 0})

	for len(queue) != 0 {
		aWord := queue[0]
		queue = queue[1:]
		if aWord.start == wLen {
			if aWord.count != 0 {
				aWord.Value = aWord.Value + fmt.Sprintf("%d", aWord.count)
			}
			ans = append(ans, aWord.Value)
		} else {
			queue = append(queue, AbbreviatedWord{aWord.Value, aWord.start + 1, aWord.count + 1})

			if aWord.count != 0 {
				aWord.Value = aWord.Value + fmt.Sprintf("%d", aWord.count)
			}
			aWord.Value = aWord.Value + string(word[aWord.start])
			queue = append(queue, AbbreviatedWord{aWord.Value, aWord.start + 1, 0})
		}
	}

	return ans
}

type AbbreviatedWord struct {
	Value string
	start int
	count int
}

//func generateAbbreviations(word string) []string {
//	if word == "" {
//		return []string{}
//	}
//	var ans []string
//	var queue []string
//	queue = append(queue, "")
//
//	for len(queue) != 0 {
//		curStr := queue[0]
//		queue = queue[1:]
//
//		if len(curStr) == len(word) {
//			// 替换_
//			newStr := replaceUnderline(curStr)
//			ans = append(ans, newStr)
//			continue
//		}
//
//		queue = append(queue, curStr+"_")
//		queue = append(queue, curStr+string(word[len(curStr)]))
//	}
//
//	return ans
//}
//
//func replaceUnderline(str string) string {
//	var (
//		start  = -1
//		end    = -1
//		newStr = ""
//	)
//	for i, v := range str {
//		if v != '_' {
//			if start == end && end == -1 {
//				newStr = newStr + string(v)
//				continue
//			}
//			newStr = newStr + fmt.Sprintf("%d", end-start+1) + string(v)
//			start = -1
//			end = -1
//		} else {
//			if start == -1 {
//				start = i
//			}
//			end = i
//		}
//	}
//
//	// 对于全是_的str
//	if start != -1 {
//		newStr = newStr + fmt.Sprintf("%d", end-start+1)
//	}
//
//	return newStr
//}