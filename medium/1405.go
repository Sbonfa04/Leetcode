/*
A string s is called happy if it satisfies the following conditions:

    - s only contains the letters 'a', 'b', and 'c'.
    - s does not contain any of "aaa", "bbb", or "ccc" as a substring.
    - s contains at most a occurrences of the letter 'a'.
    - s contains at most b occurrences of the letter 'b'.
    - s contains at most c occurrences of the letter 'c'.

Given three integers a, b, and c, return the longest possible happy string. If there are multiple longest happy strings, return any of them. If there is no such string, return the empty string "".

A substring is a contiguous sequence of characters within a string.



Example 1:

Input: a = 1, b = 1, c = 7
Output: "ccaccbcc"
Explanation: "ccbccacc" would also be a correct answer.

Example 2:

Input: a = 7, b = 1, c = 0
Output: "aabaa"
Explanation: It is the only correct answer in this case.



Constraints:

    0 <= a, b, c <= 100
    a + b + c > 0
*/

package main

import (
	"sort"
)

func longestDiverseString(a int, b int, c int) string {
	type charCount struct {
		count int
		char  byte
	}

	result := []byte{}
	counts := []charCount{
		{a, 'a'},
		{b, 'b'},
		{c, 'c'},
	}

	for {
		sort.Slice(counts, func(i, j int) bool {
			return counts[i].count > counts[j].count
		})

		added := false
		for i := 0; i < 3; i++ {
			if counts[i].count == 0 {
				break
			}
			n := len(result)
			if n >= 2 && result[n-1] == counts[i].char && result[n-2] == counts[i].char {
				continue
			}
			result = append(result, counts[i].char)
			counts[i].count--
			added = true
			break
		}

		if !added {
			break
		}
	}

	return string(result)
}
